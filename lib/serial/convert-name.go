package serial

import (
	"bytes"
	"unicode"
)

func fromSnakeToCamel(src string, big bool) string {
	if len(src) == 0 {
		return ""
	}
	var b = bytes.NewBuffer(make([]byte, 0, len(src)))
	for i := range src {
		if src[i] == '_' {
			big = true
		} else {
			if big {
				big = false
				b.WriteByte(byte(unicode.ToUpper(rune(src[i]))))
			} else {
				b.WriteByte(src[i])
			}
		}
	}
	return b.String()
}

func fromSnakeToSmallCamel(src string) string {
	return fromSnakeToCamel(src, false)
}

func fromSnakeToBigCamel(src string) string {
	return fromSnakeToCamel(src, true)
}

