package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Pixel string

type Screen struct {
	values [][]Pixel
	x, y   int
}

func (p Pixel) isOn() bool {
	return p == "#"
}

func (s Screen) Print() {
	for x := range s.values {
		fmt.Println(s.values[x])
	}
}

func (s *Screen) startup() {
	for a := 0; a < s.y; a++ {
		var x []Pixel
		for b := 0; b < s.x; b++ {
			x = append(x, ".")
		}
		s.values = append(s.values, x)
	}
}

func (s *Screen) rotateY(a, b int) {
	t := make([]Pixel, s.y)
	for n, v := range s.values {
		f := n + b%s.y
		if f >= s.y {
			f -= s.y
		}
		t[f] = v[a]
	}
	for n, x := range s.values {
		x[a] = t[n]
	}
}
func (s *Screen) rotateX(a, b int) {
	t := make([]Pixel, s.x)
	for x, v := range s.values[a] {
		f := x + b%s.x
		if f >= s.x {
			f -= s.x
		}
		t[f] = v
	}
	s.values[a] = t
}
func (s *Screen) turnOn(a, b int) {
	for x := 0; x < b; x++ {
		for y := 0; y < a; y++ {
			s.values[x][y] = "#"
		}
	}
}

func main() {
	sc := Screen{x: 50, y: 6}
	sc.startup()
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	for _, x := range strings.Split(string(input), "\r\n") {
		s := strings.Split(x, " ")
		if strings.HasPrefix(x, "rotate") {
			i, _ := strconv.Atoi(string(s[2][2:]))
			j, _ := strconv.Atoi(string(s[4][:]))
			if s[1] == "column" {
				sc.rotateY(i, j)
			} else {
				sc.rotateX(i, j)
			}
		} else {
			xPos := strings.Index(s[1], "x")
			i, _ := strconv.Atoi(string(s[1][:xPos]))
			j, _ := strconv.Atoi(string(s[1][xPos+1:]))
			sc.turnOn(i, j)
		}
	}
	var counter int
	for _, v := range sc.values {
		for _, w := range v {
			if w.isOn() {
				counter++
			}
		}
	}
	fmt.Println(counter)
	sc.Print()
}
