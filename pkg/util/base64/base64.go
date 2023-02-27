package base64

import "encoding/base64"

func Base64Encoding(str string) string {
	by := []byte(str)

	return base64.StdEncoding.EncodeToString(by)
}

func Base64Decoding(str string) string {
	res, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return ""
	}
	return string(res)
}
