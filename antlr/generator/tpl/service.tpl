package service


{{if .TypeSpecs}}
type (
{{- range .TypeSpecs}}
    {{.Name}} struct {
{{range .Fields -}}
    {{.Name}} {{.DataType}} {{.Tag}}
{{- end}}
    }
{{- end}}
)
{{- end}}



func {{.Handler}}Service(req {{.Request}}) (resp {{.Response}}, err error) {
	//TODO: coding by yourself
	return 
}