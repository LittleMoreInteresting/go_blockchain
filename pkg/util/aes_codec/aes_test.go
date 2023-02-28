package aes_codec

import (
	"fmt"
	"testing"
)

func TestAESEncoding(t *testing.T) {
	encoding := AESEncoding("hallen")
	fmt.Println(encoding)
}
func TestAESDecoding(t *testing.T) {
	encoding := AESDecoding("mb1SXEc/MMCFBky0pt4Iog==")
	fmt.Println(encoding)
}
