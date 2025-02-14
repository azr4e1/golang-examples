package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	encoded1 := encode('😊', []byte("ciao come va?"))
	encoded2 := encode('t', []byte("tutto bene tu?"))
	encoded := encoded1 + encoded2
	fmt.Printf("%s\n", encoded)
	decoded := decode(encoded)
	fmt.Printf("%s\n", decoded)
}

func byteToVariationSelector(b byte) rune {
	var r rune
	if b < 16 {
		r = rune(0xFE00 + uint32(b))
	} else {
		r = rune(0xE0100 + uint32(b-16))
	}

	return r
}

func encode(base rune, sentence []byte) string {
	s := new(strings.Builder)
	s.WriteRune(base)
	for _, b := range sentence {
		s.WriteRune(byteToVariationSelector(b))
	}

	return s.String()
}

func variationSelectorToByte(vs rune) (byte, error) {
	varSel := uint32(vs)
	var range1S, range1E uint32 = 0xFE00, 0xFE0F
	for i := range1S; i <= range1E; i++ {
		if varSel == i {
			return byte(varSel - range1S), nil
		}
	}
	var range2S, range2E uint32 = 0xE0100, 0xE01EF
	for i := range2S; i <= range2E; i++ {
		if varSel == i {
			return byte(varSel - range2S + 16), nil
		}
	}
	return 0, errors.New("couldn't decode")
}

func decode(varSels string) string {
	message := new(strings.Builder)
	for _, vs := range varSels {
		b, err := variationSelectorToByte(vs)
		if err == nil {
			message.WriteByte(b)
		} else {
			message.WriteByte('\n')
		}
	}
	return message.String()
}
