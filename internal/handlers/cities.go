package handlers

import (
	"helloladies/internal/model"
	"helloladies/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

//	@Summary	CreateCity
//	@Security	Token
//	@Tags		cities
//	@Accept		json
//	@Produce	json
//	@Param		input	body		model.City	true	"city data"
//	@Success	200		{object}	response.Success
//	@Failure	400		{object}	response.Error
//	@Router		/logged/cities/create [post]
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

//	@Summary	ListCities
//	@Security	Token
//	@Tags		cities
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	[]model.City
//	@Failure	400	{object}	response.Error
//	@Router		/logged/cities/all [get]
func (h *Handler) ListCities(c *gin.Context) {
	cities, err := h.services.CitiesService.List()
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusAccepted, cities)
}

//	@Summary	UpdateCity
//	@Security	Token
//	@Tags		cities
//	@Accept		json
//	@Produce	json
//	@Param		input	body		model.City	true	"city data"
//	@Success	200		{object}	response.Success
//	@Failure	400		{object}	response.Error
//	@Router		/logged/cities/update [put]
func (h *Handler) UpdateCity(c *gin.Context) {
	var city model.City
	if err := c.Bind(&city); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, errIncorrectBody)
		return
	}

	_, err := h.services.CitiesService.UpdateCity(city)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response.NewSuccessResponse(c, http.StatusAccepted, "city was updated")
}

//	@Summary	DeleteCity
//	@Security	Token
//	@Tags		cities
//	@Accept		json
//	@Produce	json
//	@Param		city_code	query		string	true	"city's code"
//	@Success	200			{object}	response.Success
//	@Failure	400			{object}	response.Error
//	@Router		/logged/cities/delete [delete]
func (h *Handler) DeleteCity(c *gin.Context) {
	code := c.Query("city_code")

	err := h.services.CitiesService.DeleteCity(code)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response.NewSuccessResponse(c, http.StatusAccepted, "city was deleted")
}
