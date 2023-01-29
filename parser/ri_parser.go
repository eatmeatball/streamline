// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Ri

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type RiParser struct {
	*antlr.BaseParser
}

var riParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func riParserInit() {
	staticData := &riParserStaticData
	staticData.literalNames = []string{
		"", "'('", "')'", "'{'", "'}'", "'='", "';'", "'echo'", "'if'", "'for'",
		"'bool'", "'string'", "'float'", "'any'", "'func'", "'true'", "'false'",
		"", "", "", "", "", "'*'", "'/'", "'+'", "'-'", "'%'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "ECHO", "IF", "FOR", "BOOL", "STRING", "FLOAT",
		"ANY", "FUNC", "TRUE", "FALSE", "ID", "INT", "NEWLINE", "WS", "DECIMAL_FLOAT_LIT",
		"MUL", "DIV", "ADD", "SUB", "MOD",
	}
	staticData.ruleNames = []string{
		"expr", "parExpression", "ifstat", "stat", "prog",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 26, 84, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 19, 8, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 30, 8, 0, 10, 0,
		12, 0, 33, 9, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 4, 2, 40, 8, 2, 11, 2, 12,
		2, 41, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 4, 3, 70, 8, 3, 11, 3, 12, 3, 71, 1, 3, 1, 3, 1, 3, 3,
		3, 77, 8, 3, 1, 4, 4, 4, 80, 8, 4, 11, 4, 12, 4, 81, 1, 4, 0, 1, 0, 5,
		0, 2, 4, 6, 8, 0, 3, 1, 0, 15, 16, 1, 0, 22, 23, 1, 0, 24, 25, 92, 0, 18,
		1, 0, 0, 0, 2, 34, 1, 0, 0, 0, 4, 39, 1, 0, 0, 0, 6, 76, 1, 0, 0, 0, 8,
		79, 1, 0, 0, 0, 10, 11, 6, 0, -1, 0, 11, 19, 7, 0, 0, 0, 12, 19, 5, 18,
		0, 0, 13, 19, 5, 17, 0, 0, 14, 15, 5, 1, 0, 0, 15, 16, 3, 0, 0, 0, 16,
		17, 5, 2, 0, 0, 17, 19, 1, 0, 0, 0, 18, 10, 1, 0, 0, 0, 18, 12, 1, 0, 0,
		0, 18, 13, 1, 0, 0, 0, 18, 14, 1, 0, 0, 0, 19, 31, 1, 0, 0, 0, 20, 21,
		10, 7, 0, 0, 21, 22, 5, 26, 0, 0, 22, 30, 3, 0, 0, 8, 23, 24, 10, 6, 0,
		0, 24, 25, 7, 1, 0, 0, 25, 30, 3, 0, 0, 7, 26, 27, 10, 5, 0, 0, 27, 28,
		7, 2, 0, 0, 28, 30, 3, 0, 0, 6, 29, 20, 1, 0, 0, 0, 29, 23, 1, 0, 0, 0,
		29, 26, 1, 0, 0, 0, 30, 33, 1, 0, 0, 0, 31, 29, 1, 0, 0, 0, 31, 32, 1,
		0, 0, 0, 32, 1, 1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 34, 35, 5, 1, 0, 0, 35,
		36, 3, 0, 0, 0, 36, 37, 5, 2, 0, 0, 37, 3, 1, 0, 0, 0, 38, 40, 3, 6, 3,
		0, 39, 38, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0, 41, 42,
		1, 0, 0, 0, 42, 5, 1, 0, 0, 0, 43, 44, 5, 7, 0, 0, 44, 45, 5, 1, 0, 0,
		45, 46, 5, 17, 0, 0, 46, 47, 5, 2, 0, 0, 47, 77, 5, 19, 0, 0, 48, 49, 5,
		8, 0, 0, 49, 50, 3, 2, 1, 0, 50, 51, 5, 3, 0, 0, 51, 52, 3, 4, 2, 0, 52,
		53, 5, 4, 0, 0, 53, 77, 1, 0, 0, 0, 54, 55, 3, 0, 0, 0, 55, 56, 5, 19,
		0, 0, 56, 77, 1, 0, 0, 0, 57, 58, 5, 17, 0, 0, 58, 59, 5, 5, 0, 0, 59,
		60, 3, 0, 0, 0, 60, 61, 5, 19, 0, 0, 61, 77, 1, 0, 0, 0, 62, 63, 5, 9,
		0, 0, 63, 64, 5, 1, 0, 0, 64, 65, 5, 6, 0, 0, 65, 66, 5, 6, 0, 0, 66, 67,
		5, 2, 0, 0, 67, 69, 5, 3, 0, 0, 68, 70, 3, 6, 3, 0, 69, 68, 1, 0, 0, 0,
		70, 71, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 73, 1,
		0, 0, 0, 73, 74, 5, 4, 0, 0, 74, 77, 1, 0, 0, 0, 75, 77, 5, 19, 0, 0, 76,
		43, 1, 0, 0, 0, 76, 48, 1, 0, 0, 0, 76, 54, 1, 0, 0, 0, 76, 57, 1, 0, 0,
		0, 76, 62, 1, 0, 0, 0, 76, 75, 1, 0, 0, 0, 77, 7, 1, 0, 0, 0, 78, 80, 3,
		6, 3, 0, 79, 78, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 81, 79, 1, 0, 0, 0, 81,
		82, 1, 0, 0, 0, 82, 9, 1, 0, 0, 0, 7, 18, 29, 31, 41, 71, 76, 81,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// RiParserInit initializes any static state used to implement RiParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRiParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RiParserInit() {
	staticData := &riParserStaticData
	staticData.once.Do(riParserInit)
}

// NewRiParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRiParser(input antlr.TokenStream) *RiParser {
	RiParserInit()
	this := new(RiParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &riParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// RiParser tokens.
const (
	RiParserEOF               = antlr.TokenEOF
	RiParserT__0              = 1
	RiParserT__1              = 2
	RiParserT__2              = 3
	RiParserT__3              = 4
	RiParserT__4              = 5
	RiParserT__5              = 6
	RiParserECHO              = 7
	RiParserIF                = 8
	RiParserFOR               = 9
	RiParserBOOL              = 10
	RiParserSTRING            = 11
	RiParserFLOAT             = 12
	RiParserANY               = 13
	RiParserFUNC              = 14
	RiParserTRUE              = 15
	RiParserFALSE             = 16
	RiParserID                = 17
	RiParserINT               = 18
	RiParserNEWLINE           = 19
	RiParserWS                = 20
	RiParserDECIMAL_FLOAT_LIT = 21
	RiParserMUL               = 22
	RiParserDIV               = 23
	RiParserADD               = 24
	RiParserSUB               = 25
	RiParserMOD               = 26
)

// RiParser rules.
const (
	RiParserRULE_expr          = 0
	RiParserRULE_parExpression = 1
	RiParserRULE_ifstat        = 2
	RiParserRULE_stat          = 3
	RiParserRULE_prog          = 4
)

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RiParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RiParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParensContext struct {
	*ExprContext
}

func NewParensContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParensContext {
	var p = new(ParensContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ParensContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParensContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParensContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiVisitor:
		return t.VisitParens(s)

	default:
		return t.VisitChildren(s)
	}
}

type ModContext struct {
	*ExprContext
}

func NewModContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ModContext {
	var p = new(ModContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ModContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ModContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ModContext) MOD() antlr.TerminalNode {
	return s.GetToken(RiParserMOD, 0)
}

func (s *ModContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiVisitor:
		return t.VisitMod(s)

	default:
		return t.VisitChildren(s)
	}
}

type BoolContext struct {
	*ExprContext
	op antlr.Token
}

func NewBoolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolContext {
	var p = new(BoolContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *BoolContext) GetOp() antlr.Token { return s.op }

func (s *BoolContext) SetOp(v antlr.Token) { s.op = v }

func (s *BoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolContext) FALSE() antlr.TerminalNode {
	return s.GetToken(RiParserFALSE, 0)
}

func (s *BoolContext) TRUE() antlr.TerminalNode {
	return s.GetToken(RiParserTRUE, 0)
}

func (s *BoolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiVisitor:
		return t.VisitBool(s)

	default:
		return t.VisitChildren(s)
	}
}

type MulDivContext struct {
	*ExprContext
	op antlr.Token
}

func NewMulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivContext {
	var p = new(MulDivContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MulDivContext) GetOp() antlr.Token { return s.op }

func (s *MulDivContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *MulDivContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MulDivContext) MUL() antlr.TerminalNode {
	return s.GetToken(RiParserMUL, 0)
}

func (s *MulDivContext) DIV() antlr.TerminalNode {
	return s.GetToken(RiParserDIV, 0)
}

func (s *MulDivContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiVisitor:
		return t.VisitMulDiv(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddSubContext struct {
	*ExprContext
	op antlr.Token
}

func NewAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubContext {
	var p = new(AddSubContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AddSubContext) GetOp() antlr.Token { return s.op }

func (s *AddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *AddSubContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AddSubContext) ADD() antlr.TerminalNode {
	return s.GetToken(RiParserADD, 0)
}

func (s *AddSubContext) SUB() antlr.TerminalNode {
	return s.GetToken(RiParserSUB, 0)
}

func (s *AddSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiVisitor:
		return t.VisitAddSub(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdContext struct {
	*ExprContext
}

func NewIdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdContext {
	var p = new(IdContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdContext) ID() antlr.TerminalNode {
	return s.GetToken(RiParserID, 0)
}

func (s *IdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiVisitor:
		return t.VisitId(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntContext struct {
	*ExprContext
}

func NewIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntContext {
	var p = new(IntContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *IntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntContext) INT() antlr.TerminalNode {
	return s.GetToken(RiParserINT, 0)
}

func (s *IntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiVisitor:
		return t.VisitInt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RiParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *RiParser) expr(_p int) (localctx IExprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 0
	p.EnterRecursionRule(localctx, 0, RiParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(18)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RiParserTRUE, RiParserFALSE:
		localctx = NewBoolContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(11)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*BoolContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == RiParserTRUE || _la == RiParserFALSE) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*BoolContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case RiParserINT:
		localctx = NewIntContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(12)
			p.Match(RiParserINT)
		}

	case RiParserID:
		localctx = NewIdContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(13)
			p.Match(RiParserID)
		}

	case RiParserT__0:
		localctx = NewParensContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(14)
			p.Match(RiParserT__0)
		}
		{
			p.SetState(15)
			p.expr(0)
		}
		{
			p.SetState(16)
			p.Match(RiParserT__1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(29)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewModContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RiParserRULE_expr)
				p.SetState(20)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(21)
					p.Match(RiParserMOD)
				}
				{
					p.SetState(22)
					p.expr(8)
				}

			case 2:
				localctx = NewMulDivContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RiParserRULE_expr)
				p.SetState(23)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(24)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == RiParserMUL || _la == RiParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(25)
					p.expr(7)
				}

			case 3:
				localctx = NewAddSubContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RiParserRULE_expr)
				p.SetState(26)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(27)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddSubContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == RiParserADD || _la == RiParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddSubContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(28)
					p.expr(6)
				}

			}

		}
		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// IParExpressionContext is an interface to support dynamic dispatch.
type IParExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParExpressionContext differentiates from other interfaces.
	IsParExpressionContext()
}

type ParExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParExpressionContext() *ParExpressionContext {
	var p = new(ParExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RiParserRULE_parExpression
	return p
}

func (*ParExpressionContext) IsParExpressionContext() {}

func NewParExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParExpressionContext {
	var p = new(ParExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RiParserRULE_parExpression

	return p
}

func (s *ParExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ParExpressionContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiVisitor:
		return t.VisitParExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RiParser) ParExpression() (localctx IParExpressionContext) {
	this := p
	_ = this

	localctx = NewParExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RiParserRULE_parExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(34)
		p.Match(RiParserT__0)
	}
	{
		p.SetState(35)
		p.expr(0)
	}
	{
		p.SetState(36)
		p.Match(RiParserT__1)
	}

	return localctx
}

// IIfstatContext is an interface to support dynamic dispatch.
type IIfstatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfstatContext differentiates from other interfaces.
	IsIfstatContext()
}

type IfstatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfstatContext() *IfstatContext {
	var p = new(IfstatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RiParserRULE_ifstat
	return p
}

func (*IfstatContext) IsIfstatContext() {}

func NewIfstatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfstatContext {
	var p = new(IfstatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RiParserRULE_ifstat

	return p
}

func (s *IfstatContext) GetParser() antlr.Parser { return s.parser }

func (s *IfstatContext) AllStat() []IStatContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatContext); ok {
			len++
		}
	}

	tst := make([]IStatContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatContext); ok {
			tst[i] = t.(IStatContext)
			i++
		}
	}

	return tst
}

func (s *IfstatContext) Stat(i int) IStatContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *IfstatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfstatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfstatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiVisitor:
		return t.VisitIfstat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RiParser) Ifstat() (localctx IIfstatContext) {
	this := p
	_ = this

	localctx = NewIfstatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, RiParserRULE_ifstat)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1016706) != 0 {
		{
			p.SetState(38)
			p.Stat()
		}

		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatContext is an interface to support dynamic dispatch.
type IStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatContext differentiates from other interfaces.
	IsStatContext()
}

type StatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatContext() *StatContext {
	var p = new(StatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RiParserRULE_stat
	return p
}

func (*StatContext) IsStatContext() {}

func NewStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatContext {
	var p = new(StatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RiParserRULE_stat

	return p
}

func (s *StatContext) GetParser() antlr.Parser { return s.parser }

func (s *StatContext) CopyFrom(ctx *StatContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ForLogicContext struct {
	*StatContext
}

func NewForLogicContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForLogicContext {
	var p = new(ForLogicContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *ForLogicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForLogicContext) FOR() antlr.TerminalNode {
	return s.GetToken(RiParserFOR, 0)
}

func (s *ForLogicContext) AllStat() []IStatContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatContext); ok {
			len++
		}
	}

	tst := make([]IStatContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatContext); ok {
			tst[i] = t.(IStatContext)
			i++
		}
	}

	return tst
}

func (s *ForLogicContext) Stat(i int) IStatContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *ForLogicContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiVisitor:
		return t.VisitForLogic(s)

	default:
		return t.VisitChildren(s)
	}
}

type BlankContext struct {
	*StatContext
}

func NewBlankContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BlankContext {
	var p = new(BlankContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *BlankContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlankContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(RiParserNEWLINE, 0)
}

func (s *BlankContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiVisitor:
		return t.VisitBlank(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssginContext struct {
	*StatContext
}

func NewAssginContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssginContext {
	var p = new(AssginContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *AssginContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssginContext) ID() antlr.TerminalNode {
	return s.GetToken(RiParserID, 0)
}

func (s *AssginContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssginContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(RiParserNEWLINE, 0)
}

func (s *AssginContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiVisitor:
		return t.VisitAssgin(s)

	default:
		return t.VisitChildren(s)
	}
}

type EchoExprContext struct {
	*StatContext
}

func NewEchoExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EchoExprContext {
	var p = new(EchoExprContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *EchoExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EchoExprContext) ECHO() antlr.TerminalNode {
	return s.GetToken(RiParserECHO, 0)
}

func (s *EchoExprContext) ID() antlr.TerminalNode {
	return s.GetToken(RiParserID, 0)
}

func (s *EchoExprContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(RiParserNEWLINE, 0)
}

func (s *EchoExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiVisitor:
		return t.VisitEchoExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfLogicContext struct {
	*StatContext
}

func NewIfLogicContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfLogicContext {
	var p = new(IfLogicContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *IfLogicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfLogicContext) IF() antlr.TerminalNode {
	return s.GetToken(RiParserIF, 0)
}

func (s *IfLogicContext) ParExpression() IParExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParExpressionContext)
}

func (s *IfLogicContext) Ifstat() IIfstatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfstatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfstatContext)
}

func (s *IfLogicContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiVisitor:
		return t.VisitIfLogic(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrintExprContext struct {
	*StatContext
}

func NewPrintExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintExprContext {
	var p = new(PrintExprContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *PrintExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PrintExprContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(RiParserNEWLINE, 0)
}

func (s *PrintExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiVisitor:
		return t.VisitPrintExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RiParser) Stat() (localctx IStatContext) {
	this := p
	_ = this

	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, RiParserRULE_stat)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewEchoExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(43)
			p.Match(RiParserECHO)
		}
		{
			p.SetState(44)
			p.Match(RiParserT__0)
		}
		{
			p.SetState(45)
			p.Match(RiParserID)
		}
		{
			p.SetState(46)
			p.Match(RiParserT__1)
		}
		{
			p.SetState(47)
			p.Match(RiParserNEWLINE)
		}

	case 2:
		localctx = NewIfLogicContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(48)
			p.Match(RiParserIF)
		}
		{
			p.SetState(49)
			p.ParExpression()
		}
		{
			p.SetState(50)
			p.Match(RiParserT__2)
		}
		{
			p.SetState(51)
			p.Ifstat()
		}
		{
			p.SetState(52)
			p.Match(RiParserT__3)
		}

	case 3:
		localctx = NewPrintExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(54)
			p.expr(0)
		}
		{
			p.SetState(55)
			p.Match(RiParserNEWLINE)
		}

	case 4:
		localctx = NewAssginContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(57)
			p.Match(RiParserID)
		}
		{
			p.SetState(58)
			p.Match(RiParserT__4)
		}
		{
			p.SetState(59)
			p.expr(0)
		}
		{
			p.SetState(60)
			p.Match(RiParserNEWLINE)
		}

	case 5:
		localctx = NewForLogicContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(62)
			p.Match(RiParserFOR)
		}
		{
			p.SetState(63)
			p.Match(RiParserT__0)
		}
		{
			p.SetState(64)
			p.Match(RiParserT__5)
		}
		{
			p.SetState(65)
			p.Match(RiParserT__5)
		}
		{
			p.SetState(66)
			p.Match(RiParserT__1)
		}
		{
			p.SetState(67)
			p.Match(RiParserT__2)
		}
		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1016706) != 0 {
			{
				p.SetState(68)
				p.Stat()
			}

			p.SetState(71)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(73)
			p.Match(RiParserT__3)
		}

	case 6:
		localctx = NewBlankContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(75)
			p.Match(RiParserNEWLINE)
		}

	}

	return localctx
}

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RiParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RiParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) AllStat() []IStatContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatContext); ok {
			len++
		}
	}

	tst := make([]IStatContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatContext); ok {
			tst[i] = t.(IStatContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) Stat(i int) IStatContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiVisitor:
		return t.VisitProg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RiParser) Prog() (localctx IProgContext) {
	this := p
	_ = this

	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, RiParserRULE_prog)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1016706) != 0 {
		{
			p.SetState(78)
			p.Stat()
		}

		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

func (p *RiParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 0:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *RiParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
