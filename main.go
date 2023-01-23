package main

import (
	"bufio"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"os"
	"ri/core"
	"ri/gscript"
	"ri/mathparser"
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
	code := `a=1+2+3+4+5+6+10*10
b=1231+1
`
	is := antlr.NewInputStream(code)
	// 初始化相关内容
	lexer := parser.NewRiLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewRiParser(stream)
	p.BuildParseTrees = true
	tree := p.Prog()
	v := core.NewRiVisitor()
	result := v.Visit(tree)
	// 当前问题
	// 1 当前返回出现问题是因为循环的第一个child ，并没有进行后续循环
	// 2 存在对应访问节点出现不符合预期现象。可能g4编写错误。或者对应访问编写错误
	fmt.Println(code, result)
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
