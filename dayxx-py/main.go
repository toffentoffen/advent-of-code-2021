package main

import (
	"bufio"
	_ "embed"
	"io"
	"log"
	"strings"
)

var (
	//go:embed input.txt
	input string
)

func main() {
	//d, err := readInput(input)
	_, err := read(testInputReader())
	if err != nil {
		log.Fatal(err)
	}
}

type data struct{}

func readInput(input string) (*data, error) {
	return read(strings.NewReader(input))
}

func read(r io.Reader) (*data, error) {
	var d data

	s := bufio.NewScanner(r)
	for s.Scan() {
		// do stuff
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return &d, s.Err()
}

func testInputReader() io.Reader {
	return strings.NewReader(``)
}
