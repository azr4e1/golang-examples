package encoji

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	StatusOK int = iota
	MissingInputError
	TooManyInputsError
	MissingTextToEncodeError
	ExecutionError
)

type option func(*smuggler) error

type smuggler struct {
	stdin     io.Reader
	stdout    io.Writer
	stderr    io.Writer
	clearText string
	encode    bool
}

func (s smuggler) EncodeText(target string) (string, error) {
	if len(target) == 0 {
		return "", errors.New("target cannot be empty")
	}
	if len(s.clearText) == 0 {
		return "", errors.New("clear text cannot be empty")
	}
	encodedText := encode(rune(target[0]), []byte(s.clearText))
	if len(target) > 1 {
		encodedText += target[1:]
	}

	return encodedText, nil
}

func (s smuggler) DecodeText(target string) (string, error) {
	if len(target) == 0 {
		return "", errors.New("target cannot be empty")
	}
	return strings.TrimSpace(decode(target)), nil
}

func (s smuggler) Run() error {
	scan := bufio.NewScanner(s.stdin)
	for scan.Scan() {
		target := scan.Text()
		var res string
		var err error
		if s.encode {
			res, err = s.EncodeText(target)
		} else {
			res, err = s.DecodeText(target)
		}
		if err != nil {
			return err
		}
		fmt.Fprintln(s.stdout, res)
	}
	return nil
}

func (s *smuggler) SetEncode() {
	s.encode = true
}

func (s *smuggler) SetDecode() {
	s.encode = false
}

func (s *smuggler) SetClearText(text string) {
	s.clearText = text
}

func (s *smuggler) SetIn(stdin io.Reader) {
	s.stdin = stdin
}

func (s *smuggler) SetOut(stdout io.Writer) {
	s.stdout = stdout
}

func (s *smuggler) SetErr(stderr io.Writer) {
	s.stderr = stderr
}

func NewSmuggler(opts ...option) (*smuggler, error) {
	s := &smuggler{
		stdin:     os.Stdin,
		stdout:    os.Stdout,
		stderr:    os.Stderr,
		clearText: "",
		encode:    false,
	}

	for _, o := range opts {
		err := o(s)
		if err != nil {
			return s, err
		}
	}
	return s, nil
}

func WithInput(stdin io.Reader) option {
	return func(s *smuggler) error {
		s.stdin = stdin
		return nil
	}
}

func WithOutput(stdout io.Writer) option {
	return func(s *smuggler) error {
		s.stdout = stdout
		return nil
	}
}

func WithError(stderr io.Writer) option {
	return func(s *smuggler) error {
		s.stderr = stderr
		return nil
	}
}

func WithClearText(text string) option {
	return func(s *smuggler) error {
		s.clearText = text
		return nil
	}
}

func WithClearFile(file string) option {
	return func(s *smuggler) error {
		f, err := os.Open(file)
		if err != nil {
			return err
		}
		text, err := io.ReadAll(f)
		if err != nil {
			return err
		}
		s.clearText = string(text)
		return nil
	}
}

func WithEncodeFlag(f bool) option {
	return func(s *smuggler) error {
		s.encode = f
		return nil
	}
}

func Main(stdin, stdout, stderr io.ReadWriter) int {
	flag.Usage = func() {
		fmt.Fprintf(stderr, "Usage: %s [-encode string | -encodefile filepath | -decode] [stdin]\n", os.Args[0])
		fmt.Fprintln(stderr, "Encode/decode text using unicode variation selectors")
		fmt.Fprintln(stderr, "Flags:")
		flag.PrintDefaults()
	}
	encodeMode := flag.String("encode", "", "smuggle data within provided text")
	encodeFile := flag.String("encodefile", "", "smuggle data from file within provided text")
	decodeMode := flag.Bool("decode", false, "decode smuggled data")
	flag.Parse()

	if !(*decodeMode) && *encodeMode == "" && *encodeFile == "" {
		flag.Usage()
		return MissingInputError
	}

	var input io.Reader
	if arg := flag.Args(); len(arg) < 1 {
		input = stdin
	} else if len(arg) > 1 {
		flag.Usage()
		return TooManyInputsError
	} else {
		input = bytes.NewBuffer([]byte(arg[0]))
	}

	var encodeOption = WithClearText(*encodeMode)
	if *encodeMode == "" && *encodeFile != "" {
		encodeOption = WithClearFile(*encodeFile)
	}

	s, err := NewSmuggler(WithInput(input), WithOutput(stdout), WithError(stderr), WithEncodeFlag(!(*decodeMode)), encodeOption)
	if err != nil {
		fmt.Fprintln(stderr, err)
		return ExecutionError
	}

	err = s.Run()
	if err != nil {
		fmt.Fprintln(stderr, err)
		return ExecutionError
	}
	return StatusOK
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
