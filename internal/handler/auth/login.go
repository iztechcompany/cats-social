package auth

import (
	"cats-social/helper"
	"cats-social/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *authHandler) Login(c *gin.Context) {
	var loginRequest model.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
		return
	}

	err := h.validate.Struct(loginRequest)
	if err != nil {
		helper.GetResponse(err.Error(), "request doesn't pass validation", 400, true).Send(c)
		return
	}

	response, err := h.authService.Login(c, &loginRequest)
	if err != nil {
		helper.GetResponse("", err.Error(), 400, true).Send(c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User logged successfully",
		"data":    response,
	})
}
