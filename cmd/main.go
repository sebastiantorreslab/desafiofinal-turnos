package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	handlerDentist "github.com/sebastiantorreslab/desafiofinal-turnos/cmd/server/handler/dentist"
	"github.com/sebastiantorreslab/desafiofinal-turnos/internal/dentist"
	"github.com/sebastiantorreslab/desafiofinal-turnos/pkg/store"
)

func main() {

	db, err := store.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	storage := store.NewSqlStore(db)
	repository := dentist.NewRepository(storage)
	service := dentist.NewService(repository)
	dentistHandler := handlerDentist.NewDentistHandler(service)

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





	server.Run(":8080")
}
