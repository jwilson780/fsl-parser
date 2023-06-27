// Code generated from FSL.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type FSLLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var FSLLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func fsllexerLexerInit() {
	staticData := &FSLLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "':'", "'['", "']'", "'cmd'", "'add'", "','", "'id'", "'operand1'",
		"'operand2'", "'subtract'", "'multiply'", "'divide'", "'create'", "'value'",
		"'update'", "'delete'", "'print'", "'#'", "'$'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "ID", "INT", "FLOAT", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "ID", "INT", "FLOAT", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 23, 176, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1,
		6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 5, 19, 149, 8, 19, 10, 19, 12, 19,
		152, 9, 19, 1, 20, 4, 20, 155, 8, 20, 11, 20, 12, 20, 156, 1, 21, 4, 21,
		160, 8, 21, 11, 21, 12, 21, 161, 1, 21, 1, 21, 4, 21, 166, 8, 21, 11, 21,
		12, 21, 167, 1, 22, 4, 22, 171, 8, 22, 11, 22, 12, 22, 172, 1, 22, 1, 22,
		0, 0, 23, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19,
		10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37,
		19, 39, 20, 41, 21, 43, 22, 45, 23, 1, 0, 4, 3, 0, 65, 90, 95, 95, 97,
		122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 3, 0, 9, 10,
		13, 13, 32, 32, 180, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0,
		0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0,
		0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1,
		0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29,
		1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0,
		37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0,
		0, 45, 1, 0, 0, 0, 1, 47, 1, 0, 0, 0, 3, 49, 1, 0, 0, 0, 5, 51, 1, 0, 0,
		0, 7, 53, 1, 0, 0, 0, 9, 57, 1, 0, 0, 0, 11, 61, 1, 0, 0, 0, 13, 63, 1,
		0, 0, 0, 15, 66, 1, 0, 0, 0, 17, 75, 1, 0, 0, 0, 19, 84, 1, 0, 0, 0, 21,
		93, 1, 0, 0, 0, 23, 102, 1, 0, 0, 0, 25, 109, 1, 0, 0, 0, 27, 116, 1, 0,
		0, 0, 29, 122, 1, 0, 0, 0, 31, 129, 1, 0, 0, 0, 33, 136, 1, 0, 0, 0, 35,
		142, 1, 0, 0, 0, 37, 144, 1, 0, 0, 0, 39, 146, 1, 0, 0, 0, 41, 154, 1,
		0, 0, 0, 43, 159, 1, 0, 0, 0, 45, 170, 1, 0, 0, 0, 47, 48, 5, 58, 0, 0,
		48, 2, 1, 0, 0, 0, 49, 50, 5, 91, 0, 0, 50, 4, 1, 0, 0, 0, 51, 52, 5, 93,
		0, 0, 52, 6, 1, 0, 0, 0, 53, 54, 5, 99, 0, 0, 54, 55, 5, 109, 0, 0, 55,
		56, 5, 100, 0, 0, 56, 8, 1, 0, 0, 0, 57, 58, 5, 97, 0, 0, 58, 59, 5, 100,
		0, 0, 59, 60, 5, 100, 0, 0, 60, 10, 1, 0, 0, 0, 61, 62, 5, 44, 0, 0, 62,
		12, 1, 0, 0, 0, 63, 64, 5, 105, 0, 0, 64, 65, 5, 100, 0, 0, 65, 14, 1,
		0, 0, 0, 66, 67, 5, 111, 0, 0, 67, 68, 5, 112, 0, 0, 68, 69, 5, 101, 0,
		0, 69, 70, 5, 114, 0, 0, 70, 71, 5, 97, 0, 0, 71, 72, 5, 110, 0, 0, 72,
		73, 5, 100, 0, 0, 73, 74, 5, 49, 0, 0, 74, 16, 1, 0, 0, 0, 75, 76, 5, 111,
		0, 0, 76, 77, 5, 112, 0, 0, 77, 78, 5, 101, 0, 0, 78, 79, 5, 114, 0, 0,
		79, 80, 5, 97, 0, 0, 80, 81, 5, 110, 0, 0, 81, 82, 5, 100, 0, 0, 82, 83,
		5, 50, 0, 0, 83, 18, 1, 0, 0, 0, 84, 85, 5, 115, 0, 0, 85, 86, 5, 117,
		0, 0, 86, 87, 5, 98, 0, 0, 87, 88, 5, 116, 0, 0, 88, 89, 5, 114, 0, 0,
		89, 90, 5, 97, 0, 0, 90, 91, 5, 99, 0, 0, 91, 92, 5, 116, 0, 0, 92, 20,
		1, 0, 0, 0, 93, 94, 5, 109, 0, 0, 94, 95, 5, 117, 0, 0, 95, 96, 5, 108,
		0, 0, 96, 97, 5, 116, 0, 0, 97, 98, 5, 105, 0, 0, 98, 99, 5, 112, 0, 0,
		99, 100, 5, 108, 0, 0, 100, 101, 5, 121, 0, 0, 101, 22, 1, 0, 0, 0, 102,
		103, 5, 100, 0, 0, 103, 104, 5, 105, 0, 0, 104, 105, 5, 118, 0, 0, 105,
		106, 5, 105, 0, 0, 106, 107, 5, 100, 0, 0, 107, 108, 5, 101, 0, 0, 108,
		24, 1, 0, 0, 0, 109, 110, 5, 99, 0, 0, 110, 111, 5, 114, 0, 0, 111, 112,
		5, 101, 0, 0, 112, 113, 5, 97, 0, 0, 113, 114, 5, 116, 0, 0, 114, 115,
		5, 101, 0, 0, 115, 26, 1, 0, 0, 0, 116, 117, 5, 118, 0, 0, 117, 118, 5,
		97, 0, 0, 118, 119, 5, 108, 0, 0, 119, 120, 5, 117, 0, 0, 120, 121, 5,
		101, 0, 0, 121, 28, 1, 0, 0, 0, 122, 123, 5, 117, 0, 0, 123, 124, 5, 112,
		0, 0, 124, 125, 5, 100, 0, 0, 125, 126, 5, 97, 0, 0, 126, 127, 5, 116,
		0, 0, 127, 128, 5, 101, 0, 0, 128, 30, 1, 0, 0, 0, 129, 130, 5, 100, 0,
		0, 130, 131, 5, 101, 0, 0, 131, 132, 5, 108, 0, 0, 132, 133, 5, 101, 0,
		0, 133, 134, 5, 116, 0, 0, 134, 135, 5, 101, 0, 0, 135, 32, 1, 0, 0, 0,
		136, 137, 5, 112, 0, 0, 137, 138, 5, 114, 0, 0, 138, 139, 5, 105, 0, 0,
		139, 140, 5, 110, 0, 0, 140, 141, 5, 116, 0, 0, 141, 34, 1, 0, 0, 0, 142,
		143, 5, 35, 0, 0, 143, 36, 1, 0, 0, 0, 144, 145, 5, 36, 0, 0, 145, 38,
		1, 0, 0, 0, 146, 150, 7, 0, 0, 0, 147, 149, 7, 1, 0, 0, 148, 147, 1, 0,
		0, 0, 149, 152, 1, 0, 0, 0, 150, 148, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0,
		151, 40, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0, 153, 155, 7, 2, 0, 0, 154, 153,
		1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156, 154, 1, 0, 0, 0, 156, 157, 1, 0,
		0, 0, 157, 42, 1, 0, 0, 0, 158, 160, 7, 2, 0, 0, 159, 158, 1, 0, 0, 0,
		160, 161, 1, 0, 0, 0, 161, 159, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162,
		163, 1, 0, 0, 0, 163, 165, 5, 46, 0, 0, 164, 166, 7, 2, 0, 0, 165, 164,
		1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 165, 1, 0, 0, 0, 167, 168, 1, 0,
		0, 0, 168, 44, 1, 0, 0, 0, 169, 171, 7, 3, 0, 0, 170, 169, 1, 0, 0, 0,
		171, 172, 1, 0, 0, 0, 172, 170, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173,
		174, 1, 0, 0, 0, 174, 175, 6, 22, 0, 0, 175, 46, 1, 0, 0, 0, 6, 0, 150,
		156, 161, 167, 172, 1, 6, 0, 0,
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

