// Code generated from FSL.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // FSL
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type FSLParser struct {
	*antlr.BaseParser
}

var FSLParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func fslParserInit() {
	staticData := &FSLParserStaticData
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
		"program", "statement", "variableDeclaration", "functionDefinition",
		"command", "arithmeticOperation", "crudOperation", "addOperation", "subtractOperation",
		"multiplyOperation", "divideOperation", "createOperation", "updateOperation",
		"deleteOperation", "printOperation", "functionCall", "passedParameter",
		"operationParameter", "identificationParameter", "functionParameter",
		"variableReference", "value",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 23, 214, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 1, 0, 5, 0, 46, 8, 0, 10, 0, 12, 0, 49, 9, 0, 1, 1, 1, 1, 1,
		1, 3, 1, 54, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3,
		64, 8, 3, 10, 3, 12, 3, 67, 9, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 3, 4, 77, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 83, 8, 5, 1, 6,
		1, 6, 1, 6, 3, 6, 88, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 3, 12, 166, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 180, 8, 14, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 188, 8, 15, 10, 15, 12, 15, 191, 9,
		15, 1, 16, 1, 16, 3, 16, 195, 8, 16, 1, 17, 1, 17, 1, 17, 3, 17, 200, 8,
		17, 1, 18, 1, 18, 3, 18, 204, 8, 18, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20,
		1, 20, 1, 21, 1, 21, 1, 21, 0, 0, 22, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18,
		20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 0, 2, 2, 0, 7, 7, 20, 20,
		1, 0, 21, 22, 210, 0, 47, 1, 0, 0, 0, 2, 53, 1, 0, 0, 0, 4, 55, 1, 0, 0,
		0, 6, 59, 1, 0, 0, 0, 8, 70, 1, 0, 0, 0, 10, 82, 1, 0, 0, 0, 12, 87, 1,
		0, 0, 0, 14, 89, 1, 0, 0, 0, 16, 103, 1, 0, 0, 0, 18, 117, 1, 0, 0, 0,
		20, 131, 1, 0, 0, 0, 22, 145, 1, 0, 0, 0, 24, 155, 1, 0, 0, 0, 26, 167,
		1, 0, 0, 0, 28, 173, 1, 0, 0, 0, 30, 181, 1, 0, 0, 0, 32, 194, 1, 0, 0,
		0, 34, 199, 1, 0, 0, 0, 36, 203, 1, 0, 0, 0, 38, 205, 1, 0, 0, 0, 40, 208,
		1, 0, 0, 0, 42, 211, 1, 0, 0, 0, 44, 46, 3, 2, 1, 0, 45, 44, 1, 0, 0, 0,
		46, 49, 1, 0, 0, 0, 47, 45, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 1, 1, 0,
		0, 0, 49, 47, 1, 0, 0, 0, 50, 54, 3, 4, 2, 0, 51, 54, 3, 6, 3, 0, 52, 54,
		3, 8, 4, 0, 53, 50, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 53, 52, 1, 0, 0, 0,
		54, 3, 1, 0, 0, 0, 55, 56, 5, 20, 0, 0, 56, 57, 5, 1, 0, 0, 57, 58, 3,
		42, 21, 0, 58, 5, 1, 0, 0, 0, 59, 60, 5, 20, 0, 0, 60, 61, 5, 1, 0, 0,
		61, 65, 5, 2, 0, 0, 62, 64, 3, 8, 4, 0, 63, 62, 1, 0, 0, 0, 64, 67, 1,
		0, 0, 0, 65, 63, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 68, 1, 0, 0, 0, 67,
		65, 1, 0, 0, 0, 68, 69, 5, 3, 0, 0, 69, 7, 1, 0, 0, 0, 70, 71, 5, 4, 0,
		0, 71, 76, 5, 1, 0, 0, 72, 77, 3, 30, 15, 0, 73, 77, 3, 10, 5, 0, 74, 77,
		3, 12, 6, 0, 75, 77, 3, 28, 14, 0, 76, 72, 1, 0, 0, 0, 76, 73, 1, 0, 0,
		0, 76, 74, 1, 0, 0, 0, 76, 75, 1, 0, 0, 0, 77, 9, 1, 0, 0, 0, 78, 83, 3,
		14, 7, 0, 79, 83, 3, 16, 8, 0, 80, 83, 3, 18, 9, 0, 81, 83, 3, 20, 10,
		0, 82, 78, 1, 0, 0, 0, 82, 79, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0, 82, 81,
		1, 0, 0, 0, 83, 11, 1, 0, 0, 0, 84, 88, 3, 22, 11, 0, 85, 88, 3, 24, 12,
		0, 86, 88, 3, 26, 13, 0, 87, 84, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 87, 86,
		1, 0, 0, 0, 88, 13, 1, 0, 0, 0, 89, 90, 5, 5, 0, 0, 90, 91, 5, 6, 0, 0,
		91, 92, 5, 7, 0, 0, 92, 93, 5, 1, 0, 0, 93, 94, 3, 36, 18, 0, 94, 95, 5,
		6, 0, 0, 95, 96, 5, 8, 0, 0, 96, 97, 5, 1, 0, 0, 97, 98, 3, 34, 17, 0,
		98, 99, 5, 6, 0, 0, 99, 100, 5, 9, 0, 0, 100, 101, 5, 1, 0, 0, 101, 102,
		3, 34, 17, 0, 102, 15, 1, 0, 0, 0, 103, 104, 5, 10, 0, 0, 104, 105, 5,
		6, 0, 0, 105, 106, 5, 7, 0, 0, 106, 107, 5, 1, 0, 0, 107, 108, 3, 36, 18,
		0, 108, 109, 5, 6, 0, 0, 109, 110, 5, 8, 0, 0, 110, 111, 5, 1, 0, 0, 111,
		112, 3, 34, 17, 0, 112, 113, 5, 6, 0, 0, 113, 114, 5, 9, 0, 0, 114, 115,
		5, 1, 0, 0, 115, 116, 3, 34, 17, 0, 116, 17, 1, 0, 0, 0, 117, 118, 5, 11,
		0, 0, 118, 119, 5, 6, 0, 0, 119, 120, 5, 7, 0, 0, 120, 121, 5, 1, 0, 0,
		121, 122, 3, 36, 18, 0, 122, 123, 5, 6, 0, 0, 123, 124, 5, 8, 0, 0, 124,
		125, 5, 1, 0, 0, 125, 126, 3, 34, 17, 0, 126, 127, 5, 6, 0, 0, 127, 128,
		5, 9, 0, 0, 128, 129, 5, 1, 0, 0, 129, 130, 3, 34, 17, 0, 130, 19, 1, 0,
		0, 0, 131, 132, 5, 12, 0, 0, 132, 133, 5, 6, 0, 0, 133, 134, 5, 7, 0, 0,
		134, 135, 5, 1, 0, 0, 135, 136, 3, 36, 18, 0, 136, 137, 5, 6, 0, 0, 137,
		138, 5, 8, 0, 0, 138, 139, 5, 1, 0, 0, 139, 140, 3, 34, 17, 0, 140, 141,
		5, 6, 0, 0, 141, 142, 5, 9, 0, 0, 142, 143, 5, 1, 0, 0, 143, 144, 3, 34,
		17, 0, 144, 21, 1, 0, 0, 0, 145, 146, 5, 13, 0, 0, 146, 147, 5, 6, 0, 0,
		147, 148, 5, 7, 0, 0, 148, 149, 5, 1, 0, 0, 149, 150, 5, 20, 0, 0, 150,
		151, 5, 6, 0, 0, 151, 152, 5, 14, 0, 0, 152, 153, 5, 1, 0, 0, 153, 154,
		3, 42, 21, 0, 154, 23, 1, 0, 0, 0, 155, 156, 5, 15, 0, 0, 156, 157, 5,
		6, 0, 0, 157, 158, 5, 7, 0, 0, 158, 159, 5, 1, 0, 0, 159, 160, 5, 20, 0,
		0, 160, 161, 5, 6, 0, 0, 161, 162, 5, 14, 0, 0, 162, 165, 5, 1, 0, 0, 163,
		166, 3, 42, 21, 0, 164, 166, 3, 40, 20, 0, 165, 163, 1, 0, 0, 0, 165, 164,
		1, 0, 0, 0, 166, 25, 1, 0, 0, 0, 167, 168, 5, 16, 0, 0, 168, 169, 5, 6,
		0, 0, 169, 170, 5, 7, 0, 0, 170, 171, 5, 1, 0, 0, 171, 172, 5, 20, 0, 0,
		172, 27, 1, 0, 0, 0, 173, 174, 5, 17, 0, 0, 174, 175, 5, 6, 0, 0, 175,
		176, 5, 14, 0, 0, 176, 179, 5, 1, 0, 0, 177, 180, 3, 42, 21, 0, 178, 180,
		3, 40, 20, 0, 179, 177, 1, 0, 0, 0, 179, 178, 1, 0, 0, 0, 180, 29, 1, 0,
		0, 0, 181, 182, 5, 18, 0, 0, 182, 189, 5, 20, 0, 0, 183, 184, 5, 6, 0,
		0, 184, 185, 7, 0, 0, 0, 185, 186, 5, 1, 0, 0, 186, 188, 3, 32, 16, 0,
		187, 183, 1, 0, 0, 0, 188, 191, 1, 0, 0, 0, 189, 187, 1, 0, 0, 0, 189,
		190, 1, 0, 0, 0, 190, 31, 1, 0, 0, 0, 191, 189, 1, 0, 0, 0, 192, 195, 5,
		20, 0, 0, 193, 195, 3, 40, 20, 0, 194, 192, 1, 0, 0, 0, 194, 193, 1, 0,
		0, 0, 195, 33, 1, 0, 0, 0, 196, 200, 3, 40, 20, 0, 197, 200, 3, 42, 21,
		0, 198, 200, 3, 38, 19, 0, 199, 196, 1, 0, 0, 0, 199, 197, 1, 0, 0, 0,
		199, 198, 1, 0, 0, 0, 200, 35, 1, 0, 0, 0, 201, 204, 5, 20, 0, 0, 202,
		204, 3, 38, 19, 0, 203, 201, 1, 0, 0, 0, 203, 202, 1, 0, 0, 0, 204, 37,
		1, 0, 0, 0, 205, 206, 5, 19, 0, 0, 206, 207, 7, 0, 0, 0, 207, 39, 1, 0,
		0, 0, 208, 209, 5, 18, 0, 0, 209, 210, 5, 20, 0, 0, 210, 41, 1, 0, 0, 0,
		211, 212, 7, 1, 0, 0, 212, 43, 1, 0, 0, 0, 12, 47, 53, 65, 76, 82, 87,
		165, 179, 189, 194, 199, 203,
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

// FSLParserInit initializes any static state used to implement FSLParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewFSLParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func FSLParserInit() {
	staticData := &FSLParserStaticData
	staticData.once.Do(fslParserInit)
}

