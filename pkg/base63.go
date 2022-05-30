package pkg

import (
	"strings"
)

const (
	characters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"
	length     = int64(len(characters))
)

func EncodeBase63(n int64) (string, error) {
	if n < 0 {
		return "", ErrInvalidInput
	}

	if n == 0 {
		return string(characters[0]), nil
	}
	var s string

	for ; n > 0; n = n / length {
		s = string(characters[n%length]) + s
	}

	s = strings.Repeat("0", 10-len(s)) + s
	return s, nil
}

func DecodeBase63(shortenedURL string) (int64, error) {
	shortenedURL = strings.TrimLeft(shortenedURL, "0")

	var n int64
	for _, c := range []byte(shortenedURL) {
		i := strings.IndexByte(characters, c)
		if i < 0 {
			return 0, ErrInvalidInput
		}
		n = length*n + int64(i)
	}

	return n, nil
}
