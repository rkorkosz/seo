package checker

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

// Checker holds values needed to check the url
type Checker struct {
	client *http.Client
	target io.Writer
	source io.Reader
}

// New creates a new Checker
func New(client *http.Client, source io.Reader, target io.Writer) *Checker {
	if client == nil {
		client = http.DefaultClient
	}
	return &Checker{
		client: client,
		source: source,
		target: target,
	}
}

// Check performs the url status code checking
func (c *Checker) Check() {
	s := bufio.NewScanner(c.source)
	for s.Scan() {
		u := &url.URL{}
		ubytes := s.Bytes()
		err := u.UnmarshalBinary(ubytes)
		if err != nil {
			log.Println(err)
			continue
		}
		resp, err := c.client.Head(u.String())
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println(c.target)
		_, err = c.target.Write(ubytes)
		_, err = c.target.Write([]byte("\t"))
		_, err = c.target.Write([]byte(resp.Status))
		_, err = c.target.Write([]byte("\n"))
		if err != nil {
			log.Println(err)
			continue
		}
	}
	err := s.Err()
	if err != nil {
		panic(err)
	}
}
