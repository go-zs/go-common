package random

import (
	"math/rand"
	"time"
)

var (
	Digits = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	LowerCaseChars = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	UpperCaseChars = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	Chars = append(Digits, append(LowerCaseChars, UpperCaseChars...)...)
)

// 验证随机手机验证码
func RandomCode() string {
	rand.Seed(time.Now().UnixNano())
	code := make([]byte, 6)
	for i := 0; i < 6; i++ {
		c := rand.Intn(10) + 48
		code[i] = uint8(c)
	}
	return string(code)
}

// 随机字符串
func RandomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	str := make([]byte, n)
	for i := 0; i < n; i++ {
		idx := rand.Intn(62)
		str[i] = Chars[idx]
	}

	return string(str)
}

