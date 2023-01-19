package core

import (
	"fmt"
	"ri/parser"
	"strconv"
)

type RiListener struct {
	*parser.BaseRiListener

	stack []int
}

func (l *RiListener) Push(i int) {
	l.stack = append(l.stack, i)
}

func (l *RiListener) Pop() int {
	if len(l.stack) < 1 {
		panic("stack is empty unable to Pop")
	}

	// Get the last value from the stack.
	result := l.stack[len(l.stack)-1]

	// Remove the last element from the stack.
	l.stack = l.stack[:len(l.stack)-1]

	return result
}

func (l *RiListener) ExitMulDiv(c *parser.MulDivContext) {
	right, left := l.Pop(), l.Pop()

	switch c.GetOp().GetTokenType() {
	case parser.RiParserMUL:
		l.Push(left * right)
	case parser.RiParserDIV:
		l.Push(left / right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
}

func (l *RiListener) ExitAddSub(c *parser.AddSubContext) {
	right, left := l.Pop(), l.Pop()

	switch c.GetOp().GetTokenType() {
	case parser.RiParserADD:
		l.Push(left + right)
	case parser.RiParserSUB:
		l.Push(left - right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
}

func (l *RiListener) ExitNumber(c *parser.NumberContext) {
	i, err := strconv.Atoi(c.GetText())
	if err != nil {
		panic(err.Error())
	}

	l.Push(i)
}
