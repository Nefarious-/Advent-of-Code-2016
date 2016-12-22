package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"hash"
	"os"
)

func testHash(s string, h *hash.Hash) (bool, byte) {
	(*h).Write([]byte(s))
	res := hex.EncodeToString((*h).Sum(nil))
	(*h).Reset()
	return res[:5] == "00000", res[5]
}

func main() {
	input := os.Args[1]
	var pass []byte
	h := md5.New()
	for i, counter := 0, 0; counter < 8; i++ {
		if ok, res := testHash(fmt.Sprintf("%s%d", input, i), &h); ok {
			pass = append(pass, res)
			counter++
		}
	}
	fmt.Println(string(pass))
}
