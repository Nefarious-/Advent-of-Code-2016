package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	r := strings.NewReplacer(" ", "", "\n", "", "\r", "")
	i := r.Replace(string(input))
	for n := 0; n < len(i); n++ {
		if i[n] == '(' {
			pos := strings.Index(i[n:], ")") + n
			sl := i[n+1 : pos]
			cmd := strings.Split(sl, "x")
			p, _ := strconv.Atoi(cmd[0])
			x, _ := strconv.Atoi(cmd[1])
			pos += 1
			end := pos + p
			s := strings.Repeat(i[pos:end], x)
			var b bytes.Buffer
			b.WriteString(i[:n])
			b.WriteString(s)
			b.WriteString(i[end:])
			i = b.String()
			n += p*x - 1
		}
	}
	fmt.Println(len(i))
}
