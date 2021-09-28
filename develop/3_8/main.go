package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"syscall"

	"github.com/mitchellh/go-ps"
)

const (
	delim      = "|"
	fork       = "&"
	colorGreen = "\033[32m"
	colorWhite = "\033[37m"
	r          = true
)

var (
	writer io.Writer = os.Stdout
	reader io.Reader = os.Stdin
	isUdp  bool
)

func init() {
	flag.BoolVar(&isUdp, "u", false, "Use UDP instead of the default option of TCP.")
}

func main() {
	mainminic()
}

func mainminic() {
	// родительский процесс
	for {

		reader := bufio.NewReader(os.Stdin)
		if !r {
			hat()
		}
		line, _ := reader.ReadString('\n')

		if line == "exit\n" {
			break
		}
		if line == "\n" {
			continue
		}
		if r {

			subdata := strings.Split(line[:len(line)-1], ":")
			netc(subdata[0], subdata[1])
			return
		} else {

			mimicshell(line[:len(line)-1])
		}

	}

}

func mimicshell(line string) {
	subdata := strings.Split(line, fork)
	for i := range subdata {
		ret, _, err := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0, 0)
		if err != 0 {
			os.Exit(2)
		}
		if ret == 0 && len(subdata) == 1 {
			// потомок, умирает если & не имеются
			os.Exit(0)
		}

		if ret == 0 || len(subdata) == 1 {
			// потомок, если & имеются
			data := strings.Split(subdata[i], delim)
			rez := " "
			for j := range data {
				var err error

				rez, err = shell(data[j] + rez)
				if err != nil {
					println(err)
				}
			}
			if len(rez) > 1 {
				fmt.Println(rez[1:])
			}
			if ret == 0 {
				os.Exit(0)
			}
		}
	}
}

func shell(command string) (rez string, err error) {

	args := strings.Fields(command)
	if len(args) == 0 {
		return "", nil
	}
	switch args[0] {
	case "pwd":
		rez, err = os.Getwd()
	case "cd":
		err = os.Chdir(args[1])
	case "echo":
		rez = args[len(args)-1]
	case "ps":
		var A, o bool
		psargc := make([]string, 0)
		for i := 0; i < len(args); i++ {
			switch args[i] {
			case "-A":
				A = true
			case "-o":
				o = true
				i++
				buf := strings.Split(args[i], ",")
				psargc = append(psargc, buf...)
			}
		}
		var procs []ps.Process
		if !o {
			rez = "PID\tCMD\n"
		}
		procs, err = ps.Processes()
		for _, p := range procs {
			if os.Getppid() == p.Pid() || A {
				if !o {
					rez += fmt.Sprintf("%d\t%s\n", p.Pid(), p.Executable())
				} else {
					for j := range psargc {
						switch psargc[j] {
						case "pid":
							rez += strconv.Itoa(p.Pid())
						case "CMD":
							rez += p.Executable()
						default:
						}
					}
				}
			}
		}
	case "kill":
		for i := range args {
			if pid, _ := strconv.Atoi(args[i]); i > 0 {
				syscall.Kill(pid, syscall.SIGINT)
			} else {
				println(err)
			}
		}
	case "exec":
		rez, err = exec(args[1])
	}
	return " " + rez, err
}

func exec(path string) (string, error) {
	rez := ""
	data, err := ioutil.ReadFile(path)
	if err != nil {
		mimicshell(path)
	} else {
		comm := strings.Split(string(data), "\n")

		for i := 0; i < len(comm); i++ {
			if comm[i] == "" {
				continue
			}
			mimicshell(comm[i])
		}
	}
	return rez, nil
}

func hat() {
	wd, _ := os.Getwd()
	args := strings.Split(wd, "/")
	fmt.Print(string(colorGreen), "[minicshell@minicshell ", string(colorWhite), args[len(args)-1], string(colorGreen), "]$ ", string(colorWhite))
}

func netc(host, Port string) {
	if len(Port) > 0 {
		host += ":" + Port
	}
	flag.Parse()

	if !isUdp {
		log.Println("TCP protocol")
		if host != "" {
			con, err := net.Dial("tcp", host)
			if err != nil {
				log.Fatalln(err)
			}
			log.Println("Connected to", host)
			tcp_handle(con)
		} else {
			flag.Usage()
		}
	} else {
		log.Println("UDP protocol")
		if host != "" {
			addr, err := net.ResolveUDPAddr("udp", host)
			if err != nil {
				log.Fatalln(err)
			}
			log.Println("Resolved UDP address:", addr)
			con, err := net.DialUDP("udp", nil, addr)
			if err != nil {
				log.Fatalln(err)
			}
			udp_handle(con)
		}
	}
}

func tcp_handle(con net.Conn) {
	chan_to_stdout := stream_copy(con, writer)
	chan_to_remote := stream_copy(reader, con)
	select {
	case <-chan_to_stdout:
		log.Println("Remote connection is closed")
	case <-chan_to_remote:
		log.Println("Local program is terminated")
	}
}

// Performs copy operation between streams: os and tcp streams
func stream_copy(src io.Reader, dst io.Writer) <-chan int {
	buf := make([]byte, 1024)
	sync_channel := make(chan int)

	go func() {
		defer func() {
			if con, ok := dst.(net.Conn); ok {
				con.Close()
				log.Printf("Connection from %v is closed\n", con.RemoteAddr())
			}
			sync_channel <- 0 // Notify that processing is finished
		}()
		for {
			var nBytes int
			var err error
			nBytes, err = src.Read(buf)
			if err != nil {
				// конец ввода
				if err != io.EOF {
					if _, ok := src.(strings.Reader); ok {

						log.Println("Reader")
					}
					log.Printf("Read error: %s\n", err)
				}
				break
			}
			// fmt.Println("=======================")
			// fmt.Print(string(buf) + "\n")
			// fmt.Println("=======================")
			_, err = dst.Write(buf[0:nBytes])
			if err != nil {
				log.Fatalf("Write error: %s\n", err)
			}
		}
	}()
	return sync_channel
}

func udp_handle(con net.Conn) {
	in_channel := accept_from_udp_to_stream(con, writer)
	log.Println("Waiting for remote connection")
	remoteAddr := <-in_channel
	log.Println("Connected from", remoteAddr)
	out_channel := put_from_stream_to_udp(reader, con, remoteAddr)
	select {
	case <-in_channel:
		log.Println("Remote connection is closed")
	case <-out_channel:
		log.Println("Local program is terminated")
	}
}

//Accept data from UPD connection and copy it to the stream
func accept_from_udp_to_stream(src net.Conn, dst io.Writer) <-chan net.Addr {
	buf := make([]byte, 1024)
	sync_channel := make(chan net.Addr)
	con, ok := src.(*net.UDPConn)
	if !ok {
		log.Printf("Input must be UDP connection")
		return sync_channel
	}
	go func() {
		var remoteAddr net.Addr
		for {
			var nBytes int
			var err error
			var addr net.Addr
			nBytes, addr, err = con.ReadFromUDP(buf)
			if err != nil {
				// конец ввода
				if err != io.EOF {
					log.Printf("Read error: %s\n", err)
				}
				break
			}
			if remoteAddr == nil && remoteAddr != addr {
				remoteAddr = addr
				sync_channel <- remoteAddr
			}
			_, err = dst.Write(buf[0:nBytes])
			if err != nil {
				log.Fatalf("Write error: %s\n", err)
			}
		}
	}()
	log.Println("Exit write_from_udp_to_stream")
	return sync_channel
}

// Put input date from the stream to UDP connection
func put_from_stream_to_udp(src io.Reader, dst net.Conn, remoteAddr net.Addr) <-chan net.Addr {
	buf := make([]byte, 1024)
	sync_channel := make(chan net.Addr)
	go func() {
		for {
			var nBytes int
			var err error
			nBytes, err = src.Read(buf)
			if err != nil {
				// конец ввода
				if err != io.EOF {
					log.Printf("Read error: %s\n", err)
				}
				break
			}
			log.Println("Write to the remote address:", remoteAddr)
			if con, ok := dst.(*net.UDPConn); ok && remoteAddr != nil {
				_, err = con.WriteTo(buf[0:nBytes], remoteAddr)
			}
			if err != nil {
				log.Fatalf("Write error: %s\n", err)
			}
		}
	}()
	return sync_channel
}
