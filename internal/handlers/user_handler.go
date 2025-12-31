package handlers

import (
	"api-management/internal/services"
	"api-management/internal/utils"
	"net/http"
)

type UserHandler struct {
	service *services.UserService
	jwt     *utils.JWTManager
}

func NewUserHandler(service *services.UserService, jwt *utils.JWTManager) *UserHandler {
	return &UserHandler{
		service: service,
		jwt:     jwt,
	}
}

type CreateUserRequest struct {
	Email        string `json:"email"`
	PasswordHash string `json:"password"`
}

type UserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req CreateUserRequest

	utils.CheckInvalidRequest[CreateUserRequest](w, r, &req)

	if req.Email == "" || req.PasswordHash == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "email and password are required")
		return
	}

	user, err := h.service.CreateUser(r.Context(), req.PasswordHash, req.Email)

	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	token, err := h.jwt.GenerateToken(user.ID, user.Email)

	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Could not generate token")
		return
	}

	utils.JSONResponse(w, http.StatusCreated, map[string]string{
		"message": "User created successfully",
		"id":      user.ID,
		"token":   token,
		"email":   user.Email,
	})
}

func (h *UserHandler) GetUserByEmailHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "email is required in query parameter")
		return
	}
	user, err := h.service.GetUserByEmail(r.Context(), email)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.JSONResponse(w, http.StatusOK, UserResponse{
		ID:    user.ID,
		Email: user.Email,
	})
}
