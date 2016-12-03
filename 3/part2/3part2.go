package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func fmtStr(s string) int {
	s = strings.TrimSpace(s)
	i, _ := strconv.Atoi(s)
	return i
}

func main() {
	args, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	triangles := strings.Split(string(args), "\n")
	var c1, c2, c3 []int
	for _, y := range triangles {
		c1 = append(c1, fmtStr(y[2:5]))
		c2 = append(c2, fmtStr(y[7:10]))
		c3 = append(c3, fmtStr(y[12:]))
	}
	c1 = append(c1, c2...)
	c1 = append(c1, c3...)
	var set [][]int
	for x := 2; x <= len(c1); x++ {
		if (x+1)%3 == 0 {
			set = append(set, []int{c1[x-2], c1[x-1], c1[x]})
		}
	}
	var counter int
	for _, i := range set {
		if i[0]+i[1] > i[2] && i[0]+i[2] > i[1] && i[1]+i[2] > i[0] {
			counter++
		}
	}
	fmt.Println(counter)
}
