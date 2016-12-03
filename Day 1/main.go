package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func parse(input string) (string, float64) {
	dir := string(input[0])
	fl, _ := strconv.Atoi(string(input[1:]))
	num := float64(fl)
	return dir, num
}

func main() {
	args := os.Args[1:]
	pos := make([]float64, 2)
	angle := math.Pi / 2
	for x := range args {
		dir, num := parse(strings.Trim(args[x], ","))
		switch dir {
		case "R":
			angle -= math.Pi / 2
		case "L":
			angle += math.Pi / 2
		}
		pos[0] += num * math.Cos(angle)
		pos[1] += num * math.Sin(angle)
	}
	x_val, _ := strconv.ParseFloat(strings.TrimPrefix(fmt.Sprintf("%.2f", pos[0]), "-"), 64)
	y_val, _ := strconv.ParseFloat(strings.TrimPrefix(fmt.Sprintf("%.2f", pos[1]), "-"), 64)
	fmt.Println(x_val + y_val)
}
