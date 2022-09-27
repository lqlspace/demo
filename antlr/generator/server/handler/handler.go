package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"lqlspace/demo/antlr/generator/server/service"
)

func GetStudentHandler(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req service.StuRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp, err := service.GetStudentService(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Write([]byte(resp.Data))
	}
}
