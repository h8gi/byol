package main

import (
	"fmt"
	"io"
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
	kind  string
	text  string
	value Any
}

type Any interface{}

// SkipSpaces consume spaces
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

func (lx *Lexer) ReadNumber() (Token, error) {
	s, size, err := lx.ReadWhile(unicode.IsDigit)
	// EOF
	if err != nil && size == 0 {
		return Token{kind: "eof", text: "EOF"}, err
	}
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("not number: {%s}", s))
	}
	return Token{kind: "number", text: s, value: num}, nil
}

func (lx *Lexer) ReadSymbol() (Token, error) {
	s, size, err := lx.ReadWhile(func(r rune) bool {
		return !unicode.IsSpace(r) && !unicode.IsDigit(r) && r != '(' && r != ')'
	})
	// EOF
	if err != nil && size == 0 {
		return Token{kind: "eof", text: "EOF"}, err
	}
	return Token{kind: "symbol", text: s}, nil
}

// ReadToken return Token structure
func (lx *Lexer) ReadToken() (Token, error) {
	if err := lx.SkipSpaces(); err != nil {
		return Token{kind: "eof", text: "EOF"}, err
	}
	r, _, _ := lx.ReadRune()
	var token Token
	switch {
	case unicode.IsDigit(r):
		lx.UnreadRune()
		token, _ = lx.ReadNumber()
	case r == '(':
		token = Token{kind: "open", text: "("}
	case r == ')':
		token = Token{kind: "close", text: ")"}
	default:
		lx.UnreadRune()
		token, _ = lx.ReadSymbol()
	}
	return token, nil
}

func main() {
	lx := Lexer{strings.NewReader("(12 3012 31 21 3 foo a() foo")}
	for {
		token, err := lx.ReadToken()
		fmt.Println(token)
		// EOF
		if err != nil {
			break
		}
	}

}
