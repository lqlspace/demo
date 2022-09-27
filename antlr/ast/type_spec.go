package ast

import "lqlspace/demo/antlr/parser"

type (
	TypeSpec struct {
		Name    string 
		Fields  []*TypeField
	}

	TypeField struct {
		Name        string
		DataType    string
		Tag         string
	}
)


func (v *ApiVisitor) VisitTypeSpec(ctx *parser.TypeSpecContext) interface{} {
	var ts TypeSpec
	ts.Name = ctx.ID().GetText()

	fields := ctx.AllField()
	for _, each := range fields {
		f := each.Accept(v)
		if f == nil {
			continue
		}
		ts.Fields = append(ts.Fields, f.(*TypeField))
	}
	return &ts
}


func (v *ApiVisitor) VisitField(ctx *parser.FieldContext) interface{} {
	var field TypeField
	field.Name = ctx.ID(0).GetText()
	field.DataType = ctx.ID(1).GetText()
	if ctx.GetTag() != nil {
		field.Tag = ctx.GetTag().GetText()
	}

	return &field
}