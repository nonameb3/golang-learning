package main

import (
	"io"
	"log"
	"os"
)

func main() {
	txtName := os.Args[1]

	file, err := os.Open(txtName)
	if err != nil {
		log.Fatal(err)
	}

	io.Copy(os.Stdout, file)
}
