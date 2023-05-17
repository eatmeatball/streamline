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
		"", "'('", "')'", "'{'", "'}'", "'='", "';'", "'import'", "'echo'",
		"'if'", "'for'", "'bool'", "'string'", "'float'", "'any'", "'func'",
		"'true'", "'false'", "", "", "", "", "", "'*'", "'/'", "'+'", "'-'",
		"'%'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "IMPORT", "ECHO", "IF", "FOR", "BOOL", "STRING",
		"FLOAT", "ANY", "FUNC", "TRUE", "FALSE", "ID", "INT", "NEWLINE", "WS",
		"DECIMAL_FLOAT_LIT", "MUL", "DIV", "ADD", "SUB", "MOD",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "IMPORT", "ECHO", "IF",
		"FOR", "BOOL", "STRING", "FLOAT", "ANY", "FUNC", "TRUE", "FALSE", "ID",
		"INT", "NEWLINE", "WS", "DECIMALS", "DECIMAL_FLOAT_LIT", "MUL", "DIV",
		"ADD", "SUB", "MOD",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 27, 173, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1,
		3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 4, 17, 128, 8, 17, 11,
		17, 12, 17, 129, 1, 18, 4, 18, 133, 8, 18, 11, 18, 12, 18, 134, 1, 19,
		3, 19, 138, 8, 19, 1, 19, 1, 19, 1, 20, 4, 20, 143, 8, 20, 11, 20, 12,
		20, 144, 1, 20, 1, 20, 1, 21, 1, 21, 3, 21, 151, 8, 21, 1, 21, 5, 21, 154,
		8, 21, 10, 21, 12, 21, 157, 9, 21, 1, 22, 1, 22, 1, 22, 3, 22, 162, 8,
		22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27,
		0, 0, 28, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19,
		10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37,
		19, 39, 20, 41, 21, 43, 0, 45, 22, 47, 23, 49, 24, 51, 25, 53, 26, 55,
		27, 1, 0, 3, 2, 0, 65, 90, 97, 122, 1, 0, 48, 57, 2, 0, 9, 9, 32, 32, 178,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0,
		0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1,
		0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39,
		1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0,
		49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0,
		1, 57, 1, 0, 0, 0, 3, 59, 1, 0, 0, 0, 5, 61, 1, 0, 0, 0, 7, 63, 1, 0, 0,
		0, 9, 65, 1, 0, 0, 0, 11, 67, 1, 0, 0, 0, 13, 69, 1, 0, 0, 0, 15, 76, 1,
		0, 0, 0, 17, 81, 1, 0, 0, 0, 19, 84, 1, 0, 0, 0, 21, 88, 1, 0, 0, 0, 23,
		93, 1, 0, 0, 0, 25, 100, 1, 0, 0, 0, 27, 106, 1, 0, 0, 0, 29, 110, 1, 0,
		0, 0, 31, 115, 1, 0, 0, 0, 33, 120, 1, 0, 0, 0, 35, 127, 1, 0, 0, 0, 37,
		132, 1, 0, 0, 0, 39, 137, 1, 0, 0, 0, 41, 142, 1, 0, 0, 0, 43, 148, 1,
		0, 0, 0, 45, 158, 1, 0, 0, 0, 47, 163, 1, 0, 0, 0, 49, 165, 1, 0, 0, 0,
		51, 167, 1, 0, 0, 0, 53, 169, 1, 0, 0, 0, 55, 171, 1, 0, 0, 0, 57, 58,
		5, 40, 0, 0, 58, 2, 1, 0, 0, 0, 59, 60, 5, 41, 0, 0, 60, 4, 1, 0, 0, 0,
		61, 62, 5, 123, 0, 0, 62, 6, 1, 0, 0, 0, 63, 64, 5, 125, 0, 0, 64, 8, 1,
		0, 0, 0, 65, 66, 5, 61, 0, 0, 66, 10, 1, 0, 0, 0, 67, 68, 5, 59, 0, 0,
		68, 12, 1, 0, 0, 0, 69, 70, 5, 105, 0, 0, 70, 71, 5, 109, 0, 0, 71, 72,
		5, 112, 0, 0, 72, 73, 5, 111, 0, 0, 73, 74, 5, 114, 0, 0, 74, 75, 5, 116,
		0, 0, 75, 14, 1, 0, 0, 0, 76, 77, 5, 101, 0, 0, 77, 78, 5, 99, 0, 0, 78,
		79, 5, 104, 0, 0, 79, 80, 5, 111, 0, 0, 80, 16, 1, 0, 0, 0, 81, 82, 5,
		105, 0, 0, 82, 83, 5, 102, 0, 0, 83, 18, 1, 0, 0, 0, 84, 85, 5, 102, 0,
		0, 85, 86, 5, 111, 0, 0, 86, 87, 5, 114, 0, 0, 87, 20, 1, 0, 0, 0, 88,
		89, 5, 98, 0, 0, 89, 90, 5, 111, 0, 0, 90, 91, 5, 111, 0, 0, 91, 92, 5,
		108, 0, 0, 92, 22, 1, 0, 0, 0, 93, 94, 5, 115, 0, 0, 94, 95, 5, 116, 0,
		0, 95, 96, 5, 114, 0, 0, 96, 97, 5, 105, 0, 0, 97, 98, 5, 110, 0, 0, 98,
		99, 5, 103, 0, 0, 99, 24, 1, 0, 0, 0, 100, 101, 5, 102, 0, 0, 101, 102,
		5, 108, 0, 0, 102, 103, 5, 111, 0, 0, 103, 104, 5, 97, 0, 0, 104, 105,
		5, 116, 0, 0, 105, 26, 1, 0, 0, 0, 106, 107, 5, 97, 0, 0, 107, 108, 5,
		110, 0, 0, 108, 109, 5, 121, 0, 0, 109, 28, 1, 0, 0, 0, 110, 111, 5, 102,
		0, 0, 111, 112, 5, 117, 0, 0, 112, 113, 5, 110, 0, 0, 113, 114, 5, 99,
		0, 0, 114, 30, 1, 0, 0, 0, 115, 116, 5, 116, 0, 0, 116, 117, 5, 114, 0,
		0, 117, 118, 5, 117, 0, 0, 118, 119, 5, 101, 0, 0, 119, 32, 1, 0, 0, 0,
		120, 121, 5, 102, 0, 0, 121, 122, 5, 97, 0, 0, 122, 123, 5, 108, 0, 0,
		123, 124, 5, 115, 0, 0, 124, 125, 5, 101, 0, 0, 125, 34, 1, 0, 0, 0, 126,
		128, 7, 0, 0, 0, 127, 126, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 127,
		1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 36, 1, 0, 0, 0, 131, 133, 7, 1,
		0, 0, 132, 131, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134, 132, 1, 0, 0, 0,
		134, 135, 1, 0, 0, 0, 135, 38, 1, 0, 0, 0, 136, 138, 5, 13, 0, 0, 137,
		136, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 140,
		5, 10, 0, 0, 140, 40, 1, 0, 0, 0, 141, 143, 7, 2, 0, 0, 142, 141, 1, 0,
		0, 0, 143, 144, 1, 0, 0, 0, 144, 142, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0,
		145, 146, 1, 0, 0, 0, 146, 147, 6, 20, 0, 0, 147, 42, 1, 0, 0, 0, 148,
		155, 7, 1, 0, 0, 149, 151, 5, 95, 0, 0, 150, 149, 1, 0, 0, 0, 150, 151,
		1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 154, 7, 1, 0, 0, 153, 150, 1, 0,
		0, 0, 154, 157, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0,
		156, 44, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0, 158, 159, 3, 43, 21, 0, 159,
		161, 5, 46, 0, 0, 160, 162, 3, 43, 21, 0, 161, 160, 1, 0, 0, 0, 161, 162,
		1, 0, 0, 0, 162, 46, 1, 0, 0, 0, 163, 164, 5, 42, 0, 0, 164, 48, 1, 0,
		0, 0, 165, 166, 5, 47, 0, 0, 166, 50, 1, 0, 0, 0, 167, 168, 5, 43, 0, 0,
		168, 52, 1, 0, 0, 0, 169, 170, 5, 45, 0, 0, 170, 54, 1, 0, 0, 0, 171, 172,
		5, 37, 0, 0, 172, 56, 1, 0, 0, 0, 8, 0, 129, 134, 137, 144, 150, 155, 161,
		1, 6, 0, 0,
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
	RiLexerT__5              = 6
	RiLexerIMPORT            = 7
	RiLexerECHO              = 8
	RiLexerIF                = 9
	RiLexerFOR               = 10
	RiLexerBOOL              = 11
	RiLexerSTRING            = 12
	RiLexerFLOAT             = 13
	RiLexerANY               = 14
	RiLexerFUNC              = 15
	RiLexerTRUE              = 16
	RiLexerFALSE             = 17
	RiLexerID                = 18
	RiLexerINT               = 19
	RiLexerNEWLINE           = 20
	RiLexerWS                = 21
	RiLexerDECIMAL_FLOAT_LIT = 22
	RiLexerMUL               = 23
	RiLexerDIV               = 24
	RiLexerADD               = 25
	RiLexerSUB               = 26
	RiLexerMOD               = 27
)
