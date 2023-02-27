package main

import (
	"fmt"

	"go_blockchain/pkg/util/base64"
)

func main() {
	str := "hello golang"
	en := base64.Base64Encoding(str)
	fmt.Println(en)
	de := base64.Base64Decoding(en)
	fmt.Println(de)
}
