package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func testHash(s string) (bool, byte, int) {
	h := md5.New()
	h.Write([]byte(s))
	res := hex.EncodeToString(h.Sum(nil))
	pos, err := strconv.ParseInt(string(res[5]), 10, 64)
	if err != nil {
		return false, 0, 0
	}
	return strings.HasPrefix(res, "00000"), res[6], int(pos)
}

func main() {
	input := os.Args[1]
	pass := make([]byte, 8)
	var counter int
	for i := 0; counter < 8; i++ {
		ok, v, pos := testHash(fmt.Sprintf("%s%d", input, i))
		if ok && pos < 8 && pos >= 0 {
			if pass[pos] == 0 {
				pass[pos] = v
				counter++
				fmt.Println(string(pass))
			}
		}
	}
}
