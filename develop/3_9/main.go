package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"

	"net/http"
	"time"
)

const (
	linkpattern = `href="http*"`
)

func main() {
	url := flag.String("url", "http://example.com", "url")
	timeout := flag.Duration("timeout", 7*time.Second, "example")
	output_path := flag.String("output", "test", "output")
	r := flag.Bool("r", false, "recursive")

	flag.Parse()

	content, err := HTTPGet(*url, *timeout)

	if err != nil {
		log.Fatalln("HTTPGET: ", err)
	}
	if *r {
		contentAll, names := HTTPGetAll(content, *timeout)
		names = append([]string{(*url)[7:]}, names...)
		os.Mkdir(*output_path, 0777)
		for i := 0; i < len(contentAll); i++ {
			err = ioutil.WriteFile(*output_path+"/"+names[i], contentAll[i], 0666)
		}
	} else {
		err = ioutil.WriteFile(*output_path, content, 0666)
	}
	if err != nil {
		log.Fatalln("WriteFile: ", err)
	}
}

func HTTPGet(url string, timeout time.Duration) (content []byte, err error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	ctx, cancel_func := context.WithTimeout(context.Background(), timeout)
	defer cancel_func()
	request = request.WithContext(ctx)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("INVALID RESPONSE; status: %s", response.Status)
	}

	return ioutil.ReadAll(response.Body)
}

func HTTPGetAll(page []byte, timeout time.Duration) (rez [][]byte, names []string) {
	re, _ := regexp.Compile(linkpattern)
	rez = append(rez, page)

	links := re.FindAll(page, -1)

	for i := 0; i < len(links); i++ {
		way := string(links[i][5 : len(links[i])-1])
		fmt.Println(way)
		buf, err := HTTPGet(way, timeout)
		if err == nil {
			rez = append(rez, buf)
			names = append(names, way)
		}
	}

	return
}
