package main

import (
	"fmt"
	"strings"
	"log"
)


func tokenize(data string) (tokens []Token) {
	data = strings.Trim(data, " ")
	tokens = make([]Token, 0)

	for len(data) > 0 {
		matched := false
		for _, matchFunc := range matchFunctions {
			if ret, succ := matchFunc(data); succ {
				data = strings.Trim(data[len(ret.Value):], " ")
				tokens = append(tokens, ret)
				matched = true
				break;
			}
		}

		if !matched {
			log.Println("skipping unknown token: " + string(data[0]))
			data = data[1:]
		}
	}

	return tokens
}

func main() {
	lines := []string{
		"mov eax, ebx",
		"mov eax, [ebx]",
		"mov eax, 123",
		"mov eax, [123]",
		"mov eax, -123",
		"mov eax, [-123]",
		"mov [eax], 123",
		"mov [eax], -123",
		"mov [eax], ebx",
		"mov [123], 123",
		"mov [123], -123",
		"mov [123], ebx",
		"mov [-123], 123",
		"mov [-123], -123",
		"mov [-123], ebx",
		"push eax",
		"push [eax]",
		"push 123",
		"push -123",
		"push [123]",
		"push [-123]",
		"halt",
	}

	for _, val := range lines {
		tokens := tokenize(val)
		fmt.Println(RightPad(val, 20) ,":", tokens)
	}
}
