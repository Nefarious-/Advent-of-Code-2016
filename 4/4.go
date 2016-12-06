package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool { return s[i] < s[j] }
func (s sortRunes) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortRunes) Len() int           { return len(s) }

type Room struct {
	Code, Checksum, Name string
	ID                   int
}

func (r *Room) Split() {
	s := strings.Split(r.Code, "[")
	idPos := len(s[0]) - 3
	r.Name = (s[0])[:idPos-1]
	r.ID, _ = strconv.Atoi((s[0])[idPos:])
	r.Checksum = (s[1])[:5]
}

func StripHyphens(s string) string {
	return strings.Replace(s, "-", "", -1)
}

func (r *Room) IsValid() bool {
	t := []rune(StripHyphens(r.Name))
	sort.Sort(sortRunes(t))
	return Checksum(string(t)) == r.Checksum
}

func Checksum(s string) string {
	c := []string{string(s[0])}
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			c = append(c, string(s[i]))
		}
	}
	var l []int
	for _, v := range c {
		l = append(l, strings.Count(s, v))
	}
	for x := 0; x < len(l); x++ {
		for y := 1; y < len(l)-x; y++ {
			if l[y] > l[y-1] {
				l[y], l[y-1] = l[y-1], l[y]
				c[y], c[y-1] = c[y-1], c[y]
			}
		}
	}
	return strings.Join(c[:5], "")
}

func cycle(s rune, i int) rune {
	switch s {
	case '-':
		return ' '
	default:
		n := s + rune(i%26)
		if n > 122 {
			n -= 26
		}
		return n
	}
}

func main() {
	var counter int
	f, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	for _, y := range strings.Split(string(f), "\n") {
		t := Room{Code: y}
		t.Split()
		if t.IsValid() {
			counter += t.ID
			var rl []rune
			for _, r := range t.Name {
				rl = append(rl, cycle(rune(r), t.ID))
			}
			s := string(rl)
			if strings.Contains(s, "northpole") {
				fmt.Printf("%s: %d\n", s, t.ID)
			}
		}
	}
	fmt.Println("number of valid rooms:", counter)
}
