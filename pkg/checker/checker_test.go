package checker

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestChecker(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`OK`))
	}))
	t.Cleanup(server.Close)

	target := &bytes.Buffer{}
	source := strings.NewReader("https://example.com\nhttps://example1.com\n")

	c := New(server.Client(), source, target)
	c.Check()
	expected := "https://example.com\t200 OK\nhttps://example1.com\t200 OK\n"
	if strings.Compare(expected, target.String()) != 0 {
		t.Errorf("Expected %s, got %s", expected, target)
	}
}
