// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Ri

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type BaseRiVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseRiVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiVisitor) VisitMulDiv(ctx *MulDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiVisitor) VisitAddSub(ctx *AddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiVisitor) VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiVisitor) VisitBlockStatement(ctx *BlockStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiVisitor) VisitBlockStms(ctx *BlockStmsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}
