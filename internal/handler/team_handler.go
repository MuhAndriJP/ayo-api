package handler

import (
	"net/http"

	"github.com/MuhAndriJP/ayo-api/internal/dto"
	"github.com/MuhAndriJP/ayo-api/internal/service"
	"github.com/MuhAndriJP/ayo-api/internal/util"
	"github.com/gin-gonic/gin"
)

type TeamHandler struct {
	svc service.TeamService
}

func NewTeamHandler(svc service.TeamService) *TeamHandler {
	return &TeamHandler{svc: svc}
}

func (h *TeamHandler) List(c *gin.Context) {
	var query dto.ListQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		util.Fail(c, http.StatusBadRequest, util.MsgParamTidakValid, err.Error())
		return
	}

	teams, total, err := h.svc.List(query)
	if err != nil {
		util.Fail(c, http.StatusInternalServerError, util.MsgGagalAmbilData, err.Error())
		return
	}

	util.OK(c, http.StatusOK, util.MsgOK, gin.H{
		"data":  teams,
		"total": total,
		"page":  query.Page,
		"limit": query.Limit,
	})
}

func (h *TeamHandler) Get(c *gin.Context) {
	id, err := util.ParseID(c, util.ParamID)
	if err != nil {
		return
	}

	team, err := h.svc.GetByID(id)
	if err != nil {
		util.Fail(c, http.StatusNotFound, err.Error(), nil)
		return
	}

	util.OK(c, http.StatusOK, util.MsgOK, team)
}

func (h *TeamHandler) Create(c *gin.Context) {
	var req dto.CreateTeamRequest
	if err := c.ShouldBind(&req); err != nil {
		util.Fail(c, http.StatusBadRequest, util.MsgValidasiGagal, err.Error())
		return
	}

	logo, _ := c.FormFile(util.FormLogo)
	err := h.svc.Create(req, logo)
	if err != nil {
		util.Fail(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	util.OK(c, http.StatusCreated, util.MsgOK, nil)
}

func (h *TeamHandler) Update(c *gin.Context) {
	id, err := util.ParseID(c, util.ParamID)
	if err != nil {
		return
	}

	var req dto.UpdateTeamRequest
	if err := c.ShouldBind(&req); err != nil {
		util.Fail(c, http.StatusBadRequest, util.MsgValidasiGagal, err.Error())
		return
	}

	logo, _ := c.FormFile(util.FormLogo)
	if err = h.svc.Update(id, req, logo); err != nil {
		util.Fail(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	util.OK(c, http.StatusOK, util.MsgOK, nil)
}

func (h *TeamHandler) Delete(c *gin.Context) {
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
