package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"hash"
	"os"
	"strconv"
)

func testHash(s string, h *hash.Hash) (bool, byte, int64) {
	(*h).Write([]byte(s))
	res := hex.EncodeToString((*h).Sum(nil))
	(*h).Reset()
	if pos, err := strconv.ParseInt(string(res[5]), 10, 64); err == nil {
		return res[:5] == "00000", res[6], pos
	}
	return false, 0, 0
}

func main() {
	input := os.Args[1]
	pass := make([]byte, 8)
	h := md5.New()
	for i, counter := 0, 0; counter < 8; i++ {
		if ok, v, pos := testHash(fmt.Sprintf("%s%d", input, i), &h); ok {
			if pos < 8 && pass[pos] == 0 {
				pass[pos] = v
				counter++
			}
		}
	}
	fmt.Println(string(pass))
}
