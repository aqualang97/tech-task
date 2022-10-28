package router

import (
	"github.com/go-chi/chi"
	"techTask/internal/controller"
)

func Router(c *controller.Controller, r *chi.Mux) {
	r.Get("/transaction/{id}", c.Data.GetByTransactionID)
	r.Get("/terminal/{id}", c.Data.GetByTerminalID)
	r.Get("/status={s}", c.Data.GetByStatus)
	r.Get("/payment-type={t}", c.Data.GetByPaymentType)
	r.Get("/from={start}&to={end}", c.Data.GetByTime)
	r.Get("/narrative={text}", c.Data.GetByNarrative)
}
