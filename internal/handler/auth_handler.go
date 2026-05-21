package handler

import (
	"net/http"

	"github.com/MuhAndriJP/ayo-api/internal/dto"
	"github.com/MuhAndriJP/ayo-api/internal/service"
	"github.com/MuhAndriJP/ayo-api/internal/util"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	svc service.AuthService
}

func NewAuthHandler(svc service.AuthService) *AuthHandler {
	return &AuthHandler{svc: svc}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.Fail(c, http.StatusBadRequest, util.MsgValidasiGagal, err.Error())
		return
	}

	if err := h.svc.Register(req); err != nil {
		util.Fail(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	util.OK(c, http.StatusCreated, util.MsgOK, nil)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.Fail(c, http.StatusBadRequest, util.MsgValidasiGagal, err.Error())
		return
	}

	token, err := h.svc.Login(req)
	if err != nil {
		util.Fail(c, http.StatusUnauthorized, err.Error(), nil)
		return
	}

	util.OK(c, http.StatusOK, util.MsgOK, dto.AuthResponse{Token: token})
}
