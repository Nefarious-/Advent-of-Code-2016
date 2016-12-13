package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func analyse(s string) (length int) {
	for n := 0; n < len(s); n++ {
		if s[n] == '(' {
			pos := strings.Index(s[n:], ")") + n
			sl := s[n+1 : pos]
			cmd := strings.Split(sl, "x")
			i, _ := strconv.Atoi(cmd[0])
			j, _ := strconv.Atoi(cmd[1])
			length += j*analyse(s[pos+1:pos+1+i]) - 1
			n = pos + i
		}
		length++
	}
	return
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	r := strings.NewReplacer(" ", "", "\n", "", "\r", "")
	s := r.Replace(string(input))
	fmt.Println(analyse(s))
}
