package main

import (
	fset "cryptochallenges/firstSet"
	"fmt"
)

func main() {
	toConvert := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	shouldProduce := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	hexBytes, err := fset.FromHex([]byte(toConvert))
	if err != nil {
		panic(err)
	}
	base64Bytes, err := fset.ToBase64(hexBytes)
	if err != nil {
		panic(err)
	}
	stringed := string(base64Bytes)
	fmt.Println(stringed)

	fmt.Println("Does it coincide?")
	if stringed == shouldProduce {
		fmt.Println("YES!")
	} else {
		fmt.Println("NO...")
	}
}
