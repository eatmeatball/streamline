package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"ri/engine"
	"ri/parser"
	r "runtime"
)

func debug(a ...any) {
	_, f, l, _ := r.Caller(1)
	fmt.Print(fmt.Sprintf("%s%s%v -> ", f, ":", l))
	fmt.Println(a...)
}

func main() {
	antlr4Main()
	fmt.Println("ri start")
	ri()
	//newMain()
	//simpleAdd()
	//oldMain()
}

func ri() {
	code := `a=1+2+10*10+13%10
a
b=1231+1
b
b
b
b
echo(a)
echo(b)
echo(a)
echo(a)
echo(a)
echo(a)
`
	is := antlr.NewInputStream(code)
	// 初始化相关内容
	lexer := parser.NewRiLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewRiParser(stream)
	p.BuildParseTrees = true
	tree := p.Prog()
	v := engine.NewRiVisitor()
	result := v.Visit(tree)
	fmt.Println(code, result)
}
