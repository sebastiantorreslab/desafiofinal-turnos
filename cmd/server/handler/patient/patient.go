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

// @Summary Obtiene todos los pacientes
// @Description Recupera una lista de todos los pacientes registrados en el sistema
// @Tags patients
// @Accept json
// @Produce json
// @Success 200 {array} domain.Patient "Lista de pacientes"
// @Failure 400 {object} object "Solicitud incorrecta"
// @Router /patients [get]
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

// @Summary Obtiene los detalles de un paciente específico
// @Description Recupera los detalles de un paciente por su ID
// @Tags patients
// @Accept json
// @Produce json
// @Param id path int true "ID del Paciente"
// @Success 200 {object} domain.Patient "Detalles del paciente"
// @Failure 404 {object} object "Paciente no encontrado"
// @Router /patients/{id} [get]
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

// @Summary Crea un nuevo paciente
// @Description Crea un nuevo paciente con los datos proporcionados
// @Tags patients
// @Accept json
// @Produce json
// @Param patient body domain.Patient true "Datos del Paciente"
// @Success 200 {object} domain.Patient "Paciente creado exitosamente"
// @Failure 400 {object} object "Solicitud incorrecta"
// @Failure 500 {object} object "Error interno del servidor"
// @Router /patients [post]
// @Security ApiKeyAuth
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

// @Summary Actualiza los datos de un paciente
// @Description Actualiza los datos de un paciente con la información proporcionada
// @Tags patients
// @Accept json
// @Produce json
// @Param id path int true "ID del Paciente"
// @Param patient body domain.Patient true "Datos actualizados del Paciente"
// @Success 200 {object} object "Paciente actualizado exitosamente"
// @Failure 400 {object} object "Solicitud incorrecta"
// @Failure 404 {object} object "Paciente no encontrado"
// @Router /patients/{id} [put]
// @Security ApiKeyAuth
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

// @Summary Actualiza campos específicos de un paciente
// @Description Actualiza uno o más campos de un paciente sin necesidad de enviar toda la información
// @Tags patients
// @Accept json
// @Produce json
// @Param id path int true "ID del Paciente"
// @Success 200 {object} object "Paciente actualizado exitosamente"
// @Failure 400 {object} object "Solicitud incorrecta"
// @Failure 404 {object} object "Paciente no encontrado"
// @Router /patients/{id}/fields [patch]
// @Security ApiKeyAuth
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

// @Summary Elimina un paciente
// @Description Elimina un paciente del sistema
// @Tags patients
// @Produce json
// @Param id path int true "ID del Paciente"
// @Success 200 {string} string "Mensaje de confirmación"
// @Failure 404 {string} string "Paciente no encontrado"
// @Router /patients/{id} [delete]
// @Security ApiKeyAuth
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
