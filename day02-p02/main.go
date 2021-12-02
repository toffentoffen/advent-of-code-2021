package main

import (
	"bufio"
	_ "embed"
	"fmt"
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
	var sub submarine

	for _, move := range d.moves {
		sub.command(move)
	}
	result := sub.p.horizontal * sub.p.depth

	fmt.Println(result)
}

type data struct{
	moves []moverFunc
}
type position struct {
	horizontal int
	depth int
	aim int
}

type submarine struct {
	p position
}

func (s *submarine) command(move moverFunc) position {
	s.p = move(s.p)
	return s.p
}

type moverFunc func(p position) position

func forward(x int) moverFunc {
	return func(p position) position {
		p.horizontal += x
		p.depth += p.aim * x
		return p
	}
}

func up(x int) moverFunc {
	return func(p position) position {
		p.aim -= x
		return p
	}
}

func down(x int) moverFunc {
	return func(p position) position {
		p.aim += x
		return p
	}
}

func readInput(input string) (*data, error) {
	return read(strings.NewReader(input))
}

func read(r io.Reader) (*data, error) {
	var d data

	s := bufio.NewScanner(r)
	for s.Scan() {
		text := s.Text()
		var command string
		var n int
		if _, err := fmt.Sscanf(text, "%s %d", &command, &n); err != nil {
			panic(err)
		}

		var move moverFunc
		switch command {
		case "forward":
			move = forward(n)
		case "up":
			move = up(n)
		case "down":
			move = down(n)
		default:
			panic(fmt.Errorf("invalid command %q", command))
		}
		d.moves = append(d.moves, move)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return &d, s.Err()
}

func testInputReader() io.Reader {
	return strings.NewReader(`forward 5
down 5
forward 8
up 3
down 8
forward 2`)
}
