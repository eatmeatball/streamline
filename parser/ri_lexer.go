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
		"'false'", "'string'", "'float'", "'any'", "", "", "", "", "", "'*'",
		"'/'", "'+'", "'-'", "'%'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "ECHO", "IF", "BOOL", "TRUE", "FALSE", "STRING",
		"FLOAT", "ANY", "ID", "INT", "NEWLINE", "WS", "DECIMAL_FLOAT_LIT", "MUL",
		"DIV", "ADD", "SUB", "MOD",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "ECHO", "IF", "BOOL", "TRUE",
		"FALSE", "STRING", "FLOAT", "ANY", "ID", "INT", "NEWLINE", "WS", "DECIMALS",
		"DECIMAL_FLOAT_LIT", "MUL", "DIV", "ADD", "SUB", "MOD",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 23, 147, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1,
		6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 13, 4, 13, 102, 8, 13, 11, 13, 12, 13, 103, 1, 14, 4, 14, 107,
		8, 14, 11, 14, 12, 14, 108, 1, 15, 3, 15, 112, 8, 15, 1, 15, 1, 15, 1,
		16, 4, 16, 117, 8, 16, 11, 16, 12, 16, 118, 1, 16, 1, 16, 1, 17, 1, 17,
		3, 17, 125, 8, 17, 1, 17, 5, 17, 128, 8, 17, 10, 17, 12, 17, 131, 9, 17,
		1, 18, 1, 18, 1, 18, 3, 18, 136, 8, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1,
		21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 0, 0, 24, 1, 1, 3, 2, 5, 3, 7, 4,
		9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 31, 16, 33, 17, 35, 0, 37, 18, 39, 19, 41, 20, 43, 21, 45, 22,
		47, 23, 1, 0, 3, 2, 0, 65, 90, 97, 122, 1, 0, 48, 57, 2, 0, 9, 9, 32, 32,
		152, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0,
		0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1,
		0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23,
		1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0,
		31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0,
		0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0,
		0, 1, 49, 1, 0, 0, 0, 3, 51, 1, 0, 0, 0, 5, 53, 1, 0, 0, 0, 7, 55, 1, 0,
		0, 0, 9, 57, 1, 0, 0, 0, 11, 59, 1, 0, 0, 0, 13, 64, 1, 0, 0, 0, 15, 67,
		1, 0, 0, 0, 17, 72, 1, 0, 0, 0, 19, 77, 1, 0, 0, 0, 21, 83, 1, 0, 0, 0,
		23, 90, 1, 0, 0, 0, 25, 96, 1, 0, 0, 0, 27, 101, 1, 0, 0, 0, 29, 106, 1,
		0, 0, 0, 31, 111, 1, 0, 0, 0, 33, 116, 1, 0, 0, 0, 35, 122, 1, 0, 0, 0,
		37, 132, 1, 0, 0, 0, 39, 137, 1, 0, 0, 0, 41, 139, 1, 0, 0, 0, 43, 141,
		1, 0, 0, 0, 45, 143, 1, 0, 0, 0, 47, 145, 1, 0, 0, 0, 49, 50, 5, 40, 0,
		0, 50, 2, 1, 0, 0, 0, 51, 52, 5, 41, 0, 0, 52, 4, 1, 0, 0, 0, 53, 54, 5,
		123, 0, 0, 54, 6, 1, 0, 0, 0, 55, 56, 5, 125, 0, 0, 56, 8, 1, 0, 0, 0,
		57, 58, 5, 61, 0, 0, 58, 10, 1, 0, 0, 0, 59, 60, 5, 101, 0, 0, 60, 61,
		5, 99, 0, 0, 61, 62, 5, 104, 0, 0, 62, 63, 5, 111, 0, 0, 63, 12, 1, 0,
		0, 0, 64, 65, 5, 105, 0, 0, 65, 66, 5, 102, 0, 0, 66, 14, 1, 0, 0, 0, 67,
		68, 5, 98, 0, 0, 68, 69, 5, 111, 0, 0, 69, 70, 5, 111, 0, 0, 70, 71, 5,
		108, 0, 0, 71, 16, 1, 0, 0, 0, 72, 73, 5, 116, 0, 0, 73, 74, 5, 114, 0,
		0, 74, 75, 5, 117, 0, 0, 75, 76, 5, 101, 0, 0, 76, 18, 1, 0, 0, 0, 77,
		78, 5, 102, 0, 0, 78, 79, 5, 97, 0, 0, 79, 80, 5, 108, 0, 0, 80, 81, 5,
		115, 0, 0, 81, 82, 5, 101, 0, 0, 82, 20, 1, 0, 0, 0, 83, 84, 5, 115, 0,
		0, 84, 85, 5, 116, 0, 0, 85, 86, 5, 114, 0, 0, 86, 87, 5, 105, 0, 0, 87,
		88, 5, 110, 0, 0, 88, 89, 5, 103, 0, 0, 89, 22, 1, 0, 0, 0, 90, 91, 5,
		102, 0, 0, 91, 92, 5, 108, 0, 0, 92, 93, 5, 111, 0, 0, 93, 94, 5, 97, 0,
		0, 94, 95, 5, 116, 0, 0, 95, 24, 1, 0, 0, 0, 96, 97, 5, 97, 0, 0, 97, 98,
		5, 110, 0, 0, 98, 99, 5, 121, 0, 0, 99, 26, 1, 0, 0, 0, 100, 102, 7, 0,
		0, 0, 101, 100, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0,
		103, 104, 1, 0, 0, 0, 104, 28, 1, 0, 0, 0, 105, 107, 7, 1, 0, 0, 106, 105,
		1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0, 108, 109, 1, 0,
		0, 0, 109, 30, 1, 0, 0, 0, 110, 112, 5, 13, 0, 0, 111, 110, 1, 0, 0, 0,
		111, 112, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113, 114, 5, 10, 0, 0, 114,
		32, 1, 0, 0, 0, 115, 117, 7, 2, 0, 0, 116, 115, 1, 0, 0, 0, 117, 118, 1,
		0, 0, 0, 118, 116, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119, 120, 1, 0, 0,
		0, 120, 121, 6, 16, 0, 0, 121, 34, 1, 0, 0, 0, 122, 129, 7, 1, 0, 0, 123,
		125, 5, 95, 0, 0, 124, 123, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 126,
		1, 0, 0, 0, 126, 128, 7, 1, 0, 0, 127, 124, 1, 0, 0, 0, 128, 131, 1, 0,
		0, 0, 129, 127, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 36, 1, 0, 0, 0,
		131, 129, 1, 0, 0, 0, 132, 133, 3, 35, 17, 0, 133, 135, 5, 46, 0, 0, 134,
		136, 3, 35, 17, 0, 135, 134, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136, 38,
		1, 0, 0, 0, 137, 138, 5, 42, 0, 0, 138, 40, 1, 0, 0, 0, 139, 140, 5, 47,
		0, 0, 140, 42, 1, 0, 0, 0, 141, 142, 5, 43, 0, 0, 142, 44, 1, 0, 0, 0,
		143, 144, 5, 45, 0, 0, 144, 46, 1, 0, 0, 0, 145, 146, 5, 37, 0, 0, 146,
		48, 1, 0, 0, 0, 8, 0, 103, 108, 111, 118, 124, 129, 135, 1, 6, 0, 0,
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
	RiLexerT__0              = 1
	RiLexerT__1              = 2
	RiLexerT__2              = 3
	RiLexerT__3              = 4
	RiLexerT__4              = 5
	RiLexerECHO              = 6
	RiLexerIF                = 7
	RiLexerBOOL              = 8
	RiLexerTRUE              = 9
	RiLexerFALSE             = 10
	RiLexerSTRING            = 11
	RiLexerFLOAT             = 12
	RiLexerANY               = 13
	RiLexerID                = 14
	RiLexerINT               = 15
	RiLexerNEWLINE           = 16
	RiLexerWS                = 17
	RiLexerDECIMAL_FLOAT_LIT = 18
	RiLexerMUL               = 19
	RiLexerDIV               = 20
	RiLexerADD               = 21
	RiLexerSUB               = 22
	RiLexerMOD               = 23
)
