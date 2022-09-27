// Code generated from ApiParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type ApiParserLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var apiparserlexerLexerStaticData struct {
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

func apiparserlexerLexerInit() {
	staticData := &apiparserlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'type'", "'struct'", "'{'", "'}'", "'('", "')'", "'returns'", "'-'",
		"'/'", "'/:'", "'@handler'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "ATHANDLER", "WS", "RAW_STRING",
		"ID",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "ATHANDLER", "WS", "RAW_STRING", "ID", "EscapeSequence", "LetterOrDigit",
		"Digits", "Digit", "Letter",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 14, 127, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 4, 11, 85, 8, 11, 11, 11, 12, 11,
		86, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 4, 12, 94, 8, 12, 11, 12, 12, 12,
		95, 1, 12, 1, 12, 1, 13, 1, 13, 5, 13, 102, 8, 13, 10, 13, 12, 13, 105,
		9, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 3, 15, 112, 8, 15, 1, 16, 1,
		16, 5, 16, 116, 8, 16, 10, 16, 12, 16, 119, 9, 16, 1, 16, 3, 16, 122, 8,
		16, 1, 17, 1, 17, 1, 18, 1, 18, 0, 0, 19, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5,
		11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29,
		0, 31, 0, 33, 0, 35, 0, 37, 0, 1, 0, 6, 3, 0, 9, 10, 12, 13, 32, 32, 4,
		0, 10, 10, 13, 13, 92, 92, 96, 96, 8, 0, 34, 34, 39, 39, 92, 92, 98, 98,
		102, 102, 110, 110, 114, 114, 116, 116, 1, 0, 48, 57, 2, 0, 48, 57, 95,
		95, 3, 0, 65, 90, 95, 95, 97, 122, 128, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0,
		0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0,
		0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0,
		0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1,
		0, 0, 0, 1, 39, 1, 0, 0, 0, 3, 44, 1, 0, 0, 0, 5, 51, 1, 0, 0, 0, 7, 53,
		1, 0, 0, 0, 9, 55, 1, 0, 0, 0, 11, 57, 1, 0, 0, 0, 13, 59, 1, 0, 0, 0,
		15, 67, 1, 0, 0, 0, 17, 69, 1, 0, 0, 0, 19, 71, 1, 0, 0, 0, 21, 74, 1,
		0, 0, 0, 23, 84, 1, 0, 0, 0, 25, 90, 1, 0, 0, 0, 27, 99, 1, 0, 0, 0, 29,
		106, 1, 0, 0, 0, 31, 111, 1, 0, 0, 0, 33, 113, 1, 0, 0, 0, 35, 123, 1,
		0, 0, 0, 37, 125, 1, 0, 0, 0, 39, 40, 5, 116, 0, 0, 40, 41, 5, 121, 0,
		0, 41, 42, 5, 112, 0, 0, 42, 43, 5, 101, 0, 0, 43, 2, 1, 0, 0, 0, 44, 45,
		5, 115, 0, 0, 45, 46, 5, 116, 0, 0, 46, 47, 5, 114, 0, 0, 47, 48, 5, 117,
		0, 0, 48, 49, 5, 99, 0, 0, 49, 50, 5, 116, 0, 0, 50, 4, 1, 0, 0, 0, 51,
		52, 5, 123, 0, 0, 52, 6, 1, 0, 0, 0, 53, 54, 5, 125, 0, 0, 54, 8, 1, 0,
		0, 0, 55, 56, 5, 40, 0, 0, 56, 10, 1, 0, 0, 0, 57, 58, 5, 41, 0, 0, 58,
		12, 1, 0, 0, 0, 59, 60, 5, 114, 0, 0, 60, 61, 5, 101, 0, 0, 61, 62, 5,
		116, 0, 0, 62, 63, 5, 117, 0, 0, 63, 64, 5, 114, 0, 0, 64, 65, 5, 110,
		0, 0, 65, 66, 5, 115, 0, 0, 66, 14, 1, 0, 0, 0, 67, 68, 5, 45, 0, 0, 68,
		16, 1, 0, 0, 0, 69, 70, 5, 47, 0, 0, 70, 18, 1, 0, 0, 0, 71, 72, 5, 47,
		0, 0, 72, 73, 5, 58, 0, 0, 73, 20, 1, 0, 0, 0, 74, 75, 5, 64, 0, 0, 75,
		76, 5, 104, 0, 0, 76, 77, 5, 97, 0, 0, 77, 78, 5, 110, 0, 0, 78, 79, 5,
		100, 0, 0, 79, 80, 5, 108, 0, 0, 80, 81, 5, 101, 0, 0, 81, 82, 5, 114,
		0, 0, 82, 22, 1, 0, 0, 0, 83, 85, 7, 0, 0, 0, 84, 83, 1, 0, 0, 0, 85, 86,
		1, 0, 0, 0, 86, 84, 1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0,
		88, 89, 6, 11, 0, 0, 89, 24, 1, 0, 0, 0, 90, 93, 5, 96, 0, 0, 91, 94, 8,
		1, 0, 0, 92, 94, 3, 29, 14, 0, 93, 91, 1, 0, 0, 0, 93, 92, 1, 0, 0, 0,
		94, 95, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 97, 1,
		0, 0, 0, 97, 98, 5, 96, 0, 0, 98, 26, 1, 0, 0, 0, 99, 103, 3, 37, 18, 0,
		100, 102, 3, 31, 15, 0, 101, 100, 1, 0, 0, 0, 102, 105, 1, 0, 0, 0, 103,
		101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 28, 1, 0, 0, 0, 105, 103, 1,
		0, 0, 0, 106, 107, 5, 92, 0, 0, 107, 108, 7, 2, 0, 0, 108, 30, 1, 0, 0,
		0, 109, 112, 3, 37, 18, 0, 110, 112, 3, 35, 17, 0, 111, 109, 1, 0, 0, 0,
		111, 110, 1, 0, 0, 0, 112, 32, 1, 0, 0, 0, 113, 121, 7, 3, 0, 0, 114, 116,
		7, 4, 0, 0, 115, 114, 1, 0, 0, 0, 116, 119, 1, 0, 0, 0, 117, 115, 1, 0,
		0, 0, 117, 118, 1, 0, 0, 0, 118, 120, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0,
		120, 122, 7, 3, 0, 0, 121, 117, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122,
		34, 1, 0, 0, 0, 123, 124, 7, 3, 0, 0, 124, 36, 1, 0, 0, 0, 125, 126, 7,
		5, 0, 0, 126, 38, 1, 0, 0, 0, 8, 0, 86, 93, 95, 103, 111, 117, 121, 1,
		0, 88, 0,
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

// ApiParserLexerInit initializes any static state used to implement ApiParserLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewApiParserLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func ApiParserLexerInit() {
	staticData := &apiparserlexerLexerStaticData
	staticData.once.Do(apiparserlexerLexerInit)
}

// NewApiParserLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewApiParserLexer(input antlr.CharStream) *ApiParserLexer {
	ApiParserLexerInit()
	l := new(ApiParserLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &apiparserlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "ApiParser.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ApiParserLexer tokens.
const (
	ApiParserLexerT__0       = 1
	ApiParserLexerT__1       = 2
	ApiParserLexerT__2       = 3
	ApiParserLexerT__3       = 4
	ApiParserLexerT__4       = 5
	ApiParserLexerT__5       = 6
	ApiParserLexerT__6       = 7
	ApiParserLexerT__7       = 8
	ApiParserLexerT__8       = 9
	ApiParserLexerT__9       = 10
	ApiParserLexerATHANDLER  = 11
	ApiParserLexerWS         = 12
	ApiParserLexerRAW_STRING = 13
	ApiParserLexerID         = 14
)
