// Code generated from ApiParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ApiParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by ApiParserParser.
type ApiParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ApiParserParser#api.
	VisitApi(ctx *ApiContext) interface{}

	// Visit a parse tree produced by ApiParserParser#spec.
	VisitSpec(ctx *SpecContext) interface{}

	// Visit a parse tree produced by ApiParserParser#typeSpec.
	VisitTypeSpec(ctx *TypeSpecContext) interface{}

	// Visit a parse tree produced by ApiParserParser#field.
	VisitField(ctx *FieldContext) interface{}

	// Visit a parse tree produced by ApiParserParser#serviceSpec.
	VisitServiceSpec(ctx *ServiceSpecContext) interface{}

	// Visit a parse tree produced by ApiParserParser#serviceRoute.
	VisitServiceRoute(ctx *ServiceRouteContext) interface{}

	// Visit a parse tree produced by ApiParserParser#atHandler.
	VisitAtHandler(ctx *AtHandlerContext) interface{}

	// Visit a parse tree produced by ApiParserParser#route.
	VisitRoute(ctx *RouteContext) interface{}

	// Visit a parse tree produced by ApiParserParser#request.
	VisitRequest(ctx *RequestContext) interface{}

	// Visit a parse tree produced by ApiParserParser#response.
	VisitResponse(ctx *ResponseContext) interface{}

	// Visit a parse tree produced by ApiParserParser#serviceName.
	VisitServiceName(ctx *ServiceNameContext) interface{}

	// Visit a parse tree produced by ApiParserParser#path.
	VisitPath(ctx *PathContext) interface{}
}
