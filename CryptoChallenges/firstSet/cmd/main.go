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
	fmt.Println("Original:", shouldProduce)
	stringed := string(base64Bytes)
	fmt.Println("Result  :", stringed)

	fmt.Print("Does it coincide? ")
	if stringed == shouldProduce {
		fmt.Println("YES!")
	} else {
		fmt.Println("NO...")
	}
	// b := byte(0b10101111)
	// hx := fset.ToHex([]byte{b})
	// fmt.Println(string(hx))
}