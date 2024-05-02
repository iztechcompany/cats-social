package cat

import (
	"cats-social/helper"
	"cats-social/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *catHandler) Create(c *gin.Context) {
	var request model.CatRequest

	// check request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
		return
	}

	err := h.validate.Struct(request)
	if err != nil {
		helper.GetResponse(err.Error(), "request doesn't pass validation", 400, true).Send(c)
		return
	}

	response, err := h.catService.Create(c, &request)
	if err != nil {
		helper.GetResponse("", err.Error(), 400, true).Send(c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"data":    response,
	})
}
