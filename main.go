package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"os"
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
	if len(os.Args) == 2 {
		file, err := os.ReadFile(os.Args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		ri(string(file))
	}
	//antlr4Main()
	//newMain()
	//simpleAdd()
	//oldMain()
}

func ri(code string) {
	is := antlr.NewInputStream(code)
	// 初始化相关内容
	lexer := parser.NewRiLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewRiParser(stream)
	p.BuildParseTrees = true
	tree := p.Prog()
	v := engine.NewRiVisitor()
	fmt.Println(code)
	result := v.Visit(tree)
	fmt.Println(result)
}
