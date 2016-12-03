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

func fmtVector(v []float64, strip_neg bool) {
	for x := range v {
		if strip_neg {
			v[x], _ = strconv.ParseFloat(strings.TrimPrefix(fmt.Sprintf("%.0f", v[x]), "-"), 64)
		} else {
			v[x], _ = strconv.ParseFloat(fmt.Sprintf("%.0f", v[x]), 64)
		}
	}
}

func main() {
	args := os.Args[1:]
	angle := math.Pi / 2
	pos := make([]float64, 2)
	prev_locs := make([][]float64, 0)
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
			prev_locs = append(prev_locs, t)
			pos = t
		}
	}
	fmt.Printf("Part 1: %.0f; Part 2: %.0f.", part1(prev_locs), part2(prev_locs))
}

func part1(v [][]float64) float64 {
	final_pos := v[len(v)-1]
	fmtVector(final_pos, true)
	return final_pos[0] + final_pos[1]
}

func part2(v [][]float64) float64 {
	vec_ans := findDouble(v)
	fmtVector(vec_ans, true)
	return vec_ans[0] + vec_ans[1]
}