// NewFSLParser produces a new parser instance for the optional input antlr.TokenStream.
func NewFSLParser(input antlr.TokenStream) *FSLParser {
	FSLParserInit()
	this := new(FSLParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &FSLParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "FSL.g4"

	return this
}

// FSLParser tokens.
const (
	FSLParserEOF   = antlr.TokenEOF
	FSLParserT__0  = 1
	FSLParserT__1  = 2
	FSLParserT__2  = 3
	FSLParserT__3  = 4
	FSLParserT__4  = 5
	FSLParserT__5  = 6
	FSLParserT__6  = 7
	FSLParserT__7  = 8
	FSLParserT__8  = 9
	FSLParserT__9  = 10
	FSLParserT__10 = 11
	FSLParserT__11 = 12
	FSLParserT__12 = 13
	FSLParserT__13 = 14
	FSLParserT__14 = 15
	FSLParserT__15 = 16
	FSLParserT__16 = 17
	FSLParserT__17 = 18
	FSLParserT__18 = 19
	FSLParserID    = 20
	FSLParserINT   = 21
	FSLParserFLOAT = 22
	FSLParserWS    = 23
)

// FSLParser rules.
const (
	FSLParserRULE_program                 = 0
	FSLParserRULE_statement               = 1
	FSLParserRULE_variableDeclaration     = 2
	FSLParserRULE_functionDefinition      = 3
	FSLParserRULE_command                 = 4
	FSLParserRULE_arithmeticOperation     = 5
	FSLParserRULE_crudOperation           = 6
	FSLParserRULE_addOperation            = 7
	FSLParserRULE_subtractOperation       = 8
	FSLParserRULE_multiplyOperation       = 9
	FSLParserRULE_divideOperation         = 10
	FSLParserRULE_createOperation         = 11
	FSLParserRULE_updateOperation         = 12
	FSLParserRULE_deleteOperation         = 13
	FSLParserRULE_printOperation          = 14
	FSLParserRULE_functionCall            = 15
	FSLParserRULE_passedParameter         = 16
	FSLParserRULE_operationParameter      = 17
	FSLParserRULE_identificationParameter = 18
	FSLParserRULE_functionParameter       = 19
	FSLParserRULE_variableReference       = 20
	FSLParserRULE_value                   = 21
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FSLParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *FSLParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FSLParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FSLParserT__3 || _la == FSLParserID {
		{
			p.SetState(44)
			p.Statement()
		}

		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VariableDeclaration() IVariableDeclarationContext
	FunctionDefinition() IFunctionDefinitionContext
	Command() ICommandContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FSLParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) VariableDeclaration() IVariableDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *StatementContext) FunctionDefinition() IFunctionDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionDefinitionContext)
}

