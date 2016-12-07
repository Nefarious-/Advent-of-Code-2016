package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func isABBA(s [4]byte) bool {
	return s[0] == s[3] && s[1] == s[2] && s[0] != s[1]
}

func isOB(s string) bool { return s == "[" }

func containsABBA(s string) bool {
	if len(s) >= 4 {
		for i := 0; i < len(s)-3; i++ {
			if isABBA([4]byte{s[i], s[i+1], s[i+2], s[i+3]}) {
				return true
			}
		}
	}
	return false
}

func walk(s string) int {
	var res []string
	var last int
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "[" {
			cb := strings.Index(s, "]")
			bl := s[i : cb+1]
			s = strings.Replace(s, bl, "", -1)
			res = append(res, s[last:i])
			last = i
			i--
			if containsABBA(bl[1 : len(bl)-1]) {
				return 0
			}
		}
	}
	res = append(res, s[last:])
	for x := range res {
		if containsABBA(res[x]) {
			return 1
		}
	}
	return 0
}

func main() {
	f, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(f), "\n")
	var counter int
	for _, i := range input {
		counter += walk(i)
	}
	fmt.Println(counter)
}
