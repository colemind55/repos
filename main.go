package main

import (
	"log"
	"net/http"
	"os"

	"allsup.assessment/api/services/controllers"
	"github.com/joho/godotenv"
)

func main() {
	
	// Try to load .env file - Though remember file is optional; environmental
	// variable values may not be.
	godotenv.Load()

	controllers.RegisterControllers()
	http.ListenAndServe(getPort(), nil)
}

func getPort() string {
	port := ":3000"

	if os.Getenv("ASPNETCORE_PORT") != "" {// get environment variable that set by ACNM 
		port = ":" + os.Getenv("ASPNETCORE_PORT")
	}

	log.Print("PORT (", port, ")")

	return port
}