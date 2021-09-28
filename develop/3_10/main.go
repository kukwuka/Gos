package main

import (
	"bufio"
	"errors"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

var (
	timeout time.Duration
	url     string
)

func init() {
	flag.DurationVar(&timeout, "timeout", time.Second*10, "connection timeout")
	flag.StringVar(&url, "url", "google.com:80", "url")
}

func main() {
	flag.Parse()

	c := NewClient(url, timeout, ioutil.NopCloser(os.Stdin), os.Stdout)
	log.Printf("Connected to %s\n", url)
	if err := c.Connect(); err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := c.Close()
		log.Fatal(err)
	}()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			if err := c.Send(); err != nil {
				log.Fatal(err)
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			if err := c.Receive(); err != nil {
				log.Fatal(err)
				return
			}
		}
	}()

	wg.Wait()
}

var (
	ErrConnectionClosed = errors.New("connection unable")
	ErrEOF              = errors.New("EOF")
)

type Client interface {
	Connect() error
	Close() error
	Send() error
	Receive() error
}

type telnet struct {
	address    string
	timeout    time.Duration
	reader     io.ReadCloser
	writer     io.Writer
	conn       net.Conn
	readerScan *bufio.Scanner
	connScan   *bufio.Scanner
}

func NewClient(address string, timeout time.Duration, reader io.ReadCloser, writer io.Writer) Client {
	return &telnet{
		address: address,
		timeout: timeout,
		reader:  reader,
		writer:  writer,
	}
}
func (t *telnet) Connect() (err error) {

	t.conn, err = net.DialTimeout("tcp", t.address, t.timeout)
	t.connScan = bufio.NewScanner(t.conn)
	t.readerScan = bufio.NewScanner(t.reader)

	return
}

func (t *telnet) Close() (err error) {
	if t.conn != nil {
		err = t.conn.Close()
	}
	return
}

func (t *telnet) Send() (err error) {
	if t.conn == nil {
		return
	}
	if !t.readerScan.Scan() {
		return ErrEOF
	}
	_, err = t.conn.Write(append(t.readerScan.Bytes(), '\n'))
	return
}

func (t *telnet) Receive() (err error) {
	if t.conn == nil {
		return
	}
	if !t.connScan.Scan() {
		return ErrConnectionClosed
	}
	_, err = t.writer.Write(append(t.connScan.Bytes(), '\n'))
	return
}
