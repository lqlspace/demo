package ast

import (
	"lqlspace/demo/antlr/parser"
)

type (
	PlaceHolder struct{}

 	Api struct {
		TypeSpec      []*TypeSpec
		ServiceSpec    []*ServiceSpec
	}
)

type ApiVisitor struct {
	*parser.BaseApiParserVisitor
}

func NewApiVisitor() *ApiVisitor {
	return &ApiVisitor{}
}

func (v *ApiVisitor) VisitApi(ctx *parser.ApiContext) interface{} {
	var final Api
	for _, each := range ctx.AllSpec() {
		root := each.Accept(v).(*Api)
		final.TypeSpec = append(final.TypeSpec, root.TypeSpec...)
		final.ServiceSpec = append(final.ServiceSpec, root.ServiceSpec...)
	}

	return &final
}

func (v *ApiVisitor) VisitSpec(ctx *parser.SpecContext) interface{} {
	var root Api
	if ctx.TypeSpec() != nil {
		tp := ctx.TypeSpec().Accept(v)
		root.TypeSpec = append(root.TypeSpec, tp.(*TypeSpec))
	}

	if ctx.ServiceSpec() != nil {
		root.ServiceSpec = []*ServiceSpec{ctx.ServiceSpec().Accept(v).(*ServiceSpec)}
	}

	return &root
}