// Code generated from ApiParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ApiParser

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ApiParserParser struct {
	*antlr.BaseParser
}

var apiparserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func apiparserParserInit() {
	staticData := &apiparserParserStaticData
	staticData.literalNames = []string{
		"", "'type'", "'struct'", "'{'", "'}'", "'('", "')'", "'returns'", "'-'",
		"'/'", "'/:'", "'@handler'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "ATHANDLER", "WS", "RAW_STRING",
		"ID",
	}
	staticData.ruleNames = []string{
		"api", "spec", "typeSpec", "field", "serviceSpec", "serviceRoute", "atHandler",
		"route", "request", "response", "serviceName", "path",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 14, 120, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 1, 0, 5, 0, 26, 8, 0, 10, 0, 12, 0, 29, 9, 0, 1, 1, 1,
		1, 3, 1, 33, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 40, 8, 2, 10, 2,
		12, 2, 43, 9, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 3, 3, 50, 8, 3, 1, 4, 1,
		4, 1, 4, 1, 4, 5, 4, 56, 8, 4, 10, 4, 12, 4, 59, 9, 4, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 3, 7, 72, 8, 7, 1, 7, 3,
		7, 75, 8, 7, 1, 8, 1, 8, 3, 8, 79, 8, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9,
		3, 9, 86, 8, 9, 1, 9, 1, 9, 1, 10, 1, 10, 3, 10, 92, 8, 10, 4, 10, 94,
		8, 10, 11, 10, 12, 10, 95, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 102, 8, 11,
		10, 11, 12, 11, 105, 9, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 111, 8,
		11, 4, 11, 113, 8, 11, 11, 11, 12, 11, 114, 1, 11, 3, 11, 118, 8, 11, 1,
		11, 0, 0, 12, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 0, 0, 123, 0,
		27, 1, 0, 0, 0, 2, 32, 1, 0, 0, 0, 4, 34, 1, 0, 0, 0, 6, 46, 1, 0, 0, 0,
		8, 51, 1, 0, 0, 0, 10, 62, 1, 0, 0, 0, 12, 65, 1, 0, 0, 0, 14, 68, 1, 0,
		0, 0, 16, 76, 1, 0, 0, 0, 18, 82, 1, 0, 0, 0, 20, 93, 1, 0, 0, 0, 22, 117,
		1, 0, 0, 0, 24, 26, 3, 2, 1, 0, 25, 24, 1, 0, 0, 0, 26, 29, 1, 0, 0, 0,
		27, 25, 1, 0, 0, 0, 27, 28, 1, 0, 0, 0, 28, 1, 1, 0, 0, 0, 29, 27, 1, 0,
		0, 0, 30, 33, 3, 4, 2, 0, 31, 33, 3, 8, 4, 0, 32, 30, 1, 0, 0, 0, 32, 31,
		1, 0, 0, 0, 33, 3, 1, 0, 0, 0, 34, 35, 5, 1, 0, 0, 35, 36, 5, 14, 0, 0,
		36, 37, 5, 2, 0, 0, 37, 41, 5, 3, 0, 0, 38, 40, 3, 6, 3, 0, 39, 38, 1,
		0, 0, 0, 40, 43, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42,
		44, 1, 0, 0, 0, 43, 41, 1, 0, 0, 0, 44, 45, 5, 4, 0, 0, 45, 5, 1, 0, 0,
		0, 46, 47, 5, 14, 0, 0, 47, 49, 5, 14, 0, 0, 48, 50, 5, 13, 0, 0, 49, 48,
		1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 7, 1, 0, 0, 0, 51, 52, 5, 14, 0, 0,
		52, 53, 3, 20, 10, 0, 53, 57, 5, 3, 0, 0, 54, 56, 3, 10, 5, 0, 55, 54,
		1, 0, 0, 0, 56, 59, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0,
		58, 60, 1, 0, 0, 0, 59, 57, 1, 0, 0, 0, 60, 61, 5, 4, 0, 0, 61, 9, 1, 0,
		0, 0, 62, 63, 3, 12, 6, 0, 63, 64, 3, 14, 7, 0, 64, 11, 1, 0, 0, 0, 65,
		66, 5, 11, 0, 0, 66, 67, 5, 14, 0, 0, 67, 13, 1, 0, 0, 0, 68, 69, 5, 14,
		0, 0, 69, 71, 3, 22, 11, 0, 70, 72, 3, 16, 8, 0, 71, 70, 1, 0, 0, 0, 71,
		72, 1, 0, 0, 0, 72, 74, 1, 0, 0, 0, 73, 75, 3, 18, 9, 0, 74, 73, 1, 0,
		0, 0, 74, 75, 1, 0, 0, 0, 75, 15, 1, 0, 0, 0, 76, 78, 5, 5, 0, 0, 77, 79,
		5, 14, 0, 0, 78, 77, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0,
		80, 81, 5, 6, 0, 0, 81, 17, 1, 0, 0, 0, 82, 83, 5, 7, 0, 0, 83, 85, 5,
		5, 0, 0, 84, 86, 5, 14, 0, 0, 85, 84, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0, 86,
		87, 1, 0, 0, 0, 87, 88, 5, 6, 0, 0, 88, 19, 1, 0, 0, 0, 89, 91, 5, 14,
		0, 0, 90, 92, 5, 8, 0, 0, 91, 90, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 94,
		1, 0, 0, 0, 93, 89, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0,
		95, 96, 1, 0, 0, 0, 96, 21, 1, 0, 0, 0, 97, 98, 5, 9, 0, 0, 98, 103, 5,
		14, 0, 0, 99, 100, 5, 8, 0, 0, 100, 102, 5, 14, 0, 0, 101, 99, 1, 0, 0,
		0, 102, 105, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104,
		113, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 106, 107, 5, 10, 0, 0, 107, 110,
		5, 14, 0, 0, 108, 109, 5, 8, 0, 0, 109, 111, 5, 14, 0, 0, 110, 108, 1,
		0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 113, 1, 0, 0, 0, 112, 97, 1, 0, 0,
		0, 112, 106, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 114,
		115, 1, 0, 0, 0, 115, 118, 1, 0, 0, 0, 116, 118, 5, 9, 0, 0, 117, 112,
		1, 0, 0, 0, 117, 116, 1, 0, 0, 0, 118, 23, 1, 0, 0, 0, 16, 27, 32, 41,
		49, 57, 71, 74, 78, 85, 91, 95, 103, 110, 112, 114, 117,
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

// ApiParserParserInit initializes any static state used to implement ApiParserParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewApiParserParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ApiParserParserInit() {
	staticData := &apiparserParserStaticData
	staticData.once.Do(apiparserParserInit)
}

// NewApiParserParser produces a new parser instance for the optional input antlr.TokenStream.
func NewApiParserParser(input antlr.TokenStream) *ApiParserParser {
	ApiParserParserInit()
	this := new(ApiParserParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &apiparserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "ApiParser.g4"

	return this
}

// ApiParserParser tokens.
const (
	ApiParserParserEOF        = antlr.TokenEOF
	ApiParserParserT__0       = 1
	ApiParserParserT__1       = 2
	ApiParserParserT__2       = 3
	ApiParserParserT__3       = 4
	ApiParserParserT__4       = 5
	ApiParserParserT__5       = 6
	ApiParserParserT__6       = 7
	ApiParserParserT__7       = 8
	ApiParserParserT__8       = 9
	ApiParserParserT__9       = 10
	ApiParserParserATHANDLER  = 11
	ApiParserParserWS         = 12
	ApiParserParserRAW_STRING = 13
	ApiParserParserID         = 14
)

// ApiParserParser rules.
const (
	ApiParserParserRULE_api          = 0
	ApiParserParserRULE_spec         = 1
	ApiParserParserRULE_typeSpec     = 2
	ApiParserParserRULE_field        = 3
	ApiParserParserRULE_serviceSpec  = 4
	ApiParserParserRULE_serviceRoute = 5
	ApiParserParserRULE_atHandler    = 6
	ApiParserParserRULE_route        = 7
	ApiParserParserRULE_request      = 8
	ApiParserParserRULE_response     = 9
	ApiParserParserRULE_serviceName  = 10
	ApiParserParserRULE_path         = 11
)

// IApiContext is an interface to support dynamic dispatch.
type IApiContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsApiContext differentiates from other interfaces.
	IsApiContext()
}

type ApiContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyApiContext() *ApiContext {
	var p = new(ApiContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_api
	return p
}

func (*ApiContext) IsApiContext() {}

func NewApiContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ApiContext {
	var p = new(ApiContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_api

	return p
}

func (s *ApiContext) GetParser() antlr.Parser { return s.parser }

func (s *ApiContext) AllSpec() []ISpecContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISpecContext); ok {
			len++
		}
	}

	tst := make([]ISpecContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISpecContext); ok {
			tst[i] = t.(ISpecContext)
			i++
		}
	}

	return tst
}

func (s *ApiContext) Spec(i int) ISpecContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpecContext); ok {
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

	return t.(ISpecContext)
}

func (s *ApiContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ApiContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ApiContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitApi(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) Api() (localctx IApiContext) {
	this := p
	_ = this

	localctx = NewApiContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ApiParserParserRULE_api)
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
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ApiParserParserT__0 || _la == ApiParserParserID {
		{
			p.SetState(24)
			p.Spec()
		}

		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISpecContext is an interface to support dynamic dispatch.
type ISpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSpecContext differentiates from other interfaces.
	IsSpecContext()
}

type SpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpecContext() *SpecContext {
	var p = new(SpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_spec
	return p
}

func (*SpecContext) IsSpecContext() {}

func NewSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpecContext {
	var p = new(SpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_spec

	return p
}

func (s *SpecContext) GetParser() antlr.Parser { return s.parser }

func (s *SpecContext) TypeSpec() ITypeSpecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeSpecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *SpecContext) ServiceSpec() IServiceSpecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServiceSpecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IServiceSpecContext)
}

func (s *SpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) Spec() (localctx ISpecContext) {
	this := p
	_ = this

	localctx = NewSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ApiParserParserRULE_spec)

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

	p.SetState(32)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ApiParserParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(30)
			p.TypeSpec()
		}

	case ApiParserParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(31)
			p.ServiceSpec()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITypeSpecContext is an interface to support dynamic dispatch.
type ITypeSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTypeToken returns the typeToken token.
	GetTypeToken() antlr.Token

	// GetStructToken returns the structToken token.
	GetStructToken() antlr.Token

	// GetLbrace returns the lbrace token.
	GetLbrace() antlr.Token

	// GetRbrace returns the rbrace token.
	GetRbrace() antlr.Token

	// SetTypeToken sets the typeToken token.
	SetTypeToken(antlr.Token)

	// SetStructToken sets the structToken token.
	SetStructToken(antlr.Token)

	// SetLbrace sets the lbrace token.
	SetLbrace(antlr.Token)

	// SetRbrace sets the rbrace token.
	SetRbrace(antlr.Token)

	// IsTypeSpecContext differentiates from other interfaces.
	IsTypeSpecContext()
}

type TypeSpecContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	typeToken   antlr.Token
	structToken antlr.Token
	lbrace      antlr.Token
	rbrace      antlr.Token
}

func NewEmptyTypeSpecContext() *TypeSpecContext {
	var p = new(TypeSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_typeSpec
	return p
}

func (*TypeSpecContext) IsTypeSpecContext() {}

func NewTypeSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSpecContext {
	var p = new(TypeSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_typeSpec

	return p
}

func (s *TypeSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSpecContext) GetTypeToken() antlr.Token { return s.typeToken }

func (s *TypeSpecContext) GetStructToken() antlr.Token { return s.structToken }

func (s *TypeSpecContext) GetLbrace() antlr.Token { return s.lbrace }

func (s *TypeSpecContext) GetRbrace() antlr.Token { return s.rbrace }

func (s *TypeSpecContext) SetTypeToken(v antlr.Token) { s.typeToken = v }

func (s *TypeSpecContext) SetStructToken(v antlr.Token) { s.structToken = v }

func (s *TypeSpecContext) SetLbrace(v antlr.Token) { s.lbrace = v }

func (s *TypeSpecContext) SetRbrace(v antlr.Token) { s.rbrace = v }

func (s *TypeSpecContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *TypeSpecContext) AllField() []IFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldContext); ok {
			len++
		}
	}

	tst := make([]IFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldContext); ok {
			tst[i] = t.(IFieldContext)
			i++
		}
	}

	return tst
}

