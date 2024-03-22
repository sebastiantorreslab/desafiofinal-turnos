package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sebastiantorreslab/desafiofinal-turnos/internal/domain/dentist"
)

type DentistHandler struct {
	DentistService dentist.IDentistService
}

func (h *DentistHandler) GetAll(ctx *gin.Context) {
	dentists, err := h.DentistService.GetAll()
	if err != nil {

	}
	ctx.JSON(http.StatusOK, &dentists)
}
