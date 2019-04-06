package main

import (
	"fmt"
)

type token struct {
	Name  string
	Value string
}

type tree struct {
	Left  *tree
	Name  string
	Value string
	Right *tree
}

func isoperand(in string) (ret bool) {
	if len(in) > 0 {
		if in[0] == '+' || in[0] == '-' {
			return true
		}
	}

	return false
}

func isdigit(in string) (ret bool) {
	if len(in) > 0 {
		if in[0] >= '0' && in[0] <= '9' {
			return true
		}
	}

	return false
}

func tokenize(in string) (tokens []token) {
	tokens = make([]token, 0)
	remaining := in

	for remaining != "" {
		if isdigit(remaining) {
			number := ""
			for isdigit(remaining) {
				number += string(remaining[0])
				remaining = remaining[1:]
			}
			tokens = append(tokens, token{Name: "number", Value: number})
		} else if isoperand(remaining) {
			tok := string(remaining[0])

			remaining = remaining[1:]
			name := ""

			if tok == "+" {
				name = "plus"
			} else if tok == "-" {
				name = "minus"
			}

			tokens = append(tokens, token{Name: name, Value: tok})
		} else {
			tok := string(remaining[0])
			remaining = remaining[1:]
			tokens = append(tokens, token{Name: "unknown", Value: tok})
		}
	}

	return tokens
}

func parse(tokens []token) (tr *tree) {
	tr = &tree{Name: "statement", Left: expression(tokens)}

	return tr
}

func term(tokens []token) (tr *tree) {
	return &tree{Name: "number", Value: tokens[0].Value}
}

func expression(tokens []token) (tr *tree) {
	left := term(tokens)

	tokens = tokens[1:]

	if len(tokens) >= 2 {
		if (tokens[0].Name == "plus"  || tokens[0].Name == "minus") {
			return &tree{Name: "expression", Left: left, Value: tokens[0].Name, Right: expression(tokens[1:])}
		}
	}

	/*	if len(tokens) >= 3 {
		if tokens[0].Name == "number" && (tokens[1].Name == "plus"  || tokens[1].Name == "minus") {
			tr.Left = &tree{Name: tokens[0].Name, Value: tokens[0].Value}
			tr.Name = tokens[1].Name
			tr.Right = expression(tokens[2:])
		}
	}*/

	return &tree{Name: "expression", Left: left}
}

func pad(d int) (r string) {
	r = ""

	for i := 0; i < d; i++ {
		r += "    "
	}

	return r
}

func print(tr *tree, d int) {
	fmt.Println(pad(d) + tr.Name, "(")

	if tr.Left != nil {
		print(tr.Left, d + 1)
	}

	if tr.Value != "" {
		fmt.Println(pad(d + 1) + tr.Value)
	}

	if tr.Right != nil {
		print(tr.Right, d + 1)
	}

	fmt.Println(pad(d) + ")")
}

func main() {
	tok := tokenize("1-2+5-10")

	tree := parse(tok)

	print(tree, 0)
}
