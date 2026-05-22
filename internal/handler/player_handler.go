package handler

import (
	"net/http"

	"github.com/MuhAndriJP/ayo-api/internal/dto"
	"github.com/MuhAndriJP/ayo-api/internal/service"
	"github.com/MuhAndriJP/ayo-api/internal/util"
	"github.com/gin-gonic/gin"
)

type PlayerHandler struct {
	svc service.PlayerService
}

func NewPlayerHandler(svc service.PlayerService) *PlayerHandler {
	return &PlayerHandler{svc: svc}
}

func (h *PlayerHandler) ListByTeam(c *gin.Context) {
	teamID, err := util.ParseID(c, util.ParamID)
	if err != nil {
		return
	}

	var query dto.ListQuery
	_ = c.ShouldBindQuery(&query)

	players, err := h.svc.ListByTeam(teamID, query)
	if err != nil {
		util.Fail(c, http.StatusNotFound, err.Error(), nil)
		return
	}

	util.OK(c, http.StatusOK, util.MsgOK, players)
}

func (h *PlayerHandler) Get(c *gin.Context) {
	id, err := util.ParseID(c, util.ParamID)
	if err != nil {
		return
	}

	player, err := h.svc.GetByID(id)
	if err != nil {
		util.Fail(c, http.StatusNotFound, err.Error(), nil)
		return
	}

	util.OK(c, http.StatusOK, util.MsgOK, player)
}

func (h *PlayerHandler) Create(c *gin.Context) {
	teamID, err := util.ParseID(c, util.ParamID)
	if err != nil {
		return
	}

	var req dto.CreatePlayerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.Fail(c, http.StatusBadRequest, util.MsgValidasiGagal, err.Error())
		return
	}

	err = h.svc.Create(teamID, req)
	if err != nil {
		util.Fail(c, http.StatusConflict, err.Error(), nil)
		return
	}

	util.OK(c, http.StatusCreated, util.MsgOK, nil)
}

func (h *PlayerHandler) Update(c *gin.Context) {
	id, err := util.ParseID(c, util.ParamID)
	if err != nil {
		return
	}

	var req dto.UpdatePlayerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.Fail(c, http.StatusBadRequest, util.MsgValidasiGagal, err.Error())
		return
	}

	err = h.svc.Update(id, req)
	if err != nil {
		util.Fail(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	util.OK(c, http.StatusOK, util.MsgOK, nil)
}

func (h *PlayerHandler) Delete(c *gin.Context) {
	id, err := util.ParseID(c, util.ParamID)
	if err != nil {
		return
	}

	if err := h.svc.Delete(id); err != nil {
		util.Fail(c, http.StatusNotFound, err.Error(), nil)
		return
	}

	util.OK(c, http.StatusOK, util.MsgOK, nil)
}
