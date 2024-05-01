package auth

import (
	"cats-social/helper"
	"cats-social/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *authHandler) Register(c *gin.Context) {
	var registerRequest model.RegisterRequest
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
		return
	}

	err := h.validate.Struct(registerRequest)
	if err != nil {
		helper.GetResponse(err.Error(), "request doesn't pass validation", 400, true).Send(c)
		return
	}

	response, err := h.authService.Register(c, &registerRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"data":    response,
	})
}
