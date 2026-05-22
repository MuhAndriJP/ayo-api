package handler

import (
	"net/http"

	"github.com/MuhAndriJP/ayo-api/internal/dto"
	"github.com/MuhAndriJP/ayo-api/internal/service"
	"github.com/MuhAndriJP/ayo-api/internal/util"
	"github.com/gin-gonic/gin"
)

type MatchHandler struct {
	svc    service.MatchService
	report service.ReportService
}

func NewMatchHandler(svc service.MatchService, report service.ReportService) *MatchHandler {
	return &MatchHandler{svc: svc, report: report}
}

func (h *MatchHandler) List(c *gin.Context) {
	var query dto.ListQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		util.Fail(c, http.StatusBadRequest, util.MsgParamTidakValid, err.Error())
		return
	}

	matches, total, err := h.svc.List(query)
	if err != nil {
		util.Fail(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	util.OK(c, http.StatusOK, util.MsgOK, gin.H{
		"data":  matches,
		"total": total,
		"page":  query.Page,
		"limit": query.Limit,
	})
}

func (h *MatchHandler) Get(c *gin.Context) {
	id, err := util.ParseID(c, util.ParamID)
	if err != nil {
		return
	}

	match, err := h.svc.GetByID(id)
	if err != nil {
		util.Fail(c, http.StatusNotFound, err.Error(), nil)
		return
	}

	util.OK(c, http.StatusOK, util.MsgOK, match)
}

func (h *MatchHandler) Create(c *gin.Context) {
	var req dto.CreateMatchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.Fail(c, http.StatusBadRequest, util.MsgValidasiGagal, err.Error())
		return
	}

	err := h.svc.Create(req)
	if err != nil {
		util.Fail(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	util.OK(c, http.StatusCreated, util.MsgOK, nil)
}

func (h *MatchHandler) Update(c *gin.Context) {
	id, err := util.ParseID(c, util.ParamID)
	if err != nil {
		return
	}

	var req dto.UpdateMatchRequest
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

func (h *MatchHandler) Delete(c *gin.Context) {
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

func (h *MatchHandler) SaveResult(c *gin.Context) {
	id, err := util.ParseID(c, util.ParamID)
	if err != nil {
		return
	}

	var req dto.MatchResultRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.Fail(c, http.StatusBadRequest, util.MsgValidasiGagal, err.Error())
		return
	}

	err = h.svc.SaveResult(id, req)
	if err != nil {
		util.Fail(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	util.OK(c, http.StatusOK, util.MsgOK, nil)
}

func (h *MatchHandler) GetReport(c *gin.Context) {
	id, err := util.ParseID(c, util.ParamID)
	if err != nil {
		return
	}

	report, err := h.report.GetMatchReport(id)
	if err != nil {
		util.Fail(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	util.OK(c, http.StatusOK, util.MsgOK, report)
}
