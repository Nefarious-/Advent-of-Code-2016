package main

import (
	"fmt"
	"os"
)

type Keypad struct {
	Layout     [][]string
	Coords     []int
	Buttons    []string
	LastButton string
}

func (k *Keypad) Move(move string) {
	switch move {
	case "U":
		if k.Layout[k.Coords[0]-1][k.Coords[1]] != "0" {
			(*k).Coords[0] -= 1
			(*k).LastButton = k.Layout[k.Coords[0]][k.Coords[1]]
		}
	case "D":
		if k.Layout[k.Coords[0]+1][k.Coords[1]] != "0" {
			(*k).Coords[0] += 1
			(*k).LastButton = k.Layout[k.Coords[0]][k.Coords[1]]
		}
	case "L":
		if k.Layout[k.Coords[0]][k.Coords[1]-1] != "0" {
			(*k).Coords[1] -= 1
			(*k).LastButton = k.Layout[k.Coords[0]][k.Coords[1]]
		}
	case "R":
		if k.Layout[k.Coords[0]][k.Coords[1]+1] != "0" {
			(*k).Coords[1] += 1
			(*k).LastButton = k.Layout[k.Coords[0]][k.Coords[1]]
		}
	}
}

func (k *Keypad) parseInput(input []string) {
	for _, y := range input {
		for i := 0; i < len(y); i++ {
			k.Move(string(y[i]))
		}
		(*k).Buttons = append(k.Buttons, k.LastButton)
	}
}

func main() {
	args := os.Args[1:]
	k := Keypad{[][]string{[]string{"0", "0", "0", "0", "0", "0", "0"},
		[]string{"0", "0", "0", "1", "0", "0", "0"},
		[]string{"0", "0", "2", "3", "4", "0", "0"},
		[]string{"0", "5", "6", "7", "8", "9", "0"},
		[]string{"0", "0", "A", "B", "C", "0", "0"},
		[]string{"0", "0", "0", "D", "0", "0", "0"},
		[]string{"0", "0", "0", "0", "0", "0", "0"}},
		[]int{3, 1},
		make([]string, 0),
		"5"}
	k.parseInput(args)
	fmt.Println(k.Buttons)
}
