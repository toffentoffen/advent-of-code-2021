package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/toffentoffen/advent-of-code-2021/utils"
	"io"
	"log"
	"strings"
)

var (
	//go:embed input.txt
	input string
)

func main() {
	d, err := readInput(input)
	//d, err := read(testInputReader())
	if err != nil {
		log.Fatal(err)
	}
	gamma, epsilon := d.result()
	fmt.Println(gamma * epsilon)
}

type bits struct {
	zeros int
	ones  int
}
type data struct {
	positions []bits
}

func (d data) result() (gamma int, epsilon int) {
	for i, position := range d.positions {

		if position.ones > position.zeros {
			gamma += 1<<(len(d.positions)-i-1)
		} else {
			epsilon += 1<<(len(d.positions)-i-1)
		}
	}
	return gamma, epsilon
}

func readInput(input string) (*data, error) {
	return read(strings.NewReader(input))
}

func read(r io.Reader) (*data, error) {
	var d data

	s := bufio.NewScanner(r)
	for s.Scan() {
		t := s.Text()
		digits := utils.Digits(t)
		if len(d.positions) == 0 {
			d.positions = make([]bits, len(digits))
		}
		for i, digit := range digits {
			if digit == 0 {
				d.positions[i].zeros++
			} else {
				d.positions[i].ones++
			}
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return &d, s.Err()
}

func testInputReader() io.Reader {
	return strings.NewReader(`00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`)
}
