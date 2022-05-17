package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getAllUsers(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	users, err := h.services.Api.GetAllUsers()

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *Handler) getUserById(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}

	user, err := h.services.Api.GetUserById(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) getUsersByServiceType(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	serviceType := c.Param("service")

	users, err := h.services.Api.GetUsersByServiceType(serviceType)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *Handler) getUsersByLocation(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	location := c.Param("location")

	users, err := h.services.Api.GetUsersByLocation(location)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}
