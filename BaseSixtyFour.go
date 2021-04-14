package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	s := "Hello base64"

	sEnc := encodeStd(s)
	fmt.Println(sEnc)

	fmt.Println(decodeStd(sEnc))

	uEnc := encodeURL(s)
	fmt.Println(uEnc)

	fmt.Println(decodeURL(uEnc))
}

func encodeStd(s string) string {
	return b64.StdEncoding.EncodeToString([]byte(s))
}

func decodeStd(s string) string {
	res, err := b64.StdEncoding.DecodeString(s)
	if err != nil {
		fmt.Println(err)
	}
	return string(res)
}

func encodeURL(s string) string {
	return b64.URLEncoding.EncodeToString([]byte(s))
}

func decodeURL(s string) string {
	res, err := b64.URLEncoding.DecodeString(s)
	if err != nil {
		fmt.Println(err)
	}
	return string(res)
}
