package base58

import (
	"fmt"
	"testing"
)

func TestBase58Encoding(t *testing.T) {
	encoding := Base58Encoding("he")
	fmt.Println(encoding)
}
func TestBase58Decoding(t *testing.T) {
	encoding := Base58Decoding("he")
	fmt.Println(encoding)
}
