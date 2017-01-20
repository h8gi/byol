package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Lexer
type Lexer struct {
	io.RuneScanner
}

// Token
type Token struct {
	kind  int
	text  string
	value Any
}

const (
	EOF = -(iota + 1)
	Error
	Comment
	Ident
	Boolean
	Number
	Char
	String
	Open
	Close
	OpenVec
	Quote
	QuasiQuote
	Unquote
	UnqoteAt
	Dot
)

var tokenstring = map[int]string{
	EOF:        "EOF",
	Comment:    "Comment",
	Error:      "Error",
	Ident:      "Ident",
	Number:     "Number",
	Char:       "Char",
	String:     "String",
	Open:       "Open",
	Close:      "Close",
	OpenVec:    "OpenVec",
	Quote:      "Quote",
	QuasiQuote: "QuasiQuote",
	Unquote:    "Unquote",
	UnqoteAt:   "UnqoteAt",
	Dot:        "Dot",
}

func (t Token) String() string {
	return fmt.Sprintf("%s: %s", tokenstring[t.kind], t.text)
}

type Any interface{}

// SkipSpaces consume spaces
// error is io.EOF
func (lx *Lexer) SkipSpaces() error {
	for {
		r, _, err := lx.ReadRune()
		// EOF check
		if err != nil {
			return err
		}
		if !unicode.IsSpace(r) {
			lx.UnreadRune()
			return nil
		}
	}
}

// Read while pred(r) return true
// error is io.EOF
func (lx *Lexer) ReadWhile(pred func(rune) bool) (s string, size int, err error) {
	rs := make([]rune, 0)
	for {
		r, _, eof := lx.ReadRune()
		// EOF check
		if eof != nil {
			return string(rs), len(rs), eof
		}
		// pred fail
		if !pred(r) {
			lx.UnreadRune()
			return string(rs), len(rs), nil
		}
		rs = append(rs, r)
	}
}

// Read while digit
// err is EOF
func (lx *Lexer) ReadNumber() (Token, error) {
	s, size, err := lx.ReadWhile(unicode.IsDigit)
	// EOF
	if err != nil && size == 0 {
		return Token{kind: EOF}, err
	}
	num, err := strconv.Atoi(s)
	if err != nil {
		return Token{kind: Error, text: s}, fmt.Errorf("lexer: Illegal Number", s)
	}
	return Token{kind: Number, text: s, value: num}, nil
}

func (lx *Lexer) ReadIdent(head rune) (Token, error) {
	s, _, err := lx.ReadWhile(IsIdentSucc)
	// EOF
	s = string(head) + s
	return Token{kind: Ident, text: s}, err
}

// Ident character
func IsIdentHead(r rune) bool {
	return strings.ContainsRune("!$%&*/:<=>?^_~", r) || unicode.IsLetter(r)
}
func IsIdentSucc(r rune) bool {
	return IsIdentHead(r) || unicode.IsDigit(r) || strings.ContainsRune("+-.@", r)
}

func (lx *Lexer) ReadDot() (Token, error) {
	s, size, err := lx.ReadWhile(func(r rune) bool { return r == '.' })
	if size == 2 {
		return Token{kind: Ident, text: "..."}, err
	} else if size == 0 {
		return Token{kind: Dot, text: "."}, err
	} else {
		err = fmt.Errorf("lexer: Illegal dot before %s", s)
		return Token{kind: Error}, err
	}
}

// ReadToken return Token structure
func (lx *Lexer) ReadToken() (Token, error) {

	var token Token
	var err error
	if err = lx.SkipSpaces(); err != nil {
		return Token{kind: EOF}, err
	}
	r, _, _ := lx.ReadRune()

	switch {
	case unicode.IsDigit(r):
		lx.UnreadRune()
		token, err = lx.ReadNumber()
	case IsIdentHead(r):
		token, err = lx.ReadIdent(r)
	case r == '.':
		token, err = lx.ReadDot()
	case r == '+' || r == '-':
		token = Token{kind: Ident, text: string(r)}
	case r == '(':
		token = Token{kind: Open, text: "("}
	case r == ')':
		token = Token{kind: Close, text: ")"}
	case r == '\'':
		token = Token{kind: Quote, text: "'"}
	case r == '`':
		token = Token{kind: QuasiQuote, text: "`"}
	default:
		token = Token{kind: Error, text: string(r)}
	}
	return token, err
}

func repl() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("lispy> ")
		line, _, _ := reader.ReadLine()
		lx := Lexer{strings.NewReader(string(line))}
		token, _ := lx.ReadToken()

		for {
			fmt.Println(token)
			// EOF
			if token.kind == Error {
				fmt.Fprintf(os.Stderr, "Illegal token: %s\n", token.text)
				break
			}
			if token.kind == EOF {
				break
			}
			token, _ = lx.ReadToken()
		}
		fmt.Println("")
	}
}

func main() {
	fp, err := os.Open("./test.scm")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(fp)
	lx := Lexer{reader}
	token, err := lx.ReadToken()
	for {
		fmt.Println(token)
		// EOF
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			break
		} else if token.kind == Error {
			fmt.Fprintf(os.Stderr, "Illegal token: %s\n", token.text)
			break
		}
		token, err = lx.ReadToken()
	}
}
