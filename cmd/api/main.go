package main

import (
	"go-email/internal/domain/campaign"
	"go-email/internal/endpoints"
	"go-email/internal/infrastructure/database"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Compress(5))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	campaignService := campaign.Service{
		Repository: &database.CampaignRepository{},
	}

	handler := endpoints.Handler{
		CampaignService: campaignService,
	}

	r.Post("/campaigns", endpoints.HandlerError(handler.CampaignPost))

	r.Get("/campaigns", endpoints.HandlerError(handler.CampaignGetAll))

	http.ListenAndServe(":3000", r)
}
