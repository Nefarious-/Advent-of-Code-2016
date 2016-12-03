package main

import (
	"fmt"
	"os"
)

type Keypad struct {
	lastButton int
	Buttons    []int
}

func (k Keypad) isValid(move string) bool {
	t := k.CalcMove(move)
	if (move == "U" && t < 1) || (move == "D" && t > 9) || (move == "L" && (t == 0 || t == 3 || t == 6)) || (move == "R" && (t == 7 || t == 4 || t == 10)) {
		return false
	}
	return true
}

func (k *Keypad) Move(move string) {
	(*k).lastButton = k.CalcMove(move)
}

func (k Keypad) CalcMove(move string) int {
	var res int
	switch move {
	case "U":
		res = -3
	case "D":
		res = 3
	case "L":
		res = -1
	case "R":
		res = 1
	}
	return k.lastButton + res
}

func (k *Keypad) parseInput(input []string) {
	for _, y := range input {
		for i := 0; i < len(y); i++ {
			move := string(y[i])
			if k.isValid(move) {
				k.Move(move)
			}
		}
		(*k).Buttons = append(k.Buttons, k.lastButton)
	}
}

func main() {
	args := os.Args[1:]
	k := Keypad{5, make([]int, 0)}
	k.parseInput(args)
	fmt.Println(k.Buttons)
}
