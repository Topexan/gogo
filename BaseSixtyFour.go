package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	s := "Hello base64"

	sEnc := b64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println(sEnc)

	sDec, err := b64.StdEncoding.DecodeString(sEnc)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(sDec))

	uEnc := b64.URLEncoding.EncodeToString([]byte(s))
	fmt.Println(uEnc)
	
	uDec, err := b64.URLEncoding.DecodeString(uEnc)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(uDec))
}