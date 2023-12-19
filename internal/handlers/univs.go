package handlers

import (
	"helloladies/internal/model"
	"helloladies/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

//	@Summary	CreateUniversity
//	@Security	Token
//	@Tags		universities
//	@Accept		json
//	@Produce	json
//	@Param		city_code	query		string		true	"city's code"
//	@Param		input		body		model.Univ	true	"university data"
//	@Success	200			{object}	response.Success
//	@Failure	400			{object}	response.Error
//	@Router		/logged/univs/create [post]
func (h *Handler) CreateUniv(c *gin.Context) {
	var univ model.Univ

	cityCode := c.Query("city_code")

	if err := c.Bind(&univ); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, errIncorrectBody)
		return
	}

	if err := h.services.UnivsService.CreateUniv(cityCode, univ); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response.NewSuccessResponse(c, http.StatusCreated, "university was created")
}

//	@Summary	ListUniversities
//	@Security	Token
//	@Tags		universities
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	[]model.Univ
//	@Failure	400	{object}	response.Error
//	@Router		/logged/univs/all [get]
func (h *Handler) ListUnivs(c *gin.Context) {
	univs, err := h.services.UnivsService.List()
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, univs)
}

//	@Summary	UpdateUniversity
//	@Security	Token
//	@Tags		universities
//	@Accept		json
//	@Produce	json
//	@Param		city_code	query		string		true	"city's code"
//	@Param		input		body		model.Univ	true	"university data"
//	@Success	200			{object}	response.Success
//	@Failure	400			{object}	response.Error
//	@Router		/logged/univs/update [put]
func (h *Handler) UpdateUniv(c *gin.Context) {
	cityCode := c.Query("city_code")

	var univ model.Univ
	if err := c.Bind(&univ); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, errIncorrectBody)
		return
	}

	if _, err := h.services.UnivsService.UpdateUniv(cityCode, univ); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response.NewSuccessResponse(c, http.StatusAccepted, "university was updated")
}

//	@Summary	DeleteUniversity
//	@Security	Token
//	@Tags		universities
//	@Accept		json
//	@Produce	json
//	@Param		univ_code	query		string	true	"university's code"
//	@Success	200			{object}	response.Success
//	@Failure	400			{object}	response.Error
//	@Router		/logged/univs/delete [delete]
func (h *Handler) DeleteUniv(c *gin.Context) {
	univCode := c.Query("univ_code")

	if err := h.services.UnivsService.DeleteUniv(univCode); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	response.NewSuccessResponse(c, http.StatusAccepted, "university was deleted")
}

//	@Summary	GetUniversitiesByCity
//	@Security	Token
//	@Tags		universities
//	@Accept		json
//	@Produce	json
//	@Param		city_code	query		string	true	"city's code"
//	@Success	200			{object}	[]model.Univ
//	@Failure	400			{object}	response.Error
//	@Router		/logged/univs/list [get]
func (h *Handler) GetUnivsByCity(c *gin.Context) {
	cityCode := c.Query("city_code")

	univs, err := h.services.GetUnivsByCity(cityCode)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, univs)
}
