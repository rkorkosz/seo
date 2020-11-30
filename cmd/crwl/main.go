package main

import (
	"flag"
	"os"

	"github.com/rkorkosz/seo/pkg/crawler"
)

func main() {
	flag.Parse()
	c := crawler.New(nil, flag.Arg(0), os.Stdout)
	c.Crawl()
	if err := c.Err(); err != nil {
		panic(err)
	}
}
