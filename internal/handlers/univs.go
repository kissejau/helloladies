package handlers

import (
	"helloladies/internal/model"
	"helloladies/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

func (h *Handler) ListUnivs(c *gin.Context) {
	univs, err := h.services.UnivsService.List()
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, univs)
}

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

func (h *Handler) DeleteUniv(c *gin.Context) {
	univCode := c.Query("univ_code")

	if err := h.services.UnivsService.DeleteUniv(univCode); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	response.NewSuccessResponse(c, http.StatusAccepted, "university was deleted")
}

func (h *Handler) GetUnivsByCity(c *gin.Context) {
	cityCode := c.Query("city_code")

	univs, err := h.services.GetUnivsByCity(cityCode)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, univs)
}
