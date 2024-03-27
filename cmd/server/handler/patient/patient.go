package patient

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sebastiantorreslab/desafiofinal-turnos/internal/domain"
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
	return func(c *gin.Context) {
		patients, err := h.s.GetAll()
		if err != nil {
			web.Failure(c, http.StatusBadRequest, errors.New("unable to retrieve patients"))
			return
		}
		web.Success(c, http.StatusOK, patients)
	}
}

func (h *patientHandler) GetById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Failure(c, http.StatusBadRequest, errors.New("invalid ID format"))
			return
		}

		patient, err := h.s.GetById(id)
		if err != nil {
			web.Failure(c, http.StatusNotFound, errors.New("patient not found"))
			return
		}
		web.Success(c, http.StatusOK, patient)
	}
}

func (h *patientHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.Patient
		if err := c.Bind(&req); err != nil {
			web.Failure(c, http.StatusBadRequest, errors.New("bad request data"))
			return
		}

		d, err := h.s.Create(req)
		if err != nil {
			web.Failure(c, http.StatusInternalServerError, errors.New("unable to create patient"))
			return
		}
		web.Success(c, http.StatusOK, d)
	}
}

func (h *patientHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.Patient
		if err := c.Bind(&req); err != nil {
			web.Failure(c, http.StatusBadRequest, errors.New("bad request data"))
			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Failure(c, http.StatusBadRequest, errors.New("invalid ID format"))
			return
		}

		patient, err := h.s.Update(req, id)
		if err != nil {
			web.Failure(c, http.StatusInternalServerError, errors.New("unable to update patient"))
			return
		}
		web.Success(c, http.StatusOK, gin.H{"data": patient})
	}
}

func (h *patientHandler) UpdateByField() gin.HandlerFunc {
	type Request struct {
		DNI           int    `json:"dni,omitempty"`
		Name          string `json:"name,omitempty"`
		LastName      string `json:"last_name,omitempty"`
		Address       string `json:"address,omitempty"`
		AdmissionDate string `json:"admission_date,omitempty"`
	}
	return func(c *gin.Context) {
		var req Request
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Failure(c, http.StatusBadRequest, errors.New("invalid ID format"))
			return
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			web.Failure(c, http.StatusBadRequest, errors.New("invalid JSON format"))
			return
		}

		update := domain.Patient{
			DNI:           req.DNI,
			Name:          req.Name,
			LastName:      req.LastName,
			Address:       req.Address,
			AdmissionDate: req.AdmissionDate, // Directly assigning the string value
		}

		patient, err := h.s.Update(update, id)
		if err != nil {
			web.Failure(c, http.StatusInternalServerError, errors.New("unable to update patient fields"))
			return
		}
		web.Success(c, http.StatusOK, gin.H{"data": patient})
	}
}

func (h *patientHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Failure(c, http.StatusBadRequest, errors.New("invalid ID format"))
			return
		}

		err = h.s.Delete(id)
		if err != nil {
			web.Failure(c, http.StatusNotFound, errors.New("unable to delete patient"))
			return
		}

		web.Success(c, http.StatusOK, gin.H{"message": "patient deleted"})
	}
}
