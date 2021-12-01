package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/toffentoffen/advent-of-code-2021/utils"
	"io"
	"log"
	"strconv"
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
	fmt.Println(d.largerSlidingWindow())
}

type data struct{
	numbers []int
}

func (d data) largerSlidingWindow() int {
	var count = 0
	for i := 0; i < len(d.numbers)-3; i++ {

		if utils.IntSum(d.numbers[i:i+3]) < utils.IntSum(d.numbers[i+1:i+4]){
			count++
		}
	}
	return count
}

func readInput(in string) (*data, error) {
	return read(strings.NewReader(in))
}

func read(r io.Reader) (*data, error) {
	var d data

	s := bufio.NewScanner(r)
	for s.Scan() {
		x, err := strconv.Atoi(s.Text())
		if err != nil {
			return nil, err
		}
		d.numbers = append(d.numbers, x)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return &d, s.Err()
}

func testInputReader() io.Reader {
	return strings.NewReader(`199
200
208
210
200
207
240
269
260
263`)
}