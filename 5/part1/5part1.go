package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func testHash(s string) (bool, byte) {
	h := md5.New()
	h.Write([]byte(s))
	res := hex.EncodeToString(h.Sum(nil))
	return strings.HasPrefix(res, "00000"), res[5]
}

func main() {
	input := os.Args[1]
	var pass []byte
	var counter int
	for i := 0; counter < 8; i++ {
		ok, res := testHash(fmt.Sprintf("%s%d", input, i))
		if ok {
			pass = append(pass, res)
			counter++
		}
	}
	fmt.Println(string(pass))
}
