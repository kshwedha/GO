package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

type StringManipulator func(string) string

func ToLower(m StringManipulator) StringManipulator {
	return func(s string) string {
		lower := strings.ToLower(s)
		return m(lower)
	}
}

func ToBase64(m StringManipulator) StringManipulator {
	return func(s string) string {
		b64 := base64.StdEncoding.EncodeToString([]byte(s))
		return m(b64)
	}
}

func AppendDecorator(x string) func(m StringManipulator) StringManipulator {
	return func(m StringManipulator) StringManipulator {
		return func(s string) string {
			return m(s + x)
		}
	}
}

func PrependDecorator(x string, m StringManipulator) StringManipulator {
	return func(s string) string {
		return m(x + s)
	}
}

// "identity" just return the same string
func ident(s string) string {
	return s
}

func main() {
	s := "Hello, playground"

	var fn StringManipulator = ident
	fmt.Println(fn(s))

	fn = ToBase64(ToLower(fn))
	fmt.Println(fn(s))

	var fn2 StringManipulator = ident
	fn2 = ToLower(ToBase64(fn2))
	fmt.Println(fn2(s))

	var fn3 StringManipulator = ident
	fn3 = AppendDecorator(" GOLANG")(ToLower(PrependDecorator("DECORATED ", fn3)))
	fmt.Println(fn3(s))
}
