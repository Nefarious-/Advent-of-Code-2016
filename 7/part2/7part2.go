package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func isRev(b string) bool {
	if len(b) == 3 {
		return b[0] == b[2] && b[0] != b[1]
	}
	return false
}

func check(s1, s2 []string) bool {
	for i := range s1 {
		for j := range s2 {
			if isRev(s1[i]) && isRev(s2[j]) && s1[i][1:] == s2[j][:2] {
				return true
			}
		}
	}
	return false
}

func split(s []string) []string {
	var res []string
	for _, b := range s {
		if len(b) >= 3 {
			for i := 0; i < len(b)-2; i++ {
				res = append(res, b[i:i+3])
			}
		}
	}
	return res
}

func walk(s string) bool {
	var brs []string
	var res []string
	var last int
	for a := 0; a < len(s); a++ {
		if string(s[a]) == "[" {
			ob := strings.Index(s[a:], "]")
			inBr := s[a : ob+1+a]
			brs = append(brs, inBr[1:len(inBr)-1])
			s = strings.Replace(s, inBr, "", -1)
			res = append(res, s[last:a])
			last = a
			a--
		}
	}
	res = append(res, s[last:])
	return check(split(res), split(brs))
}

func main() {
	f, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	var counter int
	for _, i := range strings.Split(string(f), "\n") {
		if walk(i) {
			counter += 1
		}
	}
	fmt.Println(counter)
}
