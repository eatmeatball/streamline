// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Ri

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// A complete Visitor for a parse tree produced by RiParser.
type RiVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by RiParser#Number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by RiParser#MulDiv.
	VisitMulDiv(ctx *MulDivContext) interface{}

	// Visit a parse tree produced by RiParser#AddSub.
	VisitAddSub(ctx *AddSubContext) interface{}

	// Visit a parse tree produced by RiParser#functionDeclaration.
	VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{}

	// Visit a parse tree produced by RiParser#blockStatement.
	VisitBlockStatement(ctx *BlockStatementContext) interface{}

	// Visit a parse tree produced by RiParser#BlockStms.
	VisitBlockStms(ctx *BlockStmsContext) interface{}

	// Visit a parse tree produced by RiParser#start.
	VisitStart(ctx *StartContext) interface{}
}
