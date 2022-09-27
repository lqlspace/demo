package main

import (
	"fmt"
	"log"

	"lqlspace/demo/antlr/ast"
	"lqlspace/demo/antlr/generator"
)

func main() {
	parser := ast.NewParser()
	apiobj, err := parser.Parse("dsl/dsl.api")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(apiobj)
	m := make(map[string]interface{})

	m["Handler"] = apiobj.ServiceSpec[0].ServiceRoute[0].AtHandler.Name
	m["Path"] = apiobj.ServiceSpec[0].ServiceRoute[0].Route.Path
	m["Request"] = apiobj.ServiceSpec[0].ServiceRoute[0].Route.Req
	m["Response"] = apiobj.ServiceSpec[0].ServiceRoute[0].Route.Resp
	m["TypeSpecs"] = apiobj.TypeSpec

	generator.GenMapper(m, "generator/server", "generator/tpl/main.tpl", "generator/server/main.go")
	generator.GenMapper(m, "generator/server/handler", "generator/tpl/router.tpl", "generator/server/handler/router.go")
	generator.GenMapper(m, "generator/server/handler", "generator/tpl/handler.tpl", "generator/server/handler/handler.go")
	generator.GenMapper(m, "generator/server/service", "generator/tpl/service.tpl", "generator/server/service/service.go")
}