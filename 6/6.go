package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type column []string

func (r column) Less(a, b int) bool { return r[a] < r[b] }
func (r column) Swap(a, b int)      { r[a], r[b] = r[b], r[a] }
func (r column) Len() int           { return len(r) }

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	var l [9][]string
	for _, s := range strings.Split(string(input), "\n") {
		for x, y := range s {
			l[x] = append(l[x], string(y))
		}
	}
	var chars []string
	for _, i := range l {
		sort.Sort(column(i))
		ch := i[0]
		tot := strings.Count(strings.Join(i, ""), i[0])
		for c := 1; c < len(i); c++ {
			count := strings.Count(strings.Join(i, ""), i[c])
			if count > tot { //modify to < for part 2
				ch = i[c]
				tot = count
			}
		}
		chars = append(chars, ch)
	}
	fmt.Println(strings.Join(chars, ""))
}
