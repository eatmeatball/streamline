package main

import (
	"bufio"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"os"
	"ri/gscript"
	"ri/mathparser"
	"strconv"
	"strings"
)

func antlr4Main() {
	// Setup the input
	is := antlr.NewInputStream("1+  1*1-1+2")

	// Create the Lexer
	lexer := mathparser.NewHelloLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := mathparser.NewHelloParser(stream)

	// Finally parse the expression (by walking the tree)
	var listener HelloListener
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())

	fmt.Println(listener.pop())
}

func oldMain() {
	// tokens := tokenizer("( add 1  ( subtract 4 100 ))")
	tokens := gscript.Tokenizer("( add 2 ( subtract 4 2))")
	fmt.Println()
	fmt.Println(tokens)
	fmt.Println(gscript.SimpleParser(tokens))
}

func newMain() {
	gscript.InitRuntime(false)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("gscript")
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		root, err := gscript.Parse(text)
		if err != nil {
			fmt.Println(err)
		} else {
			gscript.EvaluateWithRuntime(root, "")
		}
	}
}

func simpleAdd() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("ripl")
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		dataList := strings.Split(text, " ")
		if len(dataList) != 3 {
			fmt.Println("input : [num1] [op(+|-|*|/)] [num2] ")
		} else {
			num1, err := s2int64(dataList[0])
			if err != nil {
				fmt.Println(err)
				continue
			}
			op := dataList[1]
			num2, err := s2int64(dataList[0])
			if err != nil {
				fmt.Println(err)
				continue
			}
			switch op {
			case "+":
				fmt.Println(num1 + num2)
				break
			case "-":
				fmt.Println(num1 - num2)
				break
			case "*":
				fmt.Println(num1 * num2)
			case "/":
				fmt.Println(num1 / num2)
			}
		}
	}
}

func s2int64(s string) (int64, error) {
	v, err := strconv.ParseInt(trimZeroDecimal(s), 0, 0)
	if err == nil {
		return v, nil
	}
	return 0, fmt.Errorf("unable to cast %#v of type %T to int64", s, s)
}

func trimZeroDecimal(s string) string {
	var foundZero bool
	for i := len(s); i > 0; i-- {
		switch s[i-1] {
		case '.':
			if foundZero {
				return s[:i-1]
			}
		case '0':
			foundZero = true
		default:
			return s
		}
	}
	return s
}
