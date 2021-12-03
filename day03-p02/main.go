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
	oxygen, co2 := d.rating(func(bs bits) uint8 {
		if bs.ones >= bs.zeros {
			return '1'
		}
		return '0'
	}), d.rating(func(bs bits) uint8 {
		if bs.zeros <= bs.ones {
			return '0'
		}
		return '1'
	})
	fmt.Println(oxygen * co2)
}

type bits struct {
	zeros int
	ones  int
}
type data struct {
	numbers []string
}

func (d data) rating(keeper func(bits) uint8) (rating int64) {
	numOfDigits := len(d.numbers[0])
	actNumbers := d.numbers
	for i := 0; i < numOfDigits; i++ {
		var remainingNumbers []string
		bs := bitsCount(actNumbers, i)
		var keep = keeper(bs)
		for _, number := range actNumbers {
			if number[i] == keep {
				remainingNumbers = append(remainingNumbers, number)
			}
		}
		if len(remainingNumbers) == 1 {
			rating, _ = strconv.ParseInt(remainingNumbers[0], 2, 64)
			break
		}
		actNumbers = remainingNumbers
	}
	return rating
}

func readInput(input string) (*data, error) {
	return read(strings.NewReader(input))
}

func read(r io.Reader) (*data, error) {
	var d data

	s := bufio.NewScanner(r)
	for s.Scan() {
		t := s.Text()
		d.numbers = append(d.numbers, t)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return &d, s.Err()
}

func bitsCount(numbers []string, pos int) bits {
	var bs bits
	for _, num := range numbers {
		digits := utils.Digits(num)

		if digits[pos] == 0 {
			bs.zeros++
		} else {
			bs.ones++
		}
	}
	return bs
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
