package patient

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/sebastiantorreslab/desafiofinal-turnos/internal/patient"
	"github.com/sebastiantorreslab/desafiofinal-turnos/pkg/web"
)

type patientHandler struct {
	s patient.IPatientService
}

func NewPatientHandler(s patient.IPatientService) *patientHandler {
	return &patientHandler{
		s: s,
	}
}

func (h *patientHandler) GetAll() gin.HandlerFunc {

	patient, err := h.s.GetAll()

	return func(c *gin.Context) {
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid"))
			return
		}
		web.Success(c, 200, patient)
	}
}
