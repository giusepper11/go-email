package endpoints

import (
	"net/http"
)

func (h *Handler) CampaignGetAll(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	campaigns, err := h.CampaignService.Repository.Get()
	return campaigns, 200, err
}
