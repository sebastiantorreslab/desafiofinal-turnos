package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sebastiantorreslab/desafiofinal-turnos/internal/domain"
	"github.com/sebastiantorreslab/desafiofinal-turnos/internal/shift"
	"github.com/sebastiantorreslab/desafiofinal-turnos/pkg/web"
)

type shiftHandler struct {
	s shift.IShiftService
}

func NewShiftHandler(s shift.IShiftService) *shiftHandler {
	return &shiftHandler{
		s: s,
	}
}

func (h *shiftHandler) GetAll() gin.HandlerFunc {

	shifts, err := h.s.GetAll()

	return func(c *gin.Context) {
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid"))
			return
		}
		web.Success(c, 200, shifts)
	}
}

func (h *shiftHandler) GetById() gin.HandlerFunc {

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

func (h *shiftHandler) Create() gin.HandlerFunc {

	return func(c *gin.Context) {

		var req domain.Shift
		err := c.Bind(&req)
		if err != nil {
			web.Failure(c, 400, errors.New("Bad request"))
			return

		}

		s, err := h.s.Create(req)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid"))
			return

		}
		web.Success(c, 200, s)

	}

}

func (h *shiftHandler) Update() gin.HandlerFunc {
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

		err = h.s.Update(req, id)
		if err != nil {
			web.Failure(c, 500, errors.New("Internal server error"))
			return
		}

		web.Success(c, http.StatusOK, gin.H{
			"data": "shift updated",
		})

	}

}
func (h *shiftHandler) UpdateByField() gin.HandlerFunc {
	type Request struct {
		ID        int    `json:"ID"`
		IdPatient int    `json:"patient" binding:"required,omitempty"`
		IdDentist int    `json:"dentist" binding:"required,omitempty"`
		ShiftHour string `json:"shift_hour" binding:"required,omitempty"`
		ShiftDate string `json:"shift_date" binding:"required,omitempty"`
	}
	return func(c *gin.Context) {

		/*token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}*/

		var req Request
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid param"))
			return
		}
		_, err = h.s.GetById(id)
		if err != nil {
			web.Failure(c, 404, errors.New("product not found"))
			return
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		update := domain.Shift{
			IdPatient: req.IdPatient,
			IdDentist: req.IdDentist,
			ShiftHour: req.ShiftHour,
			ShiftDate: req.ShiftDate,
		}

		err = h.s.Update(update, id)
		if err != nil {
			web.Failure(c, 404, errors.New("Invalid reques getById"))
			return
		}
		web.Success(c, http.StatusOK, gin.H{
			"data": "fields updated",
		})

	}
}

func (h *shiftHandler) Delete() gin.HandlerFunc {
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
			"message": "dentist deleted",
		})

	}
}
