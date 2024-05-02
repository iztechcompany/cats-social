package cat

import (
	"cats-social/internal/service/cat"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CatHandler interface {
	Create(c *gin.Context)
}

type catHandler struct {
	catService cat.CatService
	validate   *validator.Validate
}

func NewCatHandler(catService cat.CatService, validate *validator.Validate) CatHandler {
	return &catHandler{
		catService: catService,
		validate:   validate,
	}
}
