package main

import (
	"bufio"
	fset "cryptochallenges/firstSet"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Challenge 1")
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

	fmt.Println("\nChallenge 2")
	input1 := "1c0111001f010100061a024b53535009181c"
	input2 := "686974207468652062756c6c277320657965"
	expected := "746865206b696420646f6e277420706c6179"
	hx1, err := fset.FromHex([]byte(input1))
	if err != nil {
		panic(err)
	}
	hx2, err := fset.FromHex([]byte(input2))
	if err != nil {
		panic(err)
	}
	result := fset.XOR(hx1, hx2)
	hexedResult := fset.ToHex(result)

	fmt.Println("Result:", string(hexedResult))
	fmt.Print("Does it coincide? ")
	if expected == string(hexedResult) {
		fmt.Println("YES!")
	} else {
		fmt.Println("NO...")
	}

	fmt.Println("\nChallenge 3")
	message := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	hx, err := fset.FromHex([]byte(message))
	if err != nil {
		panic(err)
	}

	decrypted, key := fset.FrequencyXORCypher(hx)
	if key != nil {
		fmt.Println("Key: ", string(key[0]))
		fmt.Println("Decrypted message: ", string(decrypted))
	}

	fmt.Println("\nChallenge 4")
	f, err := os.Open("../data/4.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	lineNr := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		hx, err = fset.FromHex(line)
		if err != nil {
			panic(err)
		}
		decrypted, key = fset.FrequencyXORCypher(hx)
		if key != nil {
			fmt.Println("Key: ", string(key[0]))
			fmt.Println("Decrypted message: ", string(decrypted))
			break
		}
		lineNr++
	}
}
