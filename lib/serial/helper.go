package serial

import (
	"bytes"
	"reflect"
	"unicode"
)

type packageSet = map[string]bool

func clonePackage(m packageSet) (n packageSet) {
	if m == nil {
		return nil
	}
	n = make(packageSet)
	for k, v := range m {
		n[k] = v
	}
	return n
}

func mergePackage(pac packageSet, oth packageSet) packageSet {
	newPac := make(packageSet)
	for k, v := range pac {
		newPac[k] = v
	}
	for k, v := range oth {
		newPac[k] = v
	}
	return newPac
}

func inplaceMergePackage(pac packageSet, oth packageSet) packageSet {
	if pac == nil {
		pac = make(packageSet)
	}
	for k, v := range oth {
		pac[k] = v
	}
	return pac
}

func getElements(i interface{}) (reflect.Value, reflect.Type) {
	v, t := reflect.ValueOf(i), reflect.TypeOf(i)
	for t.Kind() == reflect.Ptr {
		v, t = v.Elem(), t.Elem()
	}
	return v, t
}

func getElementValue(i interface{}) reflect.Value {
	v, _ := getElements(i)
	return v
}

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
