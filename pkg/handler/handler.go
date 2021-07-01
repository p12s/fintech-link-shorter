package handler

import (
	"github.com/p12s/fintech-link-shorter/pkg/service"
	"github.com/rs/zerolog/log"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (*Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err := w.Write([]byte("Hello from simple server."))
	if err != nil {
		log.Err(err)
	}
	w.WriteHeader(http.StatusOK)
}
