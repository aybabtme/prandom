package main

import (
	"github.com/dustin/randbo"
	"io"
	"log"
	"os"
)

func main() {
	log.SetPrefix("prandom: ")
	_, err := io.Copy(os.Stdout, randbo.New())
	if err != nil {
		log.Fatal(err)
	}
}
