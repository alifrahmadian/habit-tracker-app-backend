package handlers

import (
	"net/http"
	"time"

	"github.com/alifrahmadian/habit-tracker-app-backend/internal/constants"
	"github.com/alifrahmadian/habit-tracker-app-backend/internal/handlers/dtos"
	"github.com/alifrahmadian/habit-tracker-app-backend/internal/handlers/responses"
	"github.com/alifrahmadian/habit-tracker-app-backend/internal/models"
	"github.com/alifrahmadian/habit-tracker-app-backend/internal/services"
	"github.com/alifrahmadian/habit-tracker-app-backend/pkg/errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	AuthService services.AuthService
	SecretKey   string
	TTL         int
}

func NewAuthHandler(authService *services.AuthService, secretKey string, ttl int) *AuthHandler {
	return &AuthHandler{
		AuthService: *authService,
		SecretKey:   secretKey,
		TTL:         ttl,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	req := dtos.RegisterRequest{}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "FirstName":
				responses.FailedResponse(c, http.StatusBadRequest, errors.ErrUserFirstNameRequired.Error())
				return
			case "LastName":
				responses.FailedResponse(c, http.StatusBadRequest, errors.ErrUserLastNameRequired.Error())
				return
			case "Username":
				responses.FailedResponse(c, http.StatusBadRequest, errors.ErrUserUsernameRequired.Error())
				return
			case "Email":
				responses.FailedResponse(c, http.StatusBadRequest, errors.ErrUserEmailRequired.Error())
				return
			case "Password":
				responses.FailedResponse(c, http.StatusBadRequest, errors.ErrPasswordRequired.Error())
				return
			}
		}
	}

	user := &models.User{
		RoleId:    constants.USER_ROLE_USER,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Username:  req.Username,
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	userModel, err := h.AuthService.Register(user)
	if err != nil {
		if err == errors.ErrUsernameAlreadyExist || err == errors.ErrEmailAlreadyExist {
			responses.FailedResponse(c, http.StatusConflict, err.Error())
			return
		}

		responses.FailedResponse(c, http.StatusInternalServerError, err.Error())

	}

	resp := &dtos.RegisterResponse{
		Id:        userModel.Id.String(),
		RoleId:    userModel.RoleId,
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
		Username:  userModel.Username,
		Email:     userModel.Email,
	}

	responses.SuccessResponse(c, "register success!", resp)
}
