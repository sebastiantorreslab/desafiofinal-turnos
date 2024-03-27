package dentist

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sebastiantorreslab/desafiofinal-turnos/internal/dentist"
	"github.com/sebastiantorreslab/desafiofinal-turnos/internal/domain"
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

	return func(c *gin.Context) {
		dentists, err := h.s.GetAll()
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid"))
			return
		}
		web.Success(c, 200, dentists)
	}
}

func (h *dentistHandler) GetById() gin.HandlerFunc {

	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid id"))
			return
		}

		dentist, err := h.s.GetById(id)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid"))
			return
		}
		web.Success(c, 200, dentist)
	}

}

func (h *dentistHandler) Create() gin.HandlerFunc {

	return func(c *gin.Context) {

		var req domain.Dentist
		err := c.Bind(&req)
		if err != nil {
			web.Failure(c, 400, errors.New("Bad request"))
			return

		}

		d, err := h.s.Create(req)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid"))
			return

		}
		web.Success(c, 200, d)

	}

}

func (h *dentistHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {

		var req domain.Dentist
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

		dentist, err := h.s.Update(req, id)
		if err != nil {
			web.Failure(c, 500, errors.New("Internal server error"))
			return
		}

		web.Success(c, http.StatusOK, gin.H{
			"data": dentist,
		})

	}

}
func (h *dentistHandler) UpdateByField() gin.HandlerFunc {
	type Request struct {
		License  string `json:"license,omitempty"`
		Name     string `json:"name,omitempty"`
		LastName string `json:"last_name,omitempty"`
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

		if err := c.ShouldBindJSON(&req); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		update := domain.Dentist{
			License:  req.License,
			Name:     req.Name,
			LastName: req.LastName,
		}

		dentist, err := h.s.Update(update, id)
		if err != nil {
			web.Failure(c, 404, errors.New("Invalid reques getById"))
			return
		}
		web.Success(c, http.StatusOK, gin.H{
			"data": dentist,
		})

	}
}

func (h *dentistHandler) Delete() gin.HandlerFunc {
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
