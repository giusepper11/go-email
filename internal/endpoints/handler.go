package endpoints

import "go-email/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.Service
}