func (s *StatementContext) Command() ICommandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommandContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *FSLParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FSLParserRULE_statement)
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(50)
			p.VariableDeclaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(51)
			p.FunctionDefinition()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(52)
			p.Command()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariableDeclarationContext is an interface to support dynamic dispatch.
type IVariableDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Value() IValueContext

	// IsVariableDeclarationContext differentiates from other interfaces.
	IsVariableDeclarationContext()
}

type VariableDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclarationContext() *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_variableDeclaration
	return p
}

func InitEmptyVariableDeclarationContext(p *VariableDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_variableDeclaration
}

func (*VariableDeclarationContext) IsVariableDeclarationContext() {}

func NewVariableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FSLParserRULE_variableDeclaration

	return p
}

func (s *VariableDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(FSLParserID, 0)
}

func (s *VariableDeclarationContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *VariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.EnterVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.ExitVariableDeclaration(s)
	}
}

func (p *FSLParser) VariableDeclaration() (localctx IVariableDeclarationContext) {
	localctx = NewVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, FSLParserRULE_variableDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(55)
		p.Match(FSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(56)
		p.Match(FSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(57)
		p.Value()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionDefinitionContext is an interface to support dynamic dispatch.
type IFunctionDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	AllCommand() []ICommandContext
	Command(i int) ICommandContext

	// IsFunctionDefinitionContext differentiates from other interfaces.
	IsFunctionDefinitionContext()
}

type FunctionDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDefinitionContext() *FunctionDefinitionContext {
	var p = new(FunctionDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_functionDefinition
	return p
}

func InitEmptyFunctionDefinitionContext(p *FunctionDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_functionDefinition
}

func (*FunctionDefinitionContext) IsFunctionDefinitionContext() {}

func NewFunctionDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDefinitionContext {
	var p = new(FunctionDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FSLParserRULE_functionDefinition

	return p
}

func (s *FunctionDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDefinitionContext) ID() antlr.TerminalNode {
	return s.GetToken(FSLParserID, 0)
}

func (s *FunctionDefinitionContext) AllCommand() []ICommandContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICommandContext); ok {
			len++
		}
	}

	tst := make([]ICommandContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICommandContext); ok {
			tst[i] = t.(ICommandContext)
			i++
		}
	}

	return tst
}

func (s *FunctionDefinitionContext) Command(i int) ICommandContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommandContext); ok {
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

	return t.(ICommandContext)
}

func (s *FunctionDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.EnterFunctionDefinition(s)
	}
}

func (s *FunctionDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.ExitFunctionDefinition(s)
	}
}

func (p *FSLParser) FunctionDefinition() (localctx IFunctionDefinitionContext) {
	localctx = NewFunctionDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FSLParserRULE_functionDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
		p.Match(FSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(60)
		p.Match(FSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(61)
		p.Match(FSLParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FSLParserT__3 {
		{
			p.SetState(62)
			p.Command()
		}

		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(68)
		p.Match(FSLParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICommandContext is an interface to support dynamic dispatch.
type ICommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FunctionCall() IFunctionCallContext
	ArithmeticOperation() IArithmeticOperationContext
	CrudOperation() ICrudOperationContext
	PrintOperation() IPrintOperationContext

	// IsCommandContext differentiates from other interfaces.
	IsCommandContext()
}

type CommandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandContext() *CommandContext {
	var p = new(CommandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_command
	return p
}

func InitEmptyCommandContext(p *CommandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_command
}

func (*CommandContext) IsCommandContext() {}

func NewCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandContext {
	var p = new(CommandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FSLParserRULE_command

	return p
}

func (s *CommandContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *CommandContext) ArithmeticOperation() IArithmeticOperationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArithmeticOperationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArithmeticOperationContext)
}

func (s *CommandContext) CrudOperation() ICrudOperationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICrudOperationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICrudOperationContext)
}

func (s *CommandContext) PrintOperation() IPrintOperationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintOperationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintOperationContext)
}

func (s *CommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.EnterCommand(s)
	}
}

