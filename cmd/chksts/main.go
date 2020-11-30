package main

import (
	"os"

	"github.com/rkorkosz/crwl/internal/checker"
)

func main() {
	ch := checker.New(nil, os.Stdin, os.Stdout)
	ch.Check()
}
