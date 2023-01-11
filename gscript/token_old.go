package gscript

import (
	"bytes"
	"fmt"
	"regexp"
)

type tokenItem struct {
	Type  string `json:"type" simple:"sType"`
	Value string `json:"value" simple:"sValue"`
}

func Tokenizer(input string) []tokenItem {
	tokens := make([]tokenItem, 0)

	current := 0
	for current < len(input) {
		char := input[current]
		fmt.Printf("%q", char)
		if char == '(' {
			tokens = append(tokens, tokenItem{"paren", "("})
		}
		if char == ')' {
			tokens = append(tokens, tokenItem{"paren", ")"})
		}
		WHITESPACE := regexp.MustCompile(`\s`)
		if WHITESPACE.MatchString(string(char)) {
			current++
			continue
		}

		NUMBERS := regexp.MustCompile(`[0-9]`)
		if NUMBERS.MatchString(string(char)) {
			var numbersValue bytes.Buffer

			for NUMBERS.MatchString(string(char)) {
				numbersValue.WriteString(string(char))
				current++
				char = input[current]
			}
			tokens = append(tokens, tokenItem{"number", numbersValue.String()})
			continue
		}

		if string(char) == "\"" {
			var stringValue bytes.Buffer
			current++
			char = input[current]
			strChar := string(char)
			for strChar != "\"" {
				stringValue.WriteString(strChar)
				current++
				char = input[current]
				strChar = string(char)
			}

			tokens = append(tokens, tokenItem{"string", stringValue.String()})
			current++
			continue
		}

		LETTERS := regexp.MustCompile(`[a-z]`)
		if LETTERS.MatchString(string(char)) {
			var lettersValue bytes.Buffer

			for LETTERS.MatchString(string(char)) {
				lettersValue.WriteString(string(char))
				current++
				char = input[current]
			}
			tokens = append(tokens, tokenItem{"name", lettersValue.String()})
			continue
		}

		current++
	}

	return tokens
}

func walk(tokens []tokenItem, current int) (node, int) {
	token := tokens[current]
	if token.Type == "number" {
		current++
		return node{
			"NumberLiteral",
			token.Value,
			make([]interface{}, 0),
		}, current
	}
	if token.Type == "string" {
		current++
		return node{
			"StringLiteral",
			token.Value,
			make([]interface{}, 0),
		}, current
	}

	if token.Type == "paren" && token.Value == "(" {
		current++
		token = tokens[current]

		node := node{"CallExpression", token.Value, make([]interface{}, 0)}
		current++
		token = tokens[current]
		for (token.Type != "paren") ||
			(token.Type == "paren" && token.Value != ")") {
			node, current = walk(tokens, current)
			node.Params = append(node.Params, node)
			token = tokens[current]
		}
		current++
		return node, current
	}
	return node{}, current
}

func SimpleParser(tokens []tokenItem) ast {
	current := 0

	ast := ast{"Program", make([]interface{}, 0)}
	for current < len(tokens) {
		node, walkCurrent := walk(tokens, current)
		current = walkCurrent
		ast.Body = append(ast.Body, node)
	}
	return ast
}

type ast struct {
	Type string
	Body []interface{}
}

type node struct {
	Type   string
	Name   string
	Params []interface{}
}