func (s *CommandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.ExitCommand(s)
	}
}

func (p *FSLParser) Command() (localctx ICommandContext) {
	localctx = NewCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FSLParserRULE_command)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(70)
		p.Match(FSLParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(71)
		p.Match(FSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FSLParserT__17:
		{
			p.SetState(72)
			p.FunctionCall()
		}

	case FSLParserT__4, FSLParserT__9, FSLParserT__10, FSLParserT__11:
		{
			p.SetState(73)
			p.ArithmeticOperation()
		}

	case FSLParserT__12, FSLParserT__14, FSLParserT__15:
		{
			p.SetState(74)
			p.CrudOperation()
		}

	case FSLParserT__16:
		{
			p.SetState(75)
			p.PrintOperation()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArithmeticOperationContext is an interface to support dynamic dispatch.
type IArithmeticOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AddOperation() IAddOperationContext
	SubtractOperation() ISubtractOperationContext
	MultiplyOperation() IMultiplyOperationContext
	DivideOperation() IDivideOperationContext

	// IsArithmeticOperationContext differentiates from other interfaces.
	IsArithmeticOperationContext()
}

type ArithmeticOperationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArithmeticOperationContext() *ArithmeticOperationContext {
	var p = new(ArithmeticOperationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_arithmeticOperation
	return p
}

func InitEmptyArithmeticOperationContext(p *ArithmeticOperationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_arithmeticOperation
}

func (*ArithmeticOperationContext) IsArithmeticOperationContext() {}

func NewArithmeticOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArithmeticOperationContext {
	var p = new(ArithmeticOperationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FSLParserRULE_arithmeticOperation

	return p
}

func (s *ArithmeticOperationContext) GetParser() antlr.Parser { return s.parser }

func (s *ArithmeticOperationContext) AddOperation() IAddOperationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddOperationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddOperationContext)
}

func (s *ArithmeticOperationContext) SubtractOperation() ISubtractOperationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubtractOperationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubtractOperationContext)
}

func (s *ArithmeticOperationContext) MultiplyOperation() IMultiplyOperationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplyOperationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiplyOperationContext)
}

func (s *ArithmeticOperationContext) DivideOperation() IDivideOperationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDivideOperationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDivideOperationContext)
}

func (s *ArithmeticOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArithmeticOperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArithmeticOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.EnterArithmeticOperation(s)
	}
}

func (s *ArithmeticOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.ExitArithmeticOperation(s)
	}
}

func (p *FSLParser) ArithmeticOperation() (localctx IArithmeticOperationContext) {
	localctx = NewArithmeticOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FSLParserRULE_arithmeticOperation)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FSLParserT__4:
		{
			p.SetState(78)
			p.AddOperation()
		}

	case FSLParserT__9:
		{
			p.SetState(79)
			p.SubtractOperation()
		}

	case FSLParserT__10:
		{
			p.SetState(80)
			p.MultiplyOperation()
		}

	case FSLParserT__11:
		{
			p.SetState(81)
			p.DivideOperation()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICrudOperationContext is an interface to support dynamic dispatch.
type ICrudOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CreateOperation() ICreateOperationContext
	UpdateOperation() IUpdateOperationContext
	DeleteOperation() IDeleteOperationContext

	// IsCrudOperationContext differentiates from other interfaces.
	IsCrudOperationContext()
}

type CrudOperationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCrudOperationContext() *CrudOperationContext {
	var p = new(CrudOperationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_crudOperation
	return p
}

func InitEmptyCrudOperationContext(p *CrudOperationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_crudOperation
}

func (*CrudOperationContext) IsCrudOperationContext() {}

func NewCrudOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CrudOperationContext {
	var p = new(CrudOperationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FSLParserRULE_crudOperation

	return p
}

func (s *CrudOperationContext) GetParser() antlr.Parser { return s.parser }

func (s *CrudOperationContext) CreateOperation() ICreateOperationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICreateOperationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICreateOperationContext)
}

func (s *CrudOperationContext) UpdateOperation() IUpdateOperationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUpdateOperationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUpdateOperationContext)
}

func (s *CrudOperationContext) DeleteOperation() IDeleteOperationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeleteOperationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeleteOperationContext)
}

func (s *CrudOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CrudOperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CrudOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.EnterCrudOperation(s)
	}
}

func (s *CrudOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.ExitCrudOperation(s)
	}
}

func (p *FSLParser) CrudOperation() (localctx ICrudOperationContext) {
	localctx = NewCrudOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, FSLParserRULE_crudOperation)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FSLParserT__12:
		{
			p.SetState(84)
			p.CreateOperation()
		}

	case FSLParserT__14:
		{
			p.SetState(85)
			p.UpdateOperation()
		}

	case FSLParserT__15:
		{
			p.SetState(86)
			p.DeleteOperation()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAddOperationContext is an interface to support dynamic dispatch.
type IAddOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IdentificationParameter() IIdentificationParameterContext
	AllOperationParameter() []IOperationParameterContext
	OperationParameter(i int) IOperationParameterContext

	// IsAddOperationContext differentiates from other interfaces.
	IsAddOperationContext()
}

type AddOperationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddOperationContext() *AddOperationContext {
	var p = new(AddOperationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_addOperation
	return p
}

func InitEmptyAddOperationContext(p *AddOperationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_addOperation
}

func (*AddOperationContext) IsAddOperationContext() {}

func NewAddOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddOperationContext {
	var p = new(AddOperationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FSLParserRULE_addOperation

	return p
}

func (s *AddOperationContext) GetParser() antlr.Parser { return s.parser }

func (s *AddOperationContext) IdentificationParameter() IIdentificationParameterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentificationParameterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentificationParameterContext)
}

func (s *AddOperationContext) AllOperationParameter() []IOperationParameterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOperationParameterContext); ok {
			len++
		}
	}

	tst := make([]IOperationParameterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOperationParameterContext); ok {
			tst[i] = t.(IOperationParameterContext)
			i++
		}
	}

	return tst
}

func (s *AddOperationContext) OperationParameter(i int) IOperationParameterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperationParameterContext); ok {
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

	return t.(IOperationParameterContext)
}

func (s *AddOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddOperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.EnterAddOperation(s)
	}
}

func (s *AddOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.ExitAddOperation(s)
	}
}

func (p *FSLParser) AddOperation() (localctx IAddOperationContext) {
	localctx = NewAddOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, FSLParserRULE_addOperation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(89)
		p.Match(FSLParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(90)
		p.Match(FSLParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(91)
		p.Match(FSLParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(92)
		p.Match(FSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(93)
		p.IdentificationParameter()
	}
	{
		p.SetState(94)
		p.Match(FSLParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(95)
		p.Match(FSLParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(96)
		p.Match(FSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(97)
		p.OperationParameter()
	}
	{
		p.SetState(98)
		p.Match(FSLParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(99)
		p.Match(FSLParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(100)
		p.Match(FSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(101)
		p.OperationParameter()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISubtractOperationContext is an interface to support dynamic dispatch.
type ISubtractOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IdentificationParameter() IIdentificationParameterContext
	AllOperationParameter() []IOperationParameterContext
	OperationParameter(i int) IOperationParameterContext

	// IsSubtractOperationContext differentiates from other interfaces.
	IsSubtractOperationContext()
}

type SubtractOperationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubtractOperationContext() *SubtractOperationContext {
	var p = new(SubtractOperationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_subtractOperation
	return p
}

func InitEmptySubtractOperationContext(p *SubtractOperationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_subtractOperation
}

func (*SubtractOperationContext) IsSubtractOperationContext() {}

func NewSubtractOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubtractOperationContext {
	var p = new(SubtractOperationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FSLParserRULE_subtractOperation

	return p
}

func (s *SubtractOperationContext) GetParser() antlr.Parser { return s.parser }

func (s *SubtractOperationContext) IdentificationParameter() IIdentificationParameterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentificationParameterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentificationParameterContext)
}

func (s *SubtractOperationContext) AllOperationParameter() []IOperationParameterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOperationParameterContext); ok {
			len++
		}
	}

	tst := make([]IOperationParameterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOperationParameterContext); ok {
			tst[i] = t.(IOperationParameterContext)
			i++
		}
	}

	return tst
}

func (s *SubtractOperationContext) OperationParameter(i int) IOperationParameterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperationParameterContext); ok {
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

	return t.(IOperationParameterContext)
}

func (s *SubtractOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubtractOperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubtractOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.EnterSubtractOperation(s)
	}
}

func (s *SubtractOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.ExitSubtractOperation(s)
	}
}

func (p *FSLParser) SubtractOperation() (localctx ISubtractOperationContext) {
	localctx = NewSubtractOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, FSLParserRULE_subtractOperation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(103)
		p.Match(FSLParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(104)
		p.Match(FSLParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(105)
		p.Match(FSLParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)
		p.Match(FSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(107)
		p.IdentificationParameter()
	}
	{
		p.SetState(108)
		p.Match(FSLParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(109)
		p.Match(FSLParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(110)
		p.Match(FSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(111)
		p.OperationParameter()
	}
	{
		p.SetState(112)
		p.Match(FSLParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(113)
		p.Match(FSLParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(114)
		p.Match(FSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(115)
		p.OperationParameter()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMultiplyOperationContext is an interface to support dynamic dispatch.
type IMultiplyOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IdentificationParameter() IIdentificationParameterContext
	AllOperationParameter() []IOperationParameterContext
	OperationParameter(i int) IOperationParameterContext

	// IsMultiplyOperationContext differentiates from other interfaces.
	IsMultiplyOperationContext()
}

type MultiplyOperationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplyOperationContext() *MultiplyOperationContext {
	var p = new(MultiplyOperationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_multiplyOperation
	return p
}

func InitEmptyMultiplyOperationContext(p *MultiplyOperationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_multiplyOperation
}

func (*MultiplyOperationContext) IsMultiplyOperationContext() {}

func NewMultiplyOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplyOperationContext {
	var p = new(MultiplyOperationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FSLParserRULE_multiplyOperation

	return p
}

func (s *MultiplyOperationContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplyOperationContext) IdentificationParameter() IIdentificationParameterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentificationParameterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentificationParameterContext)
}

func (s *MultiplyOperationContext) AllOperationParameter() []IOperationParameterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOperationParameterContext); ok {
			len++
		}
	}

	tst := make([]IOperationParameterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOperationParameterContext); ok {
			tst[i] = t.(IOperationParameterContext)
			i++
		}
	}

	return tst
}

func (s *MultiplyOperationContext) OperationParameter(i int) IOperationParameterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperationParameterContext); ok {
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

	return t.(IOperationParameterContext)
}

func (s *MultiplyOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplyOperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplyOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.EnterMultiplyOperation(s)
	}
}

func (s *MultiplyOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.ExitMultiplyOperation(s)
	}
}

func (p *FSLParser) MultiplyOperation() (localctx IMultiplyOperationContext) {
	localctx = NewMultiplyOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, FSLParserRULE_multiplyOperation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		p.Match(FSLParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(118)
		p.Match(FSLParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(119)
		p.Match(FSLParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(120)
		p.Match(FSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(121)
		p.IdentificationParameter()
	}
	{
		p.SetState(122)
		p.Match(FSLParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(123)
		p.Match(FSLParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(124)
		p.Match(FSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(125)
		p.OperationParameter()
	}
	{
		p.SetState(126)
		p.Match(FSLParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(127)
		p.Match(FSLParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(128)
		p.Match(FSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(129)
		p.OperationParameter()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDivideOperationContext is an interface to support dynamic dispatch.
type IDivideOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IdentificationParameter() IIdentificationParameterContext
	AllOperationParameter() []IOperationParameterContext
	OperationParameter(i int) IOperationParameterContext

	// IsDivideOperationContext differentiates from other interfaces.
	IsDivideOperationContext()
}

type DivideOperationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDivideOperationContext() *DivideOperationContext {
	var p = new(DivideOperationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_divideOperation
	return p
}

func InitEmptyDivideOperationContext(p *DivideOperationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_divideOperation
}

func (*DivideOperationContext) IsDivideOperationContext() {}

func NewDivideOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DivideOperationContext {
	var p = new(DivideOperationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FSLParserRULE_divideOperation

	return p
}

func (s *DivideOperationContext) GetParser() antlr.Parser { return s.parser }

func (s *DivideOperationContext) IdentificationParameter() IIdentificationParameterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentificationParameterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentificationParameterContext)
}

func (s *DivideOperationContext) AllOperationParameter() []IOperationParameterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOperationParameterContext); ok {
			len++
		}
	}

	tst := make([]IOperationParameterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOperationParameterContext); ok {
			tst[i] = t.(IOperationParameterContext)
			i++
		}
	}

	return tst
}

func (s *DivideOperationContext) OperationParameter(i int) IOperationParameterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperationParameterContext); ok {
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

	return t.(IOperationParameterContext)
}

func (s *DivideOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DivideOperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DivideOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.EnterDivideOperation(s)
	}
}

func (s *DivideOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.ExitDivideOperation(s)
	}
}

func (p *FSLParser) DivideOperation() (localctx IDivideOperationContext) {
	localctx = NewDivideOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, FSLParserRULE_divideOperation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)
		p.Match(FSLParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(132)
		p.Match(FSLParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(133)
		p.Match(FSLParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(134)
		p.Match(FSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(135)
		p.IdentificationParameter()
	}
	{
		p.SetState(136)
		p.Match(FSLParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(137)
		p.Match(FSLParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(138)
		p.Match(FSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(139)
		p.OperationParameter()
	}
	{
		p.SetState(140)
		p.Match(FSLParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(141)
		p.Match(FSLParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(142)
		p.Match(FSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(143)
		p.OperationParameter()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICreateOperationContext is an interface to support dynamic dispatch.
type ICreateOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Value() IValueContext

	// IsCreateOperationContext differentiates from other interfaces.
	IsCreateOperationContext()
}

type CreateOperationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreateOperationContext() *CreateOperationContext {
	var p = new(CreateOperationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_createOperation
	return p
}

func InitEmptyCreateOperationContext(p *CreateOperationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_createOperation
}

func (*CreateOperationContext) IsCreateOperationContext() {}

func NewCreateOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateOperationContext {
	var p = new(CreateOperationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FSLParserRULE_createOperation

	return p
}

func (s *CreateOperationContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateOperationContext) ID() antlr.TerminalNode {
	return s.GetToken(FSLParserID, 0)
}

func (s *CreateOperationContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *CreateOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateOperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.EnterCreateOperation(s)
	}
}

func (s *CreateOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.ExitCreateOperation(s)
	}
}

func (p *FSLParser) CreateOperation() (localctx ICreateOperationContext) {
	localctx = NewCreateOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, FSLParserRULE_createOperation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
		p.Match(FSLParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(146)
		p.Match(FSLParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(147)
		p.Match(FSLParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(148)
		p.Match(FSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(149)
		p.Match(FSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(150)
		p.Match(FSLParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(151)
		p.Match(FSLParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(152)
		p.Match(FSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(153)
		p.Value()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUpdateOperationContext is an interface to support dynamic dispatch.
type IUpdateOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Value() IValueContext
	VariableReference() IVariableReferenceContext

	// IsUpdateOperationContext differentiates from other interfaces.
	IsUpdateOperationContext()
}

type UpdateOperationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdateOperationContext() *UpdateOperationContext {
	var p = new(UpdateOperationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_updateOperation
	return p
}

func InitEmptyUpdateOperationContext(p *UpdateOperationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_updateOperation
}

func (*UpdateOperationContext) IsUpdateOperationContext() {}

func NewUpdateOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UpdateOperationContext {
	var p = new(UpdateOperationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FSLParserRULE_updateOperation

	return p
}

func (s *UpdateOperationContext) GetParser() antlr.Parser { return s.parser }

func (s *UpdateOperationContext) ID() antlr.TerminalNode {
	return s.GetToken(FSLParserID, 0)
}

func (s *UpdateOperationContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *UpdateOperationContext) VariableReference() IVariableReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableReferenceContext)
}

func (s *UpdateOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UpdateOperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UpdateOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.EnterUpdateOperation(s)
	}
}

func (s *UpdateOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.ExitUpdateOperation(s)
	}
}

func (p *FSLParser) UpdateOperation() (localctx IUpdateOperationContext) {
	localctx = NewUpdateOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, FSLParserRULE_updateOperation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		p.Match(FSLParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(156)
		p.Match(FSLParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(157)
		p.Match(FSLParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(158)
		p.Match(FSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(159)
		p.Match(FSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(160)
		p.Match(FSLParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(161)
		p.Match(FSLParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(162)
		p.Match(FSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FSLParserINT, FSLParserFLOAT:
		{
			p.SetState(163)
			p.Value()
		}

	case FSLParserT__17:
		{
			p.SetState(164)
			p.VariableReference()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeleteOperationContext is an interface to support dynamic dispatch.
type IDeleteOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsDeleteOperationContext differentiates from other interfaces.
	IsDeleteOperationContext()
}

type DeleteOperationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeleteOperationContext() *DeleteOperationContext {
	var p = new(DeleteOperationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_deleteOperation
	return p
}

func InitEmptyDeleteOperationContext(p *DeleteOperationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_deleteOperation
}

func (*DeleteOperationContext) IsDeleteOperationContext() {}

func NewDeleteOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeleteOperationContext {
	var p = new(DeleteOperationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FSLParserRULE_deleteOperation

	return p
}

func (s *DeleteOperationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeleteOperationContext) ID() antlr.TerminalNode {
	return s.GetToken(FSLParserID, 0)
}

func (s *DeleteOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeleteOperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeleteOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.EnterDeleteOperation(s)
	}
}

func (s *DeleteOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.ExitDeleteOperation(s)
	}
}

func (p *FSLParser) DeleteOperation() (localctx IDeleteOperationContext) {
	localctx = NewDeleteOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, FSLParserRULE_deleteOperation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		p.Match(FSLParserT__15)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(168)
		p.Match(FSLParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(169)
		p.Match(FSLParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(170)
		p.Match(FSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(171)
		p.Match(FSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintOperationContext is an interface to support dynamic dispatch.
type IPrintOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Value() IValueContext
	VariableReference() IVariableReferenceContext

	// IsPrintOperationContext differentiates from other interfaces.
	IsPrintOperationContext()
}

type PrintOperationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintOperationContext() *PrintOperationContext {
	var p = new(PrintOperationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_printOperation
	return p
}

func InitEmptyPrintOperationContext(p *PrintOperationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_printOperation
}

func (*PrintOperationContext) IsPrintOperationContext() {}

func NewPrintOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintOperationContext {
	var p = new(PrintOperationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FSLParserRULE_printOperation

	return p
}

func (s *PrintOperationContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintOperationContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *PrintOperationContext) VariableReference() IVariableReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableReferenceContext)
}

func (s *PrintOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintOperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.EnterPrintOperation(s)
	}
}

func (s *PrintOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.ExitPrintOperation(s)
	}
}

func (p *FSLParser) PrintOperation() (localctx IPrintOperationContext) {
	localctx = NewPrintOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, FSLParserRULE_printOperation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(173)
		p.Match(FSLParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(174)
		p.Match(FSLParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(175)
		p.Match(FSLParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(176)
		p.Match(FSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FSLParserINT, FSLParserFLOAT:
		{
			p.SetState(177)
			p.Value()
		}

	case FSLParserT__17:
		{
			p.SetState(178)
			p.VariableReference()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllPassedParameter() []IPassedParameterContext
	PassedParameter(i int) IPassedParameterContext

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_functionCall
	return p
}

func InitEmptyFunctionCallContext(p *FunctionCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_functionCall
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FSLParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(FSLParserID)
}

func (s *FunctionCallContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(FSLParserID, i)
}

func (s *FunctionCallContext) AllPassedParameter() []IPassedParameterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPassedParameterContext); ok {
			len++
		}
	}

	tst := make([]IPassedParameterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPassedParameterContext); ok {
			tst[i] = t.(IPassedParameterContext)
			i++
		}
	}

	return tst
}

func (s *FunctionCallContext) PassedParameter(i int) IPassedParameterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPassedParameterContext); ok {
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

	return t.(IPassedParameterContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (p *FSLParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, FSLParserRULE_functionCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(181)
		p.Match(FSLParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(182)
		p.Match(FSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FSLParserT__5 {
		{
			p.SetState(183)
			p.Match(FSLParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(184)
			_la = p.GetTokenStream().LA(1)

			if !(_la == FSLParserT__6 || _la == FSLParserID) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(185)
			p.Match(FSLParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(186)
			p.PassedParameter()
		}

		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPassedParameterContext is an interface to support dynamic dispatch.
type IPassedParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	VariableReference() IVariableReferenceContext

	// IsPassedParameterContext differentiates from other interfaces.
	IsPassedParameterContext()
}

type PassedParameterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPassedParameterContext() *PassedParameterContext {
	var p = new(PassedParameterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_passedParameter
	return p
}

func InitEmptyPassedParameterContext(p *PassedParameterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_passedParameter
}

func (*PassedParameterContext) IsPassedParameterContext() {}

func NewPassedParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PassedParameterContext {
	var p = new(PassedParameterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FSLParserRULE_passedParameter

	return p
}

func (s *PassedParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *PassedParameterContext) ID() antlr.TerminalNode {
	return s.GetToken(FSLParserID, 0)
}

func (s *PassedParameterContext) VariableReference() IVariableReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableReferenceContext)
}

func (s *PassedParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PassedParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PassedParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.EnterPassedParameter(s)
	}
}

func (s *PassedParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.ExitPassedParameter(s)
	}
}

func (p *FSLParser) PassedParameter() (localctx IPassedParameterContext) {
	localctx = NewPassedParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, FSLParserRULE_passedParameter)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FSLParserID:
		{
			p.SetState(192)
			p.Match(FSLParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FSLParserT__17:
		{
			p.SetState(193)
			p.VariableReference()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOperationParameterContext is an interface to support dynamic dispatch.
type IOperationParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VariableReference() IVariableReferenceContext
	Value() IValueContext
	FunctionParameter() IFunctionParameterContext

	// IsOperationParameterContext differentiates from other interfaces.
	IsOperationParameterContext()
}

type OperationParameterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperationParameterContext() *OperationParameterContext {
	var p = new(OperationParameterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_operationParameter
	return p
}

func InitEmptyOperationParameterContext(p *OperationParameterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_operationParameter
}

func (*OperationParameterContext) IsOperationParameterContext() {}

func NewOperationParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationParameterContext {
	var p = new(OperationParameterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FSLParserRULE_operationParameter

	return p
}

func (s *OperationParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *OperationParameterContext) VariableReference() IVariableReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableReferenceContext)
}

func (s *OperationParameterContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *OperationParameterContext) FunctionParameter() IFunctionParameterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionParameterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionParameterContext)
}

func (s *OperationParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperationParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.EnterOperationParameter(s)
	}
}

func (s *OperationParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.ExitOperationParameter(s)
	}
}

func (p *FSLParser) OperationParameter() (localctx IOperationParameterContext) {
	localctx = NewOperationParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, FSLParserRULE_operationParameter)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FSLParserT__17:
		{
			p.SetState(196)
			p.VariableReference()
		}

	case FSLParserINT, FSLParserFLOAT:
		{
			p.SetState(197)
			p.Value()
		}

	case FSLParserT__18:
		{
			p.SetState(198)
			p.FunctionParameter()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentificationParameterContext is an interface to support dynamic dispatch.
type IIdentificationParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	FunctionParameter() IFunctionParameterContext

	// IsIdentificationParameterContext differentiates from other interfaces.
	IsIdentificationParameterContext()
}

type IdentificationParameterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentificationParameterContext() *IdentificationParameterContext {
	var p = new(IdentificationParameterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_identificationParameter
	return p
}

func InitEmptyIdentificationParameterContext(p *IdentificationParameterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_identificationParameter
}

func (*IdentificationParameterContext) IsIdentificationParameterContext() {}

func NewIdentificationParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentificationParameterContext {
	var p = new(IdentificationParameterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FSLParserRULE_identificationParameter

	return p
}

func (s *IdentificationParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentificationParameterContext) ID() antlr.TerminalNode {
	return s.GetToken(FSLParserID, 0)
}

func (s *IdentificationParameterContext) FunctionParameter() IFunctionParameterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionParameterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionParameterContext)
}

func (s *IdentificationParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentificationParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentificationParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.EnterIdentificationParameter(s)
	}
}

func (s *IdentificationParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.ExitIdentificationParameter(s)
	}
}

func (p *FSLParser) IdentificationParameter() (localctx IIdentificationParameterContext) {
	localctx = NewIdentificationParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, FSLParserRULE_identificationParameter)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FSLParserID:
		{
			p.SetState(201)
			p.Match(FSLParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FSLParserT__18:
		{
			p.SetState(202)
			p.FunctionParameter()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionParameterContext is an interface to support dynamic dispatch.
type IFunctionParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsFunctionParameterContext differentiates from other interfaces.
	IsFunctionParameterContext()
}

type FunctionParameterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionParameterContext() *FunctionParameterContext {
	var p = new(FunctionParameterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_functionParameter
	return p
}

func InitEmptyFunctionParameterContext(p *FunctionParameterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_functionParameter
}

func (*FunctionParameterContext) IsFunctionParameterContext() {}

func NewFunctionParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionParameterContext {
	var p = new(FunctionParameterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FSLParserRULE_functionParameter

	return p
}

func (s *FunctionParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionParameterContext) ID() antlr.TerminalNode {
	return s.GetToken(FSLParserID, 0)
}

func (s *FunctionParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.EnterFunctionParameter(s)
	}
}

func (s *FunctionParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.ExitFunctionParameter(s)
	}
}

func (p *FSLParser) FunctionParameter() (localctx IFunctionParameterContext) {
	localctx = NewFunctionParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, FSLParserRULE_functionParameter)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.Match(FSLParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(206)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FSLParserT__6 || _la == FSLParserID) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariableReferenceContext is an interface to support dynamic dispatch.
type IVariableReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsVariableReferenceContext differentiates from other interfaces.
	IsVariableReferenceContext()
}

type VariableReferenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableReferenceContext() *VariableReferenceContext {
	var p = new(VariableReferenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_variableReference
	return p
}

func InitEmptyVariableReferenceContext(p *VariableReferenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_variableReference
}

func (*VariableReferenceContext) IsVariableReferenceContext() {}

func NewVariableReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableReferenceContext {
	var p = new(VariableReferenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FSLParserRULE_variableReference

	return p
}

func (s *VariableReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableReferenceContext) ID() antlr.TerminalNode {
	return s.GetToken(FSLParserID, 0)
}

func (s *VariableReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.EnterVariableReference(s)
	}
}

func (s *VariableReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.ExitVariableReference(s)
	}
}

func (p *FSLParser) VariableReference() (localctx IVariableReferenceContext) {
	localctx = NewVariableReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, FSLParserRULE_variableReference)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(208)
		p.Match(FSLParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(209)
		p.Match(FSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FSLParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FSLParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) INT() antlr.TerminalNode {
	return s.GetToken(FSLParserINT, 0)
}

func (s *ValueContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(FSLParserFLOAT, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FSLListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *FSLParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, FSLParserRULE_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(211)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FSLParserINT || _la == FSLParserFLOAT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
