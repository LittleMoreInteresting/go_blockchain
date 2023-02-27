package base58

import (
	"math/big"
)

var CHAR_58 = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'J', 'K', 'L', 'M', 'N', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

func Base58Encoding(str string) string {
	//1 转 ascii 码值
	str_byte := []byte(str)

	//2 转10进制
	i := big.NewInt(0).SetBytes(str_byte)
	var mods []byte
	//3 取余 求字符
	for i.Cmp(big.NewInt(0)) > 0 {
		mod := big.NewInt(0)
		i58 := big.NewInt(58)
		i.DivMod(i, i58, mod)
		mods = append(mods, CHAR_58[mod.Int64()])
	}
	//把0使用字节1代替
	for _, s := range mods {
		if s != '0' {
			break
		}
		mods = append(mods, '1')
	}
	reverse(mods)
	return string(mods)
}

func reverse(b []byte) {
	i, j := 0, len(b)-1
	for i < j {
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
}

func Base58Decoding(str string) string {

	return ""
}
