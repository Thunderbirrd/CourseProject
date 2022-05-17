package handler

import (
	"github.com/Thunderbirrd/CourseProject/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) updateUser(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input models.UpdateUserInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.User.UpdateUser(userId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
