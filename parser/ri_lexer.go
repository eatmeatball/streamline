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
		"", "'('", "')'", "'{'", "'}'", "'func'", "'int'", "'float'", "'string'",
		"'bool'", "'any'", "'*'", "'/'", "'+'", "'-'", "", "", "'return'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "FUNC", "INT", "FLOAT", "STRING", "BOOL", "ANY",
		"MUL", "DIV", "ADD", "SUB", "NUMBER", "WHITESPACE", "RETURN", "IDENTIFIER",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "Letter", "LetterOrDigit", "FUNC", "INT",
		"FLOAT", "STRING", "BOOL", "ANY", "MUL", "DIV", "ADD", "SUB", "NUMBER",
		"WHITESPACE", "RETURN", "IDENTIFIER",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 18, 124, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 54,
		8, 4, 1, 5, 1, 5, 3, 5, 58, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16,
		4, 16, 100, 8, 16, 11, 16, 12, 16, 101, 1, 17, 4, 17, 105, 8, 17, 11, 17,
		12, 17, 106, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 19, 1, 19, 5, 19, 120, 8, 19, 10, 19, 12, 19, 123, 9, 19, 0, 0,
		20, 1, 1, 3, 2, 5, 3, 7, 4, 9, 0, 11, 0, 13, 5, 15, 6, 17, 7, 19, 8, 21,
		9, 23, 10, 25, 11, 27, 12, 29, 13, 31, 14, 33, 15, 35, 16, 37, 17, 39,
		18, 1, 0, 6, 4, 0, 36, 36, 65, 90, 95, 95, 97, 122, 2, 0, 0, 127, 55296,
		56319, 1, 0, 55296, 56319, 1, 0, 56320, 57343, 1, 0, 48, 57, 3, 0, 9, 10,
		13, 13, 32, 32, 127, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0,
		0, 0, 7, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0,
		0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1,
		0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33,
		1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 1,
		41, 1, 0, 0, 0, 3, 43, 1, 0, 0, 0, 5, 45, 1, 0, 0, 0, 7, 47, 1, 0, 0, 0,
		9, 53, 1, 0, 0, 0, 11, 57, 1, 0, 0, 0, 13, 59, 1, 0, 0, 0, 15, 64, 1, 0,
		0, 0, 17, 68, 1, 0, 0, 0, 19, 74, 1, 0, 0, 0, 21, 81, 1, 0, 0, 0, 23, 86,
		1, 0, 0, 0, 25, 90, 1, 0, 0, 0, 27, 92, 1, 0, 0, 0, 29, 94, 1, 0, 0, 0,
		31, 96, 1, 0, 0, 0, 33, 99, 1, 0, 0, 0, 35, 104, 1, 0, 0, 0, 37, 110, 1,
		0, 0, 0, 39, 117, 1, 0, 0, 0, 41, 42, 5, 40, 0, 0, 42, 2, 1, 0, 0, 0, 43,
		44, 5, 41, 0, 0, 44, 4, 1, 0, 0, 0, 45, 46, 5, 123, 0, 0, 46, 6, 1, 0,
		0, 0, 47, 48, 5, 125, 0, 0, 48, 8, 1, 0, 0, 0, 49, 54, 7, 0, 0, 0, 50,
		54, 8, 1, 0, 0, 51, 52, 7, 2, 0, 0, 52, 54, 7, 3, 0, 0, 53, 49, 1, 0, 0,
		0, 53, 50, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 54, 10, 1, 0, 0, 0, 55, 58,
		3, 9, 4, 0, 56, 58, 7, 4, 0, 0, 57, 55, 1, 0, 0, 0, 57, 56, 1, 0, 0, 0,
		58, 12, 1, 0, 0, 0, 59, 60, 5, 102, 0, 0, 60, 61, 5, 117, 0, 0, 61, 62,
		5, 110, 0, 0, 62, 63, 5, 99, 0, 0, 63, 14, 1, 0, 0, 0, 64, 65, 5, 105,
		0, 0, 65, 66, 5, 110, 0, 0, 66, 67, 5, 116, 0, 0, 67, 16, 1, 0, 0, 0, 68,
		69, 5, 102, 0, 0, 69, 70, 5, 108, 0, 0, 70, 71, 5, 111, 0, 0, 71, 72, 5,
		97, 0, 0, 72, 73, 5, 116, 0, 0, 73, 18, 1, 0, 0, 0, 74, 75, 5, 115, 0,
		0, 75, 76, 5, 116, 0, 0, 76, 77, 5, 114, 0, 0, 77, 78, 5, 105, 0, 0, 78,
		79, 5, 110, 0, 0, 79, 80, 5, 103, 0, 0, 80, 20, 1, 0, 0, 0, 81, 82, 5,
		98, 0, 0, 82, 83, 5, 111, 0, 0, 83, 84, 5, 111, 0, 0, 84, 85, 5, 108, 0,
		0, 85, 22, 1, 0, 0, 0, 86, 87, 5, 97, 0, 0, 87, 88, 5, 110, 0, 0, 88, 89,
		5, 121, 0, 0, 89, 24, 1, 0, 0, 0, 90, 91, 5, 42, 0, 0, 91, 26, 1, 0, 0,
		0, 92, 93, 5, 47, 0, 0, 93, 28, 1, 0, 0, 0, 94, 95, 5, 43, 0, 0, 95, 30,
		1, 0, 0, 0, 96, 97, 5, 45, 0, 0, 97, 32, 1, 0, 0, 0, 98, 100, 7, 4, 0,
		0, 99, 98, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 99, 1, 0, 0, 0, 101,
		102, 1, 0, 0, 0, 102, 34, 1, 0, 0, 0, 103, 105, 7, 5, 0, 0, 104, 103, 1,
		0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 104, 1, 0, 0, 0, 106, 107, 1, 0, 0,
		0, 107, 108, 1, 0, 0, 0, 108, 109, 6, 17, 0, 0, 109, 36, 1, 0, 0, 0, 110,
		111, 5, 114, 0, 0, 111, 112, 5, 101, 0, 0, 112, 113, 5, 116, 0, 0, 113,
		114, 5, 117, 0, 0, 114, 115, 5, 114, 0, 0, 115, 116, 5, 110, 0, 0, 116,
		38, 1, 0, 0, 0, 117, 121, 3, 9, 4, 0, 118, 120, 3, 11, 5, 0, 119, 118,
		1, 0, 0, 0, 120, 123, 1, 0, 0, 0, 121, 119, 1, 0, 0, 0, 121, 122, 1, 0,
		0, 0, 122, 40, 1, 0, 0, 0, 123, 121, 1, 0, 0, 0, 6, 0, 53, 57, 101, 106,
		121, 1, 6, 0, 0,
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
	RiLexerFUNC       = 5
	RiLexerINT        = 6
	RiLexerFLOAT      = 7
	RiLexerSTRING     = 8
	RiLexerBOOL       = 9
	RiLexerANY        = 10
	RiLexerMUL        = 11
	RiLexerDIV        = 12
	RiLexerADD        = 13
	RiLexerSUB        = 14
	RiLexerNUMBER     = 15
	RiLexerWHITESPACE = 16
	RiLexerRETURN     = 17
	RiLexerIDENTIFIER = 18
)
