package main

import (
	"os"

	"github.com/rkorkosz/seo/pkg/checker"
)

func main() {
	ch := checker.New(nil, os.Stdin, os.Stdout)
	ch.Check()
}
