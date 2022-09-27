package handler

import (
	"github.com/Hymiside/test-task-appmagic/pkg/service"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	handler *chi.Mux
}

type Handlers struct {
	service *service.Service
}

func NewHandlers(s service.Service) *Handlers {
	return &Handlers{service: &s}
}

// InitHandler инициализирует хэндлеры
func (h *Handler) InitHandler(s service.Service) *chi.Mux {
	h.handler = chi.NewRouter()
	// handlers := NewHandlers(s)

	h.handler.Get("/api/gas-per-month", nil)
	h.handler.Patch("/api/update", nil)
	h.handler.Delete("/api/remove", nil)
	h.handler.Get("/api/list", nil)

	return h.handler
}
