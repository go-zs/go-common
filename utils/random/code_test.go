package random

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestRandomCode(t *testing.T) {
	for i:=0;i<100;i++ {
		code := RandomCode()
		_, err := strconv.Atoi(code)
		assert.Nil(t, err)
	}
}

func TestRandomString(t *testing.T) {
	var (
		testCases = []struct{
			name string
			n  int
		}{
			{
				name: "",
				n:    5,
			},
			{
				name: "",
				n:    1,
			},
			{
				name: "",
				n:    10,
			},
		}
	)
	for _, tc := range testCases {
		for i:=0;i<100;i++ {
			str := RandomString(tc.n)
			assert.Equal(t, tc.n, len(str))
		}
	}
}

func BenchmarkRandomCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomCode()
	}
}

func BenchmarkRandomString_5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomString(10)
	}
}

func BenchmarkRandomString_10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomString(10)
	}
}