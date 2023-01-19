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
	code := `func a() int {
return 1+1 +1
}`
	is := antlr.NewInputStream(code)
	// 初始化相关内容
	fmt.Println(is)
	lexer := parser.NewRiLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewRiParser(stream)
	var listener core.RiListener
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())
	fmt.Println(listener.Pop())
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
