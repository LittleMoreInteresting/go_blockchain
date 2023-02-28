package main

import (
	"fmt"

	"go_blockchain/pkg/util/base58"
)

func main() {
	str := "hello golang"
	en := base58.Base58Encoding(str)
	fmt.Println(en)
	de := base58.Base58Decoding(en)
	fmt.Println(de)
}