func (s *TypeSpecContext) Field(i int) IFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
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

	return t.(IFieldContext)
}

func (s *TypeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitTypeSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) TypeSpec() (localctx ITypeSpecContext) {
	this := p
	_ = this

	localctx = NewTypeSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ApiParserParserRULE_typeSpec)
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
	{
		p.SetState(34)

		var _m = p.Match(ApiParserParserT__0)

		localctx.(*TypeSpecContext).typeToken = _m
	}
	{
		p.SetState(35)
		p.Match(ApiParserParserID)
	}
	{
		p.SetState(36)

		var _m = p.Match(ApiParserParserT__1)

		localctx.(*TypeSpecContext).structToken = _m
	}
	{
		p.SetState(37)

		var _m = p.Match(ApiParserParserT__2)

		localctx.(*TypeSpecContext).lbrace = _m
	}
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ApiParserParserID {
		{
			p.SetState(38)
			p.Field()
		}

		p.SetState(43)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(44)

		var _m = p.Match(ApiParserParserT__3)

		localctx.(*TypeSpecContext).rbrace = _m
	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTag returns the tag token.
	GetTag() antlr.Token

	// SetTag sets the tag token.
	SetTag(antlr.Token)

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	tag    antlr.Token
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) GetTag() antlr.Token { return s.tag }

func (s *FieldContext) SetTag(v antlr.Token) { s.tag = v }

func (s *FieldContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ApiParserParserID)
}

func (s *FieldContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, i)
}

func (s *FieldContext) RAW_STRING() antlr.TerminalNode {
	return s.GetToken(ApiParserParserRAW_STRING, 0)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) Field() (localctx IFieldContext) {
	this := p
	_ = this

	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ApiParserParserRULE_field)
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
	{
		p.SetState(46)
		p.Match(ApiParserParserID)
	}
	{
		p.SetState(47)
		p.Match(ApiParserParserID)
	}
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserRAW_STRING {
		{
			p.SetState(48)

			var _m = p.Match(ApiParserParserRAW_STRING)

			localctx.(*FieldContext).tag = _m
		}

	}

	return localctx
}

// IServiceSpecContext is an interface to support dynamic dispatch.
type IServiceSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLbrace returns the lbrace token.
	GetLbrace() antlr.Token

	// GetRbrace returns the rbrace token.
	GetRbrace() antlr.Token

	// SetLbrace sets the lbrace token.
	SetLbrace(antlr.Token)

	// SetRbrace sets the rbrace token.
	SetRbrace(antlr.Token)

	// IsServiceSpecContext differentiates from other interfaces.
	IsServiceSpecContext()
}

type ServiceSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	lbrace antlr.Token
	rbrace antlr.Token
}

func NewEmptyServiceSpecContext() *ServiceSpecContext {
	var p = new(ServiceSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_serviceSpec
	return p
}

func (*ServiceSpecContext) IsServiceSpecContext() {}

func NewServiceSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceSpecContext {
	var p = new(ServiceSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_serviceSpec

	return p
}

func (s *ServiceSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceSpecContext) GetLbrace() antlr.Token { return s.lbrace }

func (s *ServiceSpecContext) GetRbrace() antlr.Token { return s.rbrace }

func (s *ServiceSpecContext) SetLbrace(v antlr.Token) { s.lbrace = v }

func (s *ServiceSpecContext) SetRbrace(v antlr.Token) { s.rbrace = v }

func (s *ServiceSpecContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *ServiceSpecContext) ServiceName() IServiceNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServiceNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IServiceNameContext)
}

