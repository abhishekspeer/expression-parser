package main

import "unicode/utf8"

func identifier(data string) (ret Token, succ bool) {
	if str := RegexIdentifier.FindString(data); str != "" {
		return Token{Type: IDENT, Value: str}, true
	}
	return Token{}, false
}

func number(data string) (ret Token, succ bool) {
	if str := RegexNumber.FindString(data); str != "" {
		return Token{Type: NUMBER, Value: str}, true
	}
	return Token{}, false
}

func singleChar(data string) (ret Token, succ bool) {
	if val, ok := terminals[data[0]]; ok {
		return Token{Type: val, Value: string(data[0])}, true
	}
	return Token{}, false
}

func RightPad(str string, pad int) (ret string) {
	len := utf8.RuneCountInString(str)

	if len < pad {
		padding := ""

		for i := len; i <= pad; i++ {
			padding += " "
		}

		return str + padding
	} else {
		return str
	}
}