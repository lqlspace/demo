package ast

import (
	"lqlspace/demo/antlr/parser"
)

type (
	ServiceSpec struct {
		Name         string
		ServiceRoute []*ServiceRoute
	}

	ServiceRoute struct {
		AtHandler *AtHandler
		Route     *Route
	}

	AtHandler struct {
		Name           string 
	}

	Route struct {
		Method      string 
		Path        string 
		Req         string 
		Resp        string 
	}	
)


func (v *ApiVisitor) VisitServiceSpec(ctx *parser.ServiceSpecContext) interface{} {
	var serviceSpec ServiceSpec
	if ctx.ServiceName() != nil {
		serviceSpec.Name = ctx.ServiceName().Accept(v).(string)
	}

	for _, each := range ctx.AllServiceRoute() {
		serviceSpec.ServiceRoute = append(serviceSpec.ServiceRoute, each.Accept(v).(*ServiceRoute))
	}

	return &serviceSpec
}

func (v *ApiVisitor) VisitServiceRoute(ctx *parser.ServiceRouteContext) interface{} {
	var serviceRoute ServiceRoute
	if ctx.AtHandler() != nil {
		serviceRoute.AtHandler = ctx.AtHandler().Accept(v).(*AtHandler)
	}

	serviceRoute.Route = ctx.Route().Accept(v).(*Route)
	return &serviceRoute
}


func (v *ApiVisitor) VisitAtHandler(ctx *parser.AtHandlerContext) interface{} {
	var atHandler AtHandler
	atHandler.Name = ctx.ID().GetText()
	return &atHandler
}

func (v *ApiVisitor) VisitRoute(ctx *parser.RouteContext) interface{} {
	var route Route
	route.Method = ctx.ID().GetText()
	route.Path = ctx.Path().GetText()

	if ctx.Request() != nil {
		route.Req = ctx.Request().Accept(v).(string)
	}

	if ctx.Response() != nil {
		route.Resp = ctx.Response().Accept(v).(string)
	}

	return &route
}


func (v *ApiVisitor) VisitRequest(ctx *parser.RequestContext) interface{} {
	return ctx.ID().GetText()
}

func (v *ApiVisitor) VisitResponse(ctx *parser.ResponseContext) interface{} {
	return ctx.ID().GetText()
}



func (v *ApiVisitor) VisitServiceName(ctx *parser.ServiceNameContext) interface{} {
	return ctx.ID(0).GetText()
}