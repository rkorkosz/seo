package main

import (
	"flag"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/rkorkosz/crwl/internal/crawler"
)

func main() {
	flag.Parse()
	client := &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 5 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 5 * time.Second,
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 100,
			IdleConnTimeout:     60 * time.Second,
		},
	}
	c := crawler.New(client, flag.Arg(0), os.Stdout)
	c.Crawl()
	if err := c.Err(); err != nil {
		panic(err)
	}
}
