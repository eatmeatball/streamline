// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type RiLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var rilexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func rilexerLexerInit() {
	staticData := &rilexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'('", "')'", "'='", "'echo'", "", "", "", "", "'*'", "'/'", "'+'",
		"'-'", "'%'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "ECHO", "ID", "INT", "NEWLINE", "WS", "MUL", "DIV",
		"ADD", "SUB", "MOD",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "ECHO", "ID", "INT", "NEWLINE", "WS", "MUL",
		"DIV", "ADD", "SUB", "MOD",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 13, 70, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 4, 4, 40, 8, 4, 11, 4, 12, 4, 41,
		1, 5, 4, 5, 45, 8, 5, 11, 5, 12, 5, 46, 1, 6, 3, 6, 50, 8, 6, 1, 6, 1,
		6, 1, 7, 4, 7, 55, 8, 7, 11, 7, 12, 7, 56, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9,
		1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 0, 0, 13, 1, 1, 3, 2, 5,
		3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25,
		13, 1, 0, 3, 2, 0, 65, 90, 97, 122, 1, 0, 48, 57, 2, 0, 9, 9, 32, 32, 73,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0,
		0, 0, 0, 25, 1, 0, 0, 0, 1, 27, 1, 0, 0, 0, 3, 29, 1, 0, 0, 0, 5, 31, 1,
		0, 0, 0, 7, 33, 1, 0, 0, 0, 9, 39, 1, 0, 0, 0, 11, 44, 1, 0, 0, 0, 13,
		49, 1, 0, 0, 0, 15, 54, 1, 0, 0, 0, 17, 60, 1, 0, 0, 0, 19, 62, 1, 0, 0,
		0, 21, 64, 1, 0, 0, 0, 23, 66, 1, 0, 0, 0, 25, 68, 1, 0, 0, 0, 27, 28,
		5, 40, 0, 0, 28, 2, 1, 0, 0, 0, 29, 30, 5, 41, 0, 0, 30, 4, 1, 0, 0, 0,
		31, 32, 5, 61, 0, 0, 32, 6, 1, 0, 0, 0, 33, 34, 5, 101, 0, 0, 34, 35, 5,
		99, 0, 0, 35, 36, 5, 104, 0, 0, 36, 37, 5, 111, 0, 0, 37, 8, 1, 0, 0, 0,
		38, 40, 7, 0, 0, 0, 39, 38, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 39, 1,
		0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 10, 1, 0, 0, 0, 43, 45, 7, 1, 0, 0, 44,
		43, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 44, 1, 0, 0, 0, 46, 47, 1, 0, 0,
		0, 47, 12, 1, 0, 0, 0, 48, 50, 5, 13, 0, 0, 49, 48, 1, 0, 0, 0, 49, 50,
		1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 52, 5, 10, 0, 0, 52, 14, 1, 0, 0, 0,
		53, 55, 7, 2, 0, 0, 54, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 54, 1,
		0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 59, 6, 7, 0, 0, 59,
		16, 1, 0, 0, 0, 60, 61, 5, 42, 0, 0, 61, 18, 1, 0, 0, 0, 62, 63, 5, 47,
		0, 0, 63, 20, 1, 0, 0, 0, 64, 65, 5, 43, 0, 0, 65, 22, 1, 0, 0, 0, 66,
		67, 5, 45, 0, 0, 67, 24, 1, 0, 0, 0, 68, 69, 5, 37, 0, 0, 69, 26, 1, 0,
		0, 0, 5, 0, 41, 46, 49, 56, 1, 6, 0, 0,
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

// RiLexerInit initializes any static state used to implement RiLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewRiLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func RiLexerInit() {
	staticData := &rilexerLexerStaticData
	staticData.once.Do(rilexerLexerInit)
}

// NewRiLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewRiLexer(input antlr.CharStream) *RiLexer {
	RiLexerInit()
	l := new(RiLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &rilexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Ri.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// RiLexer tokens.
const (
	RiLexerT__0    = 1
	RiLexerT__1    = 2
	RiLexerT__2    = 3
	RiLexerECHO    = 4
	RiLexerID      = 5
	RiLexerINT     = 6
	RiLexerNEWLINE = 7
	RiLexerWS      = 8
	RiLexerMUL     = 9
	RiLexerDIV     = 10
	RiLexerADD     = 11
	RiLexerSUB     = 12
	RiLexerMOD     = 13
)
