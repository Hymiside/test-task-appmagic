package handler

import (
	"net/http"

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
	handlers := NewHandlers(s)

	h.handler.Get("/api/gas-per-month", handlers.getInfoGasPerMonth)
	h.handler.Post("/api/price-per-day", handlers.getInfoPricePerDay)
	h.handler.Get("/api/hourly-price", handlers.getInfoHourlyPrice)
	h.handler.Get("/api/sum-all-period", handlers.getInfoSumAllPeriod)

	return h.handler
}

func (s *Handlers) getInfoGasPerMonth(w http.ResponseWriter, r *http.Request) {}

func (s *Handlers) getInfoPricePerDay(w http.ResponseWriter, r *http.Request) {}

func (s *Handlers) getInfoHourlyPrice(w http.ResponseWriter, r *http.Request) {}

func (s *Handlers) getInfoSumAllPeriod(w http.ResponseWriter, r *http.Request) {
	res := s.service.GetInfoSumAllPeriod()
	ResponseOk(w, res)
}
