package handlers

import (
	"api-management/internal/services"
	"api-management/internal/utils"
	"net/http"
)

type APIKeyHandler struct {
	service *services.APIKeyService
}

func NewAPIKeyHandler(service *services.APIKeyService) *APIKeyHandler {
	return &APIKeyHandler{
		service: service,
	}
}

type APIKeyRequest struct {
	Name string `json:"name"`
}

func (h *APIKeyHandler) CreateAPIKeyHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := r.Context()
	claims, ok := utils.FromContext(ctx)
	if !ok {
		utils.ErrorResponse(w, http.StatusUnauthorized, "Unauthorized")
	}

	userId := claims.UserId

	req := new(APIKeyRequest)

	utils.CheckInvalidRequest[APIKeyRequest](w, r, req)

	if req.Name == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "Name is required")
		return
	}

	apikey, err := h.service.CreateApiKey(ctx, userId, req.Name)

	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Could not create Api key")
		return
	}

	utils.JSONResponse(w, http.StatusCreated, map[string]any{
		"message": "Save this API key now. You won't see it again.",
		"apiKey":  apikey.ApiKey,
		"rawKey":  apikey.RawKey,
	})
}
