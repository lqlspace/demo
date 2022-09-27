// Code generated from ApiParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ApiParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseApiParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseApiParserVisitor) VisitApi(ctx *ApiContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseApiParserVisitor) VisitSpec(ctx *SpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseApiParserVisitor) VisitTypeSpec(ctx *TypeSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseApiParserVisitor) VisitField(ctx *FieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseApiParserVisitor) VisitServiceSpec(ctx *ServiceSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseApiParserVisitor) VisitServiceRoute(ctx *ServiceRouteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseApiParserVisitor) VisitAtHandler(ctx *AtHandlerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseApiParserVisitor) VisitRoute(ctx *RouteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseApiParserVisitor) VisitRequest(ctx *RequestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseApiParserVisitor) VisitResponse(ctx *ResponseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseApiParserVisitor) VisitServiceName(ctx *ServiceNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseApiParserVisitor) VisitPath(ctx *PathContext) interface{} {
	return v.VisitChildren(ctx)
}
