package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	var counter int
	args, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	triangles := strings.Split(string(args), "\n")
	for _, y := range triangles {
		num1, _ := strconv.Atoi(strings.TrimSpace(y[2:5]))
		num2, _ := strconv.Atoi(strings.TrimSpace(y[7:10]))
		num3, _ := strconv.Atoi(strings.TrimSpace(y[12:15]))
		if num1+num2 > num3 && num1+num3 > num2 && num2+num3 > num1 {
			counter++
		}
	}
	fmt.Println(counter)
}
