package main

import (
	"fmt"
	"ri/mathparser"
	"strconv"
)

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
