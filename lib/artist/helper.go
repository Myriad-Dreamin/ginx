package artist

import (
	"bytes"
	"reflect"
	"strings"
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
	return getReflectElements(reflect.ValueOf(i))
}

func getReflectElements(v reflect.Value) (reflect.Value, reflect.Type) {
	t := v.Type()
	for t.Kind() == reflect.Ptr {
		v, t = v.Elem(), t.Elem()
	}
	return v, t
}

func getElementValue(i interface{}) reflect.Value {
	v, _ := getElements(i)
	return v
}

func getElementType(i interface{}) reflect.Type {
	_, t := getElements(i)
	return t
}

func getReflectElementType(v reflect.Value) reflect.Type {
	_, t := getReflectElements(v)
	return t
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

func allBig(name string) bool {
	for i := range name {
		if unicode.IsLower(rune(name[i])) {
			return false
		}
	}
	return true
}

func fromBigCamelToSnake(name string) string {
	// ID, etc...
	if allBig(name) {
		return strings.ToLower(name)
	}

	if len(name) == 0 {
		return ""
	}
	var b = bytes.NewBuffer(make([]byte, 0, len(name)))
	b.WriteByte(byte(unicode.ToLower(rune(name[0]))))
	name = name[1:]
	for i := range name {
		if unicode.IsUpper(rune(name[i])) {
			b.WriteByte('_')
			b.WriteByte(byte(unicode.ToLower(rune(name[i]))))
		} else {
			b.WriteByte(name[i])
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
