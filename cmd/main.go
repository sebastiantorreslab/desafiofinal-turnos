package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	handlerDentist "github.com/sebastiantorreslab/desafiofinal-turnos/cmd/server/handler/dentist"
	handlerShift "github.com/sebastiantorreslab/desafiofinal-turnos/cmd/server/handler/shift"
	"github.com/sebastiantorreslab/desafiofinal-turnos/internal/dentist"
	"github.com/sebastiantorreslab/desafiofinal-turnos/internal/shift"
	"github.com/sebastiantorreslab/desafiofinal-turnos/pkg/store"
)

func main() {

	db, err := store.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	storageDentist := store.NewSqlDentistStore(db)
	repositoryDentist := dentist.NewDentistRepository(storageDentist)
	serviceDentist := dentist.NewDentistService(repositoryDentist)
	dentistHandler := handlerDentist.NewDentistHandler(serviceDentist)

	server := gin.Default()

	server.GET("/ping", func(c *gin.Context) { c.String(200, "server OK") })

	dentists := server.Group("/dentists")
	{

		dentists.GET("", dentistHandler.GetAll())
		dentists.GET(":id", dentistHandler.GetById())
		dentists.POST("", dentistHandler.Create())
		dentists.PUT(":id", dentistHandler.Update())
		dentists.PATCH(":id", dentistHandler.UpdateByField())
		dentists.DELETE(":id", dentistHandler.Delete())

	}

	storageShift := store.NewSqlShiftStore(db)
	repositoryShift := shift.NewShiftRepository(storageShift)
	serviceShift := shift.NewShiftService(repositoryShift)
	shiftHandler := handlerShift.NewShiftHandler(serviceShift)

	shifts := server.Group("/shifts")
	{

		shifts.GET("", shiftHandler.GetAll())
		shifts.GET(":id", shiftHandler.GetById())
		shifts.POST("", shiftHandler.Create())
		shifts.PUT(":id", shiftHandler.Update())
		shifts.PATCH(":id", shiftHandler.UpdateByField())
		shifts.DELETE(":id", shiftHandler.Delete())

	}

	

	server.Run(":8080")
}
