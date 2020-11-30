package crawler

import (
	"io"
	"net/http"
	"net/url"

	"golang.org/x/net/html"
)

// Crawler holds values needed to crawl a page
type Crawler struct {
	err    error
	target io.Writer
	url    *url.URL
	client *http.Client
}

// New creates a new Crawler
func New(client *http.Client, u string, out io.Writer) *Crawler {
	if client == nil {
		client = http.DefaultClient
	}
	c := Crawler{
		client: client,
		target: out,
	}
	c.url, c.err = url.Parse(u)
	return &c
}

// Crawl performs page crawling for urls
func (c *Crawler) Crawl() {
	resp, err := c.client.Get(c.url.String())
	if err != nil {
		c.err = err
		return
	}
	defer resp.Body.Close()
	c.extractLinks(resp.Body)
}

// Err returns a crawling error
func (c *Crawler) Err() error {
	return c.err
}

func (c *Crawler) clean(u []byte) []byte {
	parsed := &url.URL{}
	err := parsed.UnmarshalBinary(u)
	if err != nil {
		c.err = err
	}
	if parsed.Scheme == "" {
		parsed.Scheme = c.url.Scheme
	}
	if parsed.Host == "" {
		parsed.Host = c.url.Host
	}
	buf, err := parsed.MarshalBinary()
	if err != nil {
		c.err = err
	}
	return buf
}

func (c *Crawler) extractLinks(body io.Reader) {
	var err error
	z := html.NewTokenizer(body)
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return
		case html.StartTagToken:
			tn, _ := z.TagName()
			if len(tn) == 1 && tn[0] == 'a' {
				for {
					key, val, more := z.TagAttr()
					if !more {
						break
					}
					if string(key) == "href" {
						_, err = c.target.Write(c.clean(val))
						_, err = c.target.Write([]byte("\n"))
						if err != nil {
							c.err = err
						}
					}
				}
			}
		}
	}
}
