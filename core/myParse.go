package core

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/spf13/cast"
	"ri/parser"
)

type RiVisitor struct {
	parser.BaseRiVisitor
	memory map[string]int
}

func NewRiVisitor() *RiVisitor {
	return &RiVisitor{
		memory: map[string]int{},
	}
}

func (v *RiVisitor) Visit(tree antlr.ParseTree) interface{} {
	return v.VisitTree(tree)
}
func (v *RiVisitor) VisitTree(tree antlr.Tree) any {
	switch ctx := tree.(type) {
	case *parser.ProgContext:
		return v.VisitProg(ctx)
	case *parser.IntContext:
		return v.VisitInt(ctx)
	case *parser.AddSubContext:
		return v.VisitAddSub(ctx)
	case *parser.MulDivContext:
		return v.VisitMulDiv(ctx)
	case *parser.AssginContext:
		return v.VisitAssgin(ctx)
	case *parser.PrintExprContext:
		return v.VisitPrintExpr(ctx)
	case *parser.BlankContext:
		return v.VisitBlank(ctx)
	case *parser.ParensContext:
		return v.VisitParens(ctx)
	case *parser.IdContext:
		return v.VisitId(ctx)
	default:
		panic("Unknown context")
	}
}

func (v *RiVisitor) VisitProg(ctx *parser.ProgContext) interface{} {
	for _, child := range ctx.GetChildren() {
		return v.VisitTree(child)
	}
	return nil
}

func (v *RiVisitor) VisitPrintExpr(ctx *parser.PrintExprContext) interface{} {
	value := v.Visit(ctx.Expr())
	fmt.Println(value)
	return v.VisitChildren(ctx)
}

func (v *RiVisitor) VisitAssgin(ctx *parser.AssginContext) interface{} {
	id := ctx.ID().GetText()
	value, _ := v.Visit(ctx.Expr()).(int)
	v.memory[id] = value
	return value
}

func (v *RiVisitor) VisitBlank(ctx *parser.BlankContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *RiVisitor) VisitParens(ctx *parser.ParensContext) interface{} {
	return v.Visit(ctx.Expr())
}

func (v *RiVisitor) VisitMulDiv(ctx *parser.MulDivContext) interface{} {
	left := cast.ToInt(v.Visit(ctx.Expr(0)))
	right := cast.ToInt(v.Visit(ctx.Expr(0)))
	if ctx.GetOp().GetTokenType() == parser.RiParserMUL {
		return left * right
	} else {
		return left / right
	}
}

func (v *RiVisitor) VisitAddSub(ctx *parser.AddSubContext) interface{} {
	left := cast.ToInt(v.Visit(ctx.Expr(0)))
	right := cast.ToInt(v.Visit(ctx.Expr(1)))
	if ctx.GetOp().GetTokenType() == parser.RiParserADD {
		return left + right
	} else {
		return left - right
	}
}

func (v *RiVisitor) VisitId(ctx *parser.IdContext) interface{} {
	id := ctx.ID().GetText()
	data, _ := v.memory[id]
	return data
}

func (v *RiVisitor) VisitInt(ctx *parser.IntContext) interface{} {
	return cast.ToInt(ctx.INT().GetText())
}
