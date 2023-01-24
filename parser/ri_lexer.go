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
		"", "'('", "')'", "'{'", "'}'", "'='", "'echo'", "'if'", "'bool'", "'true'",
		"'false'", "", "", "", "", "'*'", "'/'", "'+'", "'-'", "'%'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "ECHO", "IF", "BOOL", "TRUE", "FALSE", "ID",
		"INT", "NEWLINE", "WS", "MUL", "DIV", "ADD", "SUB", "MOD",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "ECHO", "IF", "BOOL", "TRUE",
		"FALSE", "ID", "INT", "NEWLINE", "WS", "MUL", "DIV", "ADD", "SUB", "MOD",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 19, 105, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6,
		1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 4, 10, 75, 8, 10, 11, 10, 12,
		10, 76, 1, 11, 4, 11, 80, 8, 11, 11, 11, 12, 11, 81, 1, 12, 3, 12, 85,
		8, 12, 1, 12, 1, 12, 1, 13, 4, 13, 90, 8, 13, 11, 13, 12, 13, 91, 1, 13,
		1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1,
		18, 0, 0, 19, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9,
		19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18,
		37, 19, 1, 0, 3, 2, 0, 65, 90, 97, 122, 1, 0, 48, 57, 2, 0, 9, 9, 32, 32,
		108, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0,
		0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1,
		0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23,
		1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0,
		31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0,
		1, 39, 1, 0, 0, 0, 3, 41, 1, 0, 0, 0, 5, 43, 1, 0, 0, 0, 7, 45, 1, 0, 0,
		0, 9, 47, 1, 0, 0, 0, 11, 49, 1, 0, 0, 0, 13, 54, 1, 0, 0, 0, 15, 57, 1,
		0, 0, 0, 17, 62, 1, 0, 0, 0, 19, 67, 1, 0, 0, 0, 21, 74, 1, 0, 0, 0, 23,
		79, 1, 0, 0, 0, 25, 84, 1, 0, 0, 0, 27, 89, 1, 0, 0, 0, 29, 95, 1, 0, 0,
		0, 31, 97, 1, 0, 0, 0, 33, 99, 1, 0, 0, 0, 35, 101, 1, 0, 0, 0, 37, 103,
		1, 0, 0, 0, 39, 40, 5, 40, 0, 0, 40, 2, 1, 0, 0, 0, 41, 42, 5, 41, 0, 0,
		42, 4, 1, 0, 0, 0, 43, 44, 5, 123, 0, 0, 44, 6, 1, 0, 0, 0, 45, 46, 5,
		125, 0, 0, 46, 8, 1, 0, 0, 0, 47, 48, 5, 61, 0, 0, 48, 10, 1, 0, 0, 0,
		49, 50, 5, 101, 0, 0, 50, 51, 5, 99, 0, 0, 51, 52, 5, 104, 0, 0, 52, 53,
		5, 111, 0, 0, 53, 12, 1, 0, 0, 0, 54, 55, 5, 105, 0, 0, 55, 56, 5, 102,
		0, 0, 56, 14, 1, 0, 0, 0, 57, 58, 5, 98, 0, 0, 58, 59, 5, 111, 0, 0, 59,
		60, 5, 111, 0, 0, 60, 61, 5, 108, 0, 0, 61, 16, 1, 0, 0, 0, 62, 63, 5,
		116, 0, 0, 63, 64, 5, 114, 0, 0, 64, 65, 5, 117, 0, 0, 65, 66, 5, 101,
		0, 0, 66, 18, 1, 0, 0, 0, 67, 68, 5, 102, 0, 0, 68, 69, 5, 97, 0, 0, 69,
		70, 5, 108, 0, 0, 70, 71, 5, 115, 0, 0, 71, 72, 5, 101, 0, 0, 72, 20, 1,
		0, 0, 0, 73, 75, 7, 0, 0, 0, 74, 73, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76,
		74, 1, 0, 0, 0, 76, 77, 1, 0, 0, 0, 77, 22, 1, 0, 0, 0, 78, 80, 7, 1, 0,
		0, 79, 78, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 81, 79, 1, 0, 0, 0, 81, 82,
		1, 0, 0, 0, 82, 24, 1, 0, 0, 0, 83, 85, 5, 13, 0, 0, 84, 83, 1, 0, 0, 0,
		84, 85, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 87, 5, 10, 0, 0, 87, 26, 1,
		0, 0, 0, 88, 90, 7, 2, 0, 0, 89, 88, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91,
		89, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 94, 6, 13,
		0, 0, 94, 28, 1, 0, 0, 0, 95, 96, 5, 42, 0, 0, 96, 30, 1, 0, 0, 0, 97,
		98, 5, 47, 0, 0, 98, 32, 1, 0, 0, 0, 99, 100, 5, 43, 0, 0, 100, 34, 1,
		0, 0, 0, 101, 102, 5, 45, 0, 0, 102, 36, 1, 0, 0, 0, 103, 104, 5, 37, 0,
		0, 104, 38, 1, 0, 0, 0, 5, 0, 76, 81, 84, 91, 1, 6, 0, 0,
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
	RiLexerT__3    = 4
	RiLexerT__4    = 5
	RiLexerECHO    = 6
	RiLexerIF      = 7
	RiLexerBOOL    = 8
	RiLexerTRUE    = 9
	RiLexerFALSE   = 10
	RiLexerID      = 11
	RiLexerINT     = 12
	RiLexerNEWLINE = 13
	RiLexerWS      = 14
	RiLexerMUL     = 15
	RiLexerDIV     = 16
	RiLexerADD     = 17
	RiLexerSUB     = 18
	RiLexerMOD     = 19
)