// FSLLexerInit initializes any static state used to implement FSLLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewFSLLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func FSLLexerInit() {
	staticData := &FSLLexerLexerStaticData
	staticData.once.Do(fsllexerLexerInit)
}

// NewFSLLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewFSLLexer(input antlr.CharStream) *FSLLexer {
	FSLLexerInit()
	l := new(FSLLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &FSLLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "FSL.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// FSLLexer tokens.
const (
	FSLLexerT__0  = 1
	FSLLexerT__1  = 2
	FSLLexerT__2  = 3
	FSLLexerT__3  = 4
	FSLLexerT__4  = 5
	FSLLexerT__5  = 6
	FSLLexerT__6  = 7
	FSLLexerT__7  = 8
	FSLLexerT__8  = 9
	FSLLexerT__9  = 10
	FSLLexerT__10 = 11
	FSLLexerT__11 = 12
	FSLLexerT__12 = 13
	FSLLexerT__13 = 14
	FSLLexerT__14 = 15
	FSLLexerT__15 = 16
	FSLLexerT__16 = 17
	FSLLexerT__17 = 18
	FSLLexerT__18 = 19
	FSLLexerID    = 20
	FSLLexerINT   = 21
	FSLLexerFLOAT = 22
	FSLLexerWS    = 23
)
