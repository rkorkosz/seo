package crawler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCrawler(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`<html><a href="https://example.com/test">Test</a></html>`))
	}))
	t.Cleanup(server.Close)

	target := &bytes.Buffer{}

	c := New(server.Client(), server.URL, target)
	c.Crawl()
	expected := "https://example.com/test\n"
	if strings.Compare(expected, target.String()) != 0 {
		t.Errorf("Expected %s, got %s", expected, target)
	}
}
