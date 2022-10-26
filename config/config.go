package config

import "os"

var (
	// Frontend_Endpoint = "<FRONTEND_ENDPOINT>"
	// Backend_Endpoint  = "<BACKEND_ENDPOINT>"
	// Key               = "<API_KEY>"    // API_KEY acquired from BITGIN
	// Secret            = "<SECRET_KEY>" // SECRET_KEY acquired from BITGIN
	Frontend_Endpoint = "https://stage.bitgin.app/fiat-as-a-service"
	Backend_Endpoint  = "https://api.bitgin.app"
	Key               = os.Getenv("KEY")    // API_KEY acquired from BITGIN
	Secret            = os.Getenv("SECRET") // SECRET_KEY acquired from BITGIN
)
