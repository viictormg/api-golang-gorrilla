package handler

import (
	"api/platzi/model"
	"api/platzi/server"
	"encoding/json"
	"net/http"
)

func HomeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(model.HomeResponse{
			Message: "hola care bola",
			Status:  true,
		})
	}
}
