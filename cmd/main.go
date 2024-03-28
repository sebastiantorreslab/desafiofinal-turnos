package main

import (
	"log"

	"github.com/joho/godotenv"
	_ "github.com/sebastiantorreslab/desafiofinal-turnos/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	handlerDentist "github.com/sebastiantorreslab/desafiofinal-turnos/cmd/server/handler/dentist"
	handlerPatient "github.com/sebastiantorreslab/desafiofinal-turnos/cmd/server/handler/patient"
	handlerShift "github.com/sebastiantorreslab/desafiofinal-turnos/cmd/server/handler/shift"
	"github.com/sebastiantorreslab/desafiofinal-turnos/internal/dentist"
	"github.com/sebastiantorreslab/desafiofinal-turnos/internal/patient"
	"github.com/sebastiantorreslab/desafiofinal-turnos/internal/shift"
	"github.com/sebastiantorreslab/desafiofinal-turnos/pkg/middleware"
	"github.com/sebastiantorreslab/desafiofinal-turnos/pkg/store"
)

// @title Desafío integrador Golang
// @version 3.0
// @description Sistema de reserva de turnos para odontologos y pacientes
// @termsOfService

// @contact.name Sebastian Torres, Sebastian Alfonso
// @contact.url	https://github.com/sebastiantorreslab/desafiofinal-turnos/tree/main

// @license.name Digital House
// @license.url https://www.digitalhouse.com/co?utm_source=Google&utm_medium=paidsearch&utm_campaign=Clic&utm_term=Institucional&utm_content=institucional-institucional-branding-home-adresponsive-awareness-brandkws-none-exactas-adtext-none-exactas-ad1-latam-search&gad_source=1&gclid=CjwKCAjwh4-wBhB3EiwAeJsppJ2OfeYOwheVflmB-kJD2ctgZIZwplWV89rVP21dChHQdJtreVHUZRoCzOsQAvD_BwE
// @host localhost:8080
// @BasePath /api

// @SecurityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err.Error())
	}

	db, err := store.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	storageDentist := store.NewSqlDentistStore(db)
	repositoryDentist := dentist.NewDentistRepository(storageDentist)
	serviceDentist := dentist.NewDentistService(repositoryDentist)
	dentistHandler := handlerDentist.NewDentistHandler(serviceDentist)

	storagePatient := store.NewSqlStorePatient(db)
	repositoryPatient := patient.NewPatientRepository(storagePatient)
	servicePatient := patient.NewService(repositoryPatient)
	patientHandler := handlerPatient.NewPatientHandler(servicePatient)

	storageShift := store.NewSqlShiftStore(db)
	repositoryShift := shift.NewShiftRepository(storageShift)
	serviceShift := shift.NewShiftService(repositoryShift)
	shiftHandler := handlerShift.NewShiftHandler(serviceShift)

	server := gin.Default()

	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.GET("/ping", func(c *gin.Context) { c.String(200, "server OK") })

	dentists := server.Group("/dentists")

	dentists.GET("", dentistHandler.GetAll())
	dentists.GET(":id", dentistHandler.GetById())
	dentists.POST("", middleware.Authentication(), dentistHandler.Create())
	dentists.PUT(":id", middleware.Authentication(), dentistHandler.Update())
	dentists.PATCH(":id", middleware.Authentication(), dentistHandler.UpdateByField())
	dentists.DELETE(":id", middleware.Authentication(), dentistHandler.Delete())

	shifts := server.Group("/shifts")

	shifts.GET("", shiftHandler.GetAll())
	shifts.GET(":id", shiftHandler.GetById())
	shifts.POST("", middleware.Authentication(), shiftHandler.Create())
	shifts.PUT(":id", middleware.Authentication(), shiftHandler.Update())
	shifts.PATCH(":id", middleware.Authentication(), shiftHandler.UpdateByField())
	shifts.DELETE(":id", middleware.Authentication(), shiftHandler.Delete())

	// Rutas para pacientes
	patients := server.Group("/patients")
	{
		patients.GET("", patientHandler.GetAll())
		patients.GET(":id", patientHandler.GetById())
		patients.POST("", middleware.Authentication(), patientHandler.Create())
		patients.PUT(":id", middleware.Authentication(), patientHandler.Update())
		patients.PATCH(":id", middleware.Authentication(), patientHandler.UpdateByField())
		patients.DELETE(":id", middleware.Authentication(), patientHandler.Delete())
	}

	server.Run(":8080")
}
