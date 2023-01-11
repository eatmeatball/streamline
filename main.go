package main

import (
	"bufio"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"os"
	"ri/gscript"
	"ri/mathparser"
	r "runtime"
	"strconv"
	"strings"
)

func debug(a ...any) {
	_, f, l, _ := r.Caller(1)
	fmt.Print(fmt.Sprintf("%s%s%v -> ", f, ":", l))
	fmt.Println(a...)
}

func main() {
	antlr4Main()
	//newMain()
	//simpleAdd()
	//oldMain()
}

/*
class {
func numberFunction() int {
return 1+2
}
}
*/

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

type HelloListener struct {
	*mathparser.BaseHelloListener

	stack []int
}

func (l *HelloListener) push(i int) {
	l.stack = append(l.stack, i)
}

func (l *HelloListener) pop() int {
	if len(l.stack) < 1 {
		panic("stack is empty unable to pop")
	}

	// Get the last value from the stack.
	result := l.stack[len(l.stack)-1]

	// Remove the last element from the stack.
	l.stack = l.stack[:len(l.stack)-1]

	return result
}

func (l *HelloListener) ExitMulDiv(c *mathparser.MulDivContext) {
	right, left := l.pop(), l.pop()

	switch c.GetOp().GetTokenType() {
	case mathparser.HelloParserMUL:
		l.push(left * right)
	case mathparser.HelloParserDIV:
		l.push(left / right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
}

func (l *HelloListener) ExitAddSub(c *mathparser.AddSubContext) {
	right, left := l.pop(), l.pop()

	switch c.GetOp().GetTokenType() {
	case mathparser.HelloParserADD:
		l.push(left + right)
	case mathparser.HelloParserSUB:
		l.push(left - right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
}

func (l *HelloListener) ExitNumber(c *mathparser.NumberContext) {
	i, err := strconv.Atoi(c.GetText())
	if err != nil {
		panic(err.Error())
	}

	l.push(i)
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
