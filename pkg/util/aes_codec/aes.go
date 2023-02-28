package aes_codec

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
)

var KEY = []byte("hallenhallenhall")

func AESEncoding(s string) string {
	src := []byte(s)
	block, err := aes.NewCipher(KEY)
	if err != nil {
		return ""
	}
	src = padding(src, block.BlockSize())
	dst := make([]byte, len(src))
	block.Encrypt(dst, src)
	return base64.StdEncoding.EncodeToString(dst)
}

func AESDecoding(s string) string {
	src, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	block, err := aes.NewCipher(KEY)
	if err != nil {
		return ""
	}
	dst := make([]byte, len(src))
	block.Decrypt(dst, src)
	dst = un_padding(dst)
	return string(dst)
}

func padding(src []byte, block_size int) []byte {
	pad := block_size - len(src)%block_size
	repeat := bytes.Repeat([]byte{byte(pad)}, pad)
	return append(src, repeat...)
}

func un_padding(src []byte) []byte {
	l := len(src)
	if l <= 0 {
		panic("error len")
	}
	pad := int(src[l-1])
	return src[:l-pad]
}
