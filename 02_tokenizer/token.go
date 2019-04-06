package main

import "regexp"

type SearchFunction func (data string) (ret Token, succ bool)
type TokenType int
type Token struct {
	Type TokenType
	Value string
}

func (t TokenType) String() (s string) {
	return tokens[t]
}

const (
	IDENT TokenType = iota
	NUMBER
	LBRACK
	RBRACK
	COMMA
	MINUS
	PLUS
)

var tokens = []string {
	IDENT: "IDENT",
	NUMBER: "NUMBER",
	LBRACK: "LBRACK",
	RBRACK: "RBRACK",
	COMMA: "COMMA",
	MINUS: "MINUS",
	PLUS: "PLUS",
}

var terminals = map[byte]TokenType {
	'[': LBRACK,
	']': RBRACK,
	',': COMMA,
	'-': MINUS,
	'+': PLUS,
}

var (
	RegexIdentifier = regexp.MustCompile("^[a-zA-Z]+")
	RegexNumber = regexp.MustCompile("^[1-9][0-9]*")
)

var matchFunctions = []SearchFunction{
	identifier,
	number,
	singleChar,
}