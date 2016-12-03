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

func findDouble(locs [][]float64) []float64 {
	for i := 0; i < len(locs); i++ {
		for j := i + 1; j < len(locs); j++ {
			if locs[i][0] == locs[j][0] && locs[i][1] == locs[j][1] {
				return locs[i]
			}
		}
	}
	return nil
}

func fmtVector(v []float64, stripNeg bool) {
	for x := range v {
		n := fmt.Sprintf("%.0f", v[x])
		if stripNeg {
			v[x], _ = strconv.ParseFloat(strings.TrimPrefix(n, "-"), 64)
		} else {
			v[x], _ = strconv.ParseFloat(n, 64)
		}
	}
}

func main() {
	args := os.Args[1:]
	angle := math.Pi / 2
	pos := make([]float64, 2)
	var prevLocs [][]float64
	for x := range args {
		dir, num := parse(strings.Trim(args[x], ","))
		switch dir {
		case "R":
			angle -= math.Pi / 2
		case "L":
			angle += math.Pi / 2
		}
		for y := 0; float64(y) < num; y++ {
			t := []float64{pos[0] + math.Cos(angle), pos[1] + math.Sin(angle)}
			fmtVector(t, false)
			prevLocs = append(prevLocs, t)
			pos = t
		}
	}
	fmt.Printf("Part 1: %.0f; Part 2: %.0f.", part1(prevLocs), part2(prevLocs))
}

func part1(v [][]float64) float64 {
	pos := v[len(v)-1]
	fmtVector(pos, true)
	return pos[0] + pos[1]
}

func part2(v [][]float64) float64 {
	vec := findDouble(v)
	fmtVector(vec, true)
	return vec[0] + vec[1]
}
