package handler

import (
	"context"
	"net/http"
)


func InitRouter() *http.ServeMux {
	ctx := context.Background()
	mux := http.NewServeMux()

	mux.HandleFunc("{{.Path}}", {{.Handler}}Handler(ctx))
	return mux 
}