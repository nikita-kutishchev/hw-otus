package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	if len(str) == 0 {
		return "", nil
	}

	if unicode.IsDigit(rune(str[0])) {
		return "", ErrInvalidString
	}

	builder := strings.Builder{}
	builder.Grow(len(str))

	for i := 0; i < len(str); i++ {
		if i + 1 < len(str) && unicode.IsDigit(rune(str[i + 1])) {
			if unicode.IsDigit(rune(str[i + 2])) {
				return "", ErrInvalidString
			}

			num, _ := strconv.Atoi(string(str[i + 1]))
			builder.WriteString(strings.Repeat(string(str[i]), num))
			i++
		} else {
			builder.WriteByte(str[i])
		}
	}

	return builder.String(), nil
}
