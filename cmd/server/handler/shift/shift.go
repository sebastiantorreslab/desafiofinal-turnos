package shift

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sebastiantorreslab/desafiofinal-turnos/internal/domain"
	"github.com/sebastiantorreslab/desafiofinal-turnos/internal/shift"
	"github.com/sebastiantorreslab/desafiofinal-turnos/pkg/web"
)

type ShiftHandler struct {
	s shift.IShiftService
}

func NewShiftHandler(s shift.IShiftService) *ShiftHandler {
	return &ShiftHandler{
		s: s,
	}
}

// @Summary Obtiene todos los turnos
// @Description Recupera una lista de todos los turnos registrados
// @Tags shifts
// @Accept json
// @Produce json
// @Success 200 {array} domain.Shift "Lista de turnos"
// @Router /shifts [get]
func (h *ShiftHandler) GetAll() gin.HandlerFunc {

	return func(c *gin.Context) {
		shifts, err := h.s.GetAll()
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid"))
			return
		}
		web.Success(c, 200, shifts)
	}
}

// @Summary Obtiene los detalles de un turno
// @Description Recupera los detalles de un turno por su ID
// @Tags shifts
// @Accept json
// @Produce json
// @Param id path int true "ID del Turno"
// @Success 200 {object} domain.Shift "Detalles del turno"
// @Router /shifts/{id} [get]
func (h *ShiftHandler) GetById() gin.HandlerFunc {

	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid id"))
			return
		}

		shift, err := h.s.GetById(id)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid"))
			return
		}
		web.Success(c, 200, shift)
	}

}

// @Summary Crea un nuevo turno
// @Description Crea un nuevo turno con los datos proporcionados
// @Tags shifts
// @Accept json
// @Produce json
// @Param shift body domain.Shift true "Datos del Turno"
// @Success 200 {object} domain.Shift "Turno creado"
// @Security ApiKeyAuth
// @Router /shifts [post]
func (h *ShiftHandler) Create() gin.HandlerFunc {

	return func(c *gin.Context) {

		var req domain.Shift
		err := c.Bind(&req)
		if err != nil {
			web.Failure(c, 400, errors.New("Bad request"))
			return

		}

		s, err := h.s.Create(req)
		if err != nil {
			log.Fatal(err)
			return

		}
		web.Success(c, 200, s)

	}

}

// @Summary Actualiza los datos de un turno
// @Description Actualiza los datos de un turno con la información proporcionada
// @Tags shifts
// @Accept json
// @Produce json
// @Param id path int true "ID del Turno"
// @Param shift body domain.Shift true "Datos actualizados del Turno"
// @Success 200 {object} domain.Shift "Turno actualizado exitosamente"
// @Security ApiKeyAuth
// @Router /shifts/{id} [put]
func (h *ShiftHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {

		var req domain.Shift
		err := c.Bind(&req)
		if err != nil {
			web.Failure(c, 400, errors.New("Bad request"))
			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Failure(c, 404, errors.New("Invalid param"))
			return
		}

		s, err := h.s.Update(req, id)
		if err != nil {
			web.Failure(c, 500, errors.New("Internal server error"))
			return
		}

		web.Success(c, http.StatusOK, gin.H{
			"data": s,
		})

	}

}
func (h *ShiftHandler) UpdateByField() gin.HandlerFunc {
	type Request struct {
		ID         int    `json:"ID"`
		ShiftDate  string `json:"shift_date" binding:"required,omitempty"`
		ShiftHour  string `json:"shift_hour" binding:"required,omitempty"`
		IdPatient  int    `json:"id_patient" binding:"required,omitempty"`
		IdDentist  int    `json:"id_dentist" binding:"required,omitempty"`
		PatientDNI int    `json:"patient_DNI" binding:"required,omitempty"`
	}
	return func(c *gin.Context) {

		var req Request
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid param"))
			return
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		update := domain.Shift{
			ShiftDate:  req.ShiftDate,
			ShiftHour:  req.ShiftHour,
			IdPatient:  req.IdPatient,
			IdDentist:  req.IdDentist,
			PatientDNI: req.PatientDNI,
		}

		s, err := h.s.Update(update, id)
		if err != nil {
			web.Failure(c, 404, errors.New("Invalid request getById"))
			return
		}
		web.Success(c, http.StatusOK, gin.H{
			"data": s,
		})

	}
}

// @Summary Elimina un turno
// @Description Elimina un turno del sistema por su ID
// @Tags shifts
// @Accept json
// @Produce json
// @Param id path int true "ID del Turno"
// @Success 200 {object} object "Mensaje de confirmación de eliminación"
// @Security ApiKeyAuth
// @Router /shifts/{id} [delete]
func (h *ShiftHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Failure(c, 404, errors.New("Invalid param"))
			return
		}

		err = h.s.Delete(id)
		if err != nil {
			web.Failure(c, 404, errors.New("Invalid request delete"))
			return
		}

		web.Success(c, http.StatusOK, gin.H{
			"message": "shift deleted",
		})

	}
}

func (h *ShiftHandler) GetByDNI() gin.HandlerFunc {

	return func(c *gin.Context) {
		dni, err := strconv.Atoi(c.Query("dni"))
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid dni"))
			return
		}

		shift, err := h.s.GetByDNI(dni)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid request"))
			return
		}
		web.Success(c, 200, shift)
	}

}
