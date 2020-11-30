package checker

import (
	"bufio"
	"io"
	"net/http"
	"net/url"
)

type Checker struct {
	client *http.Client
	target io.Writer
	source io.Reader
}

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

func (c *Checker) Check() {
	s := bufio.NewScanner(c.source)
	for s.Scan() {
		u := &url.URL{}
		err := u.UnmarshalBinary(s.Bytes())
		if err != nil {
			continue
		}
		resp, err := c.client.Head(u.String())
		if err != nil {
			continue
		}
		c.target.Write(s.Bytes())
		c.target.Write([]byte("\t"))
		c.target.Write([]byte(resp.Status))
		c.target.Write([]byte("\n"))
	}
	err := s.Err()
	if err != nil {
		panic(err)
	}
}
