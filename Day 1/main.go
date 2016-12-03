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

func fmtVecString(input string) [][]float64 {
	vec := strings.Split(input, " ")
	output := make([][]float64, 0)
	for i := 0; i < len(vec)-1; i++ {
		x_pos := strings.Index(vec[i], "X")
		y_pos := strings.Index(vec[i], "Y")
		x, _ := strconv.ParseFloat(vec[i][x_pos+1:y_pos], 64)
		y, _ := strconv.ParseFloat(vec[i][y_pos+1:], 64)
		output = append(output, []float64{x, y})
	}
	return output
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
	pos := make([]float64, 2)
	angle := math.Pi / 2
	var prev_locs string
	for x := range args {
		dir, num := parse(strings.Trim(args[x], ","))
		switch dir {
		case "R":
			angle -= math.Pi / 2
		case "L":
			angle += math.Pi / 2
		}
		for y := 0; float64(y) < num; y++ {
			pos[0] += math.Cos(angle)
			pos[1] += math.Sin(angle)
			fmtVector(pos, false)
			prev_locs += fmt.Sprintf("X%.0fY%.0f ", pos[0], pos[1])
		}
	}
	vec_fmt := fmtVecString(prev_locs)
	fmt.Printf("Part 1: %.0f; Part 2: %.0f.", part1(vec_fmt), part2(vec_fmt))
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