func (s *ServiceSpecContext) AllServiceRoute() []IServiceRouteContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IServiceRouteContext); ok {
			len++
		}
	}

	tst := make([]IServiceRouteContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IServiceRouteContext); ok {
			tst[i] = t.(IServiceRouteContext)
			i++
		}
	}

	return tst
}

func (s *ServiceSpecContext) ServiceRoute(i int) IServiceRouteContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServiceRouteContext); ok {
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

	return t.(IServiceRouteContext)
}

func (s *ServiceSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitServiceSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) ServiceSpec() (localctx IServiceSpecContext) {
	this := p
	_ = this

	localctx = NewServiceSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ApiParserParserRULE_serviceSpec)
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
	{
		p.SetState(51)
		p.Match(ApiParserParserID)
	}
	{
		p.SetState(52)
		p.ServiceName()
	}
	{
		p.SetState(53)

		var _m = p.Match(ApiParserParserT__2)

		localctx.(*ServiceSpecContext).lbrace = _m
	}
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ApiParserParserATHANDLER {
		{
			p.SetState(54)
			p.ServiceRoute()
		}

		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(60)

		var _m = p.Match(ApiParserParserT__3)

		localctx.(*ServiceSpecContext).rbrace = _m
	}

	return localctx
}

// IServiceRouteContext is an interface to support dynamic dispatch.
type IServiceRouteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsServiceRouteContext differentiates from other interfaces.
	IsServiceRouteContext()
}

type ServiceRouteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceRouteContext() *ServiceRouteContext {
	var p = new(ServiceRouteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_serviceRoute
	return p
}

func (*ServiceRouteContext) IsServiceRouteContext() {}

func NewServiceRouteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceRouteContext {
	var p = new(ServiceRouteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_serviceRoute

	return p
}

func (s *ServiceRouteContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceRouteContext) AtHandler() IAtHandlerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtHandlerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtHandlerContext)
}

func (s *ServiceRouteContext) Route() IRouteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRouteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRouteContext)
}

func (s *ServiceRouteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceRouteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceRouteContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitServiceRoute(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) ServiceRoute() (localctx IServiceRouteContext) {
	this := p
	_ = this

	localctx = NewServiceRouteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ApiParserParserRULE_serviceRoute)

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
		p.SetState(62)
		p.AtHandler()
	}
	{
		p.SetState(63)
		p.Route()
	}

	return localctx
}

// IAtHandlerContext is an interface to support dynamic dispatch.
type IAtHandlerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtHandlerContext differentiates from other interfaces.
	IsAtHandlerContext()
}

type AtHandlerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtHandlerContext() *AtHandlerContext {
	var p = new(AtHandlerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_atHandler
	return p
}

func (*AtHandlerContext) IsAtHandlerContext() {}

func NewAtHandlerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtHandlerContext {
	var p = new(AtHandlerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_atHandler

	return p
}

func (s *AtHandlerContext) GetParser() antlr.Parser { return s.parser }

func (s *AtHandlerContext) ATHANDLER() antlr.TerminalNode {
	return s.GetToken(ApiParserParserATHANDLER, 0)
}

func (s *AtHandlerContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *AtHandlerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtHandlerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtHandlerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitAtHandler(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) AtHandler() (localctx IAtHandlerContext) {
	this := p
	_ = this

	localctx = NewAtHandlerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ApiParserParserRULE_atHandler)

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
		p.SetState(65)
		p.Match(ApiParserParserATHANDLER)
	}
	{
		p.SetState(66)
		p.Match(ApiParserParserID)
	}

	return localctx
}

// IRouteContext is an interface to support dynamic dispatch.
type IRouteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRouteContext differentiates from other interfaces.
	IsRouteContext()
}

type RouteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRouteContext() *RouteContext {
	var p = new(RouteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_route
	return p
}

func (*RouteContext) IsRouteContext() {}

func NewRouteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RouteContext {
	var p = new(RouteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_route

	return p
}

func (s *RouteContext) GetParser() antlr.Parser { return s.parser }

func (s *RouteContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *RouteContext) Path() IPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *RouteContext) Request() IRequestContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRequestContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRequestContext)
}

func (s *RouteContext) Response() IResponseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResponseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IResponseContext)
}

func (s *RouteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RouteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RouteContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitRoute(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) Route() (localctx IRouteContext) {
	this := p
	_ = this

	localctx = NewRouteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ApiParserParserRULE_route)
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
	{
		p.SetState(68)
		p.Match(ApiParserParserID)
	}
	{
		p.SetState(69)
		p.Path()
	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserT__4 {
		{
			p.SetState(70)
			p.Request()
		}

	}
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserT__6 {
		{
			p.SetState(73)
			p.Response()
		}

	}

	return localctx
}

// IRequestContext is an interface to support dynamic dispatch.
type IRequestContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLp returns the lp token.
	GetLp() antlr.Token

	// GetRp returns the rp token.
	GetRp() antlr.Token

	// SetLp sets the lp token.
	SetLp(antlr.Token)

	// SetRp sets the rp token.
	SetRp(antlr.Token)

	// IsRequestContext differentiates from other interfaces.
	IsRequestContext()
}

type RequestContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	lp     antlr.Token
	rp     antlr.Token
}

func NewEmptyRequestContext() *RequestContext {
	var p = new(RequestContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_request
	return p
}

func (*RequestContext) IsRequestContext() {}

func NewRequestContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RequestContext {
	var p = new(RequestContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_request

	return p
}

func (s *RequestContext) GetParser() antlr.Parser { return s.parser }

func (s *RequestContext) GetLp() antlr.Token { return s.lp }

func (s *RequestContext) GetRp() antlr.Token { return s.rp }

func (s *RequestContext) SetLp(v antlr.Token) { s.lp = v }

func (s *RequestContext) SetRp(v antlr.Token) { s.rp = v }

func (s *RequestContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *RequestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RequestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RequestContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitRequest(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) Request() (localctx IRequestContext) {
	this := p
	_ = this

	localctx = NewRequestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ApiParserParserRULE_request)
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
	{
		p.SetState(76)

		var _m = p.Match(ApiParserParserT__4)

		localctx.(*RequestContext).lp = _m
	}
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserID {
		{
			p.SetState(77)
			p.Match(ApiParserParserID)
		}

	}
	{
		p.SetState(80)

		var _m = p.Match(ApiParserParserT__5)

		localctx.(*RequestContext).rp = _m
	}

	return localctx
}

// IResponseContext is an interface to support dynamic dispatch.
type IResponseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetReturnToken returns the returnToken token.
	GetReturnToken() antlr.Token

	// GetLp returns the lp token.
	GetLp() antlr.Token

	// GetRp returns the rp token.
	GetRp() antlr.Token

	// SetReturnToken sets the returnToken token.
	SetReturnToken(antlr.Token)

	// SetLp sets the lp token.
	SetLp(antlr.Token)

	// SetRp sets the rp token.
	SetRp(antlr.Token)

	// IsResponseContext differentiates from other interfaces.
	IsResponseContext()
}

type ResponseContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	returnToken antlr.Token
	lp          antlr.Token
	rp          antlr.Token
}

