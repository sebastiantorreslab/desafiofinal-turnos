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

// @Summary Obtiene todos los dentistas
// @Description Recupera una lista de todos los dentistas registrados
// @Tags dentists
// @Accept json
// @Produce json
// @Success 200 {array} domain.Dentist "Lista de dentistas"
// @Router /dentists [get]
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

// @Summary Obtiene un dentista por ID
// @Description Recupera los detalles de un dentista por su ID
// @Tags dentists
// @Accept json
// @Produce json
// @Param id path int true "ID del Dentista"
// @Success 200 {object} domain.Dentist "Detalles del dentista"
// @Router /dentists/{id} [get]
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

// @Summary Crea un nuevo dentista
// @Description Crea un nuevo dentista con los datos proporcionados
// @Tags dentists
// @Accept json
// @Produce json
// @Param dentist body domain.Dentist true "Datos del Dentista"
// @Success 200 {object} domain.Dentist "Dentista creado"
// @Security ApiKeyAuth
// @Router /dentists [post]
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

// @Summary Actualiza los datos de un dentista
// @Description Actualiza los datos de un dentista con la información proporcionada
// @Tags dentists
// @Accept json
// @Produce json
// @Param id path int true "ID del Dentista"
// @Param dentist body domain.Dentist true "Datos del Dentista a actualizar"
// @Success 200 {object} domain.Dentist "Dentista actualizado exitosamente"
// @Security ApiKeyAuth
// @Router /dentists/{id} [put]
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

// @Summary Actualiza campos específicos de un dentista
// @Description Actualiza uno o más campos específicos de un dentista
// @Tags dentists
// @Accept json
// @Produce json
// @Param id path int true "ID del Dentista"
// @Success 200 {object} domain.Dentist "Dentista actualizado exitosamente"
// @Security ApiKeyAuth
// @Router /dentists/{id}/fields [patch]
func (h *dentistHandler) UpdateByField() gin.HandlerFunc {
	type Request struct {
		License  string `json:"license,omitempty"`
		Name     string `json:"name,omitempty"`
		LastName string `json:"last_name,omitempty"`
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

// @Summary Elimina un dentista
// @Description Elimina un dentista del sistema por su ID
// @Tags dentists
// @Accept json
// @Produce json
// @Param id path int true "ID del Dentista"
// @Success 200 {object} object "Mensaje de confirmación de eliminación"
// @Security ApiKeyAuth
// @Router /dentists/{id} [delete]
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
