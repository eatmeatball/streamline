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
		"", "'('", "')'", "'{'", "'}'", "'class'", "'func'", "'numberFunction'",
		"'int'", "'float'", "'string'", "'bool'", "'any'", "'*'", "'/'", "'+'",
		"'-'", "", "", "'return'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "FUNC", "FUNCNAME", "INT", "FLOAT", "STRING",
		"BOOL", "ANY", "MUL", "DIV", "ADD", "SUB", "NUMBER", "WHITESPACE", "RETURN",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "FUNC", "FUNCNAME", "INT", "FLOAT",
		"STRING", "BOOL", "ANY", "MUL", "DIV", "ADD", "SUB", "NUMBER", "WHITESPACE",
		"RETURN",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 19, 126, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1,
		13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 4, 16, 109, 8, 16, 11, 16,
		12, 16, 110, 1, 17, 4, 17, 114, 8, 17, 11, 17, 12, 17, 115, 1, 17, 1, 17,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 0, 0, 19, 1, 1, 3, 2,
		5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25,
		13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 1, 0, 2, 1, 0, 48,
		57, 3, 0, 9, 10, 13, 13, 32, 32, 127, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0,
		0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0,
		0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0,
		0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1,
		0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35,
		1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 1, 39, 1, 0, 0, 0, 3, 41, 1, 0, 0, 0, 5,
		43, 1, 0, 0, 0, 7, 45, 1, 0, 0, 0, 9, 47, 1, 0, 0, 0, 11, 53, 1, 0, 0,
		0, 13, 58, 1, 0, 0, 0, 15, 73, 1, 0, 0, 0, 17, 77, 1, 0, 0, 0, 19, 83,
		1, 0, 0, 0, 21, 90, 1, 0, 0, 0, 23, 95, 1, 0, 0, 0, 25, 99, 1, 0, 0, 0,
		27, 101, 1, 0, 0, 0, 29, 103, 1, 0, 0, 0, 31, 105, 1, 0, 0, 0, 33, 108,
		1, 0, 0, 0, 35, 113, 1, 0, 0, 0, 37, 119, 1, 0, 0, 0, 39, 40, 5, 40, 0,
		0, 40, 2, 1, 0, 0, 0, 41, 42, 5, 41, 0, 0, 42, 4, 1, 0, 0, 0, 43, 44, 5,
		123, 0, 0, 44, 6, 1, 0, 0, 0, 45, 46, 5, 125, 0, 0, 46, 8, 1, 0, 0, 0,
		47, 48, 5, 99, 0, 0, 48, 49, 5, 108, 0, 0, 49, 50, 5, 97, 0, 0, 50, 51,
		5, 115, 0, 0, 51, 52, 5, 115, 0, 0, 52, 10, 1, 0, 0, 0, 53, 54, 5, 102,
		0, 0, 54, 55, 5, 117, 0, 0, 55, 56, 5, 110, 0, 0, 56, 57, 5, 99, 0, 0,
		57, 12, 1, 0, 0, 0, 58, 59, 5, 110, 0, 0, 59, 60, 5, 117, 0, 0, 60, 61,
		5, 109, 0, 0, 61, 62, 5, 98, 0, 0, 62, 63, 5, 101, 0, 0, 63, 64, 5, 114,
		0, 0, 64, 65, 5, 70, 0, 0, 65, 66, 5, 117, 0, 0, 66, 67, 5, 110, 0, 0,
		67, 68, 5, 99, 0, 0, 68, 69, 5, 116, 0, 0, 69, 70, 5, 105, 0, 0, 70, 71,
		5, 111, 0, 0, 71, 72, 5, 110, 0, 0, 72, 14, 1, 0, 0, 0, 73, 74, 5, 105,
		0, 0, 74, 75, 5, 110, 0, 0, 75, 76, 5, 116, 0, 0, 76, 16, 1, 0, 0, 0, 77,
		78, 5, 102, 0, 0, 78, 79, 5, 108, 0, 0, 79, 80, 5, 111, 0, 0, 80, 81, 5,
		97, 0, 0, 81, 82, 5, 116, 0, 0, 82, 18, 1, 0, 0, 0, 83, 84, 5, 115, 0,
		0, 84, 85, 5, 116, 0, 0, 85, 86, 5, 114, 0, 0, 86, 87, 5, 105, 0, 0, 87,
		88, 5, 110, 0, 0, 88, 89, 5, 103, 0, 0, 89, 20, 1, 0, 0, 0, 90, 91, 5,
		98, 0, 0, 91, 92, 5, 111, 0, 0, 92, 93, 5, 111, 0, 0, 93, 94, 5, 108, 0,
		0, 94, 22, 1, 0, 0, 0, 95, 96, 5, 97, 0, 0, 96, 97, 5, 110, 0, 0, 97, 98,
		5, 121, 0, 0, 98, 24, 1, 0, 0, 0, 99, 100, 5, 42, 0, 0, 100, 26, 1, 0,
		0, 0, 101, 102, 5, 47, 0, 0, 102, 28, 1, 0, 0, 0, 103, 104, 5, 43, 0, 0,
		104, 30, 1, 0, 0, 0, 105, 106, 5, 45, 0, 0, 106, 32, 1, 0, 0, 0, 107, 109,
		7, 0, 0, 0, 108, 107, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 108, 1, 0,
		0, 0, 110, 111, 1, 0, 0, 0, 111, 34, 1, 0, 0, 0, 112, 114, 7, 1, 0, 0,
		113, 112, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 115,
		116, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 118, 6, 17, 0, 0, 118, 36,
		1, 0, 0, 0, 119, 120, 5, 114, 0, 0, 120, 121, 5, 101, 0, 0, 121, 122, 5,
		116, 0, 0, 122, 123, 5, 117, 0, 0, 123, 124, 5, 114, 0, 0, 124, 125, 5,
		110, 0, 0, 125, 38, 1, 0, 0, 0, 3, 0, 110, 115, 1, 6, 0, 0,
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
	RiLexerT__0       = 1
	RiLexerT__1       = 2
	RiLexerT__2       = 3
	RiLexerT__3       = 4
	RiLexerT__4       = 5
	RiLexerFUNC       = 6
	RiLexerFUNCNAME   = 7
	RiLexerINT        = 8
	RiLexerFLOAT      = 9
	RiLexerSTRING     = 10
	RiLexerBOOL       = 11
	RiLexerANY        = 12
	RiLexerMUL        = 13
	RiLexerDIV        = 14
	RiLexerADD        = 15
	RiLexerSUB        = 16
	RiLexerNUMBER     = 17
	RiLexerWHITESPACE = 18
	RiLexerRETURN     = 19
)