func NewEmptyResponseContext() *ResponseContext {
	var p = new(ResponseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_response
	return p
}

func (*ResponseContext) IsResponseContext() {}

func NewResponseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResponseContext {
	var p = new(ResponseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_response

	return p
}

func (s *ResponseContext) GetParser() antlr.Parser { return s.parser }

func (s *ResponseContext) GetReturnToken() antlr.Token { return s.returnToken }

func (s *ResponseContext) GetLp() antlr.Token { return s.lp }

func (s *ResponseContext) GetRp() antlr.Token { return s.rp }

func (s *ResponseContext) SetReturnToken(v antlr.Token) { s.returnToken = v }

func (s *ResponseContext) SetLp(v antlr.Token) { s.lp = v }

func (s *ResponseContext) SetRp(v antlr.Token) { s.rp = v }

func (s *ResponseContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *ResponseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResponseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResponseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitResponse(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) Response() (localctx IResponseContext) {
	this := p
	_ = this

	localctx = NewResponseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ApiParserParserRULE_response)
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
	{
		p.SetState(82)

		var _m = p.Match(ApiParserParserT__6)

		localctx.(*ResponseContext).returnToken = _m
	}
	{
		p.SetState(83)

		var _m = p.Match(ApiParserParserT__4)

		localctx.(*ResponseContext).lp = _m
	}
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserID {
		{
			p.SetState(84)
			p.Match(ApiParserParserID)
		}

	}
	{
		p.SetState(87)

		var _m = p.Match(ApiParserParserT__5)

		localctx.(*ResponseContext).rp = _m
	}

	return localctx
}

// IServiceNameContext is an interface to support dynamic dispatch.
type IServiceNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsServiceNameContext differentiates from other interfaces.
	IsServiceNameContext()
}

type ServiceNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceNameContext() *ServiceNameContext {
	var p = new(ServiceNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_serviceName
	return p
}

func (*ServiceNameContext) IsServiceNameContext() {}

func NewServiceNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceNameContext {
	var p = new(ServiceNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_serviceName

	return p
}

func (s *ServiceNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceNameContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ApiParserParserID)
}

func (s *ServiceNameContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, i)
}

func (s *ServiceNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitServiceName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) ServiceName() (localctx IServiceNameContext) {
	this := p
	_ = this

	localctx = NewServiceNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ApiParserParserRULE_serviceName)
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
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ApiParserParserID {
		{
			p.SetState(89)
			p.Match(ApiParserParserID)
		}
		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ApiParserParserT__7 {
			{
				p.SetState(90)
				p.Match(ApiParserParserT__7)
			}

		}

		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPathContext is an interface to support dynamic dispatch.
type IPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPathContext differentiates from other interfaces.
	IsPathContext()
}

type PathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPathContext() *PathContext {
	var p = new(PathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_path
	return p
}

func (*PathContext) IsPathContext() {}

func NewPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PathContext {
	var p = new(PathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_path

	return p
}

func (s *PathContext) GetParser() antlr.Parser { return s.parser }

func (s *PathContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ApiParserParserID)
}

func (s *PathContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, i)
}

func (s *PathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitPath(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) Path() (localctx IPathContext) {
	this := p
	_ = this

	localctx = NewPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ApiParserParserRULE_path)
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

	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(112)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ApiParserParserT__8 || _la == ApiParserParserT__9 {
			p.SetState(112)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case ApiParserParserT__8:
				{
					p.SetState(97)
					p.Match(ApiParserParserT__8)
				}

				{
					p.SetState(98)
					p.Match(ApiParserParserID)
				}
				p.SetState(103)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				for _la == ApiParserParserT__7 {
					{
						p.SetState(99)
						p.Match(ApiParserParserT__7)
					}
					{
						p.SetState(100)
						p.Match(ApiParserParserID)
					}

					p.SetState(105)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}

			case ApiParserParserT__9:
				{
					p.SetState(106)
					p.Match(ApiParserParserT__9)
				}

				{
					p.SetState(107)
					p.Match(ApiParserParserID)
				}
				p.SetState(110)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == ApiParserParserT__7 {
					{
						p.SetState(108)
						p.Match(ApiParserParserT__7)
					}
					{
						p.SetState(109)
						p.Match(ApiParserParserID)
					}

				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(114)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(116)
			p.Match(ApiParserParserT__8)
		}

	}

	return localctx
}
