package handlers

import (
	"helloladies/internal/model"
	"helloladies/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCity(c *gin.Context) {
	var city model.City
	if err := c.Bind(&city); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, errIncorrectBody)
		return
	}

	if err := h.services.CitiesService.CreateCity(city); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response.NewSuccessResponse(c, http.StatusCreated, "city was created")
}

func (h *Handler) ListCities(c *gin.Context) {
	cities, err := h.services.CitiesService.List()
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusAccepted, cities)
}

func (h *Handler) UpdateCity(c *gin.Context) {
	var city model.City

	code := c.Query("code")
	if err := c.Bind(&city); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, errIncorrectBody)
		return
	}

	_, err := h.services.CitiesService.UpdateCity(code, city)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response.NewSuccessResponse(c, http.StatusAccepted, "city was updated")
}

func (h *Handler) DeleteCity(c *gin.Context) {
	var city model.City

	code := c.Query("code")
	if err := c.Bind(&city); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, errIncorrectBody)
		return
	}

	err := h.services.CitiesService.DeleteCity(code)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response.NewSuccessResponse(c, http.StatusAccepted, "city was deleted")
}
