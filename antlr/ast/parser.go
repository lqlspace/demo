package ast

import (
	"io/ioutil"
	"log"
	"strings"

	"lqlspace/demo/antlr/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type (
	Parser struct {
		antlr.DefaultErrorListener
	}
)

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(filename string) (*Api, error) {
	filename = strings.ReplaceAll(filename, `"`, "")
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return p.parse(filename, string(data))
}

func (p *Parser) parse(filename, content string) (*Api, error) {
	inputStream := antlr.NewInputStream(content)
	lexer := parser.NewApiParserLexer(inputStream)
	lexer.RemoveErrorListeners()
	tokens := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	apiParser := parser.NewApiParserParser(tokens)
	apiParser.RemoveErrorListeners()
	apiParser.AddErrorListener(p)

	visitor := NewApiVisitor()
	return apiParser.Api().Accept(visitor).(*Api), nil 
}