// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Ri

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseRiListener is a complete listener for a parse tree produced by RiParser.
type BaseRiListener struct{}

var _ RiListener = &BaseRiListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRiListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRiListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRiListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRiListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterNumber is called when production Number is entered.
func (s *BaseRiListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production Number is exited.
func (s *BaseRiListener) ExitNumber(ctx *NumberContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BaseRiListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BaseRiListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseRiListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseRiListener) ExitAddSub(ctx *AddSubContext) {}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BaseRiListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BaseRiListener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterBlockStatement is called when production blockStatement is entered.
func (s *BaseRiListener) EnterBlockStatement(ctx *BlockStatementContext) {}

// ExitBlockStatement is called when production blockStatement is exited.
func (s *BaseRiListener) ExitBlockStatement(ctx *BlockStatementContext) {}

// EnterBlockStms is called when production BlockStms is entered.
func (s *BaseRiListener) EnterBlockStms(ctx *BlockStmsContext) {}

// ExitBlockStms is called when production BlockStms is exited.
func (s *BaseRiListener) ExitBlockStms(ctx *BlockStmsContext) {}

// EnterStart is called when production start is entered.
func (s *BaseRiListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseRiListener) ExitStart(ctx *StartContext) {}
