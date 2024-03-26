package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/sebastiantorreslab/desafiofinal-turnos/internal/dentist"
	"github.com/sebastiantorreslab/desafiofinal-turnos/pkg/web"
)

type dentistHandler struct {
	s dentist.IDentistService
}

func NewDentistHandler(s dentist.IDentistService) *dentistHandler {
	return &dentistHandler{
		s: s,
	}
}

func (h *dentistHandler) GetAll() gin.HandlerFunc {

	dentist,err := h.s.GetAll()
	
	return func(c *gin.Context) {
		if err != nil {
			web.Failure(c,400,errors.New("Invalid"))
			return
		}
		web.Success(c, 200, dentist)
	}
}
