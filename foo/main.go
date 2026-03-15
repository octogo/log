package main

import (
	"io"

	"github.com/octogo/log/v2"
)

func main() {
	log.Fatalf("Error: %s", io.EOF)
}
