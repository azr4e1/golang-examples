package main

import (
	"bufio"
	"bytes"
	fset "cryptochallenges/firstSet"
	"fmt"
	"io"
	"os"
	"slices"
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

	key := fset.FrequencyXORCypher(hx)
	if key != 0 {
		decrypted := fset.XOR(hx, bytes.Repeat([]byte{key}, len(hx)))
		fmt.Println("Key: ", string(key))
		fmt.Println("Decrypted message: ", string(decrypted))
	}

	fmt.Println("\nChallenge 4")
	f, err := os.Open("../data/4.txt")
	defer f.Close()
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
		key = fset.FrequencyXORCypher(hx)
		if key != 0 {
			decrypted := fset.XOR(hx, bytes.Repeat([]byte{key}, len(hx)))
			fmt.Println("Key: ", string(key))
			fmt.Println("Decrypted message: ", string(decrypted))
			break
		}
		lineNr++
	}

	fmt.Println("\nChallenge 5")
	// f, err = os.Open("./main.go")
	// if err != nil {
	// 	panic(err)
	// }
	// toEncrypt, err := io.ReadAll(f)
	// if err != nil {
	// 	panic(err)
	// }
	toEncrypt := `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`
	expected = "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	keyRept := []byte("ICE")
	encrypted := fset.RepeatingKeyXOREncrypt(keyRept, []byte(toEncrypt))

	hexedResult = fset.ToHex(encrypted)
	fmt.Println("Result:", string(hexedResult))
	fmt.Print("Does it coincide? ")
	if expected == string(hexedResult) {
		fmt.Println("YES!")
	} else {
		fmt.Println("NO...")
	}

	fmt.Println("\nChallenge 6")
	f, err = os.Open("../data/6.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	dataL := bytes.Split(data, []byte{'\n'})
	encrypted = []byte{}
	for _, line := range dataL {
		b64, err := fset.FromBase64(line)
		if err != nil {
			panic(err)
		}
		encrypted = append(encrypted, b64...)
	}
	keyLengths := fset.GetKeyLengths(encrypted)
	lengths := []int{}
	for k := range keyLengths {
		lengths = append(lengths, k)
	}
	slices.SortFunc(lengths, func(a, b int) int {
		if keyLengths[a] < keyLengths[b] {
			return -1
		}
		if keyLengths[a] > keyLengths[b] {
			return 1
		}
		return 0
	})
	topLengths := lengths[:5]
	blocks := fset.GetBlocks(encrypted, topLengths[4])
	fmt.Println(topLengths)
	keys := []byte{}
	for _, block := range blocks {
		key := fset.FrequencyXORCypher(block)
		keys = append(keys, key)
	}
	fmt.Println(keys)
	fmt.Println(string(fset.RepeatingKeyXOREncrypt([]byte{keys[0]}, blocks[0])))
}
