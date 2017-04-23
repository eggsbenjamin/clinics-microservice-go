package main

import (
	"log"
	"net/http"

	"github.com/eggsbenjamin/clinics-microservice-go/handlers"
	"github.com/eggsbenjamin/clinics-microservice-go/mappers"
	clinics "github.com/eggsbenjamin/clinics-microservice-go/services"
	"github.com/eggsbenjamin/clinics-microservice-go/utils"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	client := &http.Client{}
	utils := &utils.Utils{}
	mapper := &mappers.Mapper{}
	clinicsService := clinics.NewClinicsService(client, utils)
	clinicsHandlers := handlers.NewClinicsHandlers(clinicsService, mapper)

	router.GET("/clinics/postcode/:postcode", clinicsHandlers.ClinicsByPostcode)

	log.Fatal(http.ListenAndServe(":8080", router))
}
