package config

var (
	// Frontend_Endpoint = "<FRONTEND_ENDPOINT>"
	// Backend_Endpoint  = "<BACKEND_ENDPOINT>"
	// Key               = "<API_KEY>"    // API_KEY acquired from BITGIN
	// Secret            = "<SECRET_KEY>" // SECRET_KEY acquired from BITGIN
	Frontend_Endpoint = "https://stage.bitgin.app/fiat-as-a-service"
	Backend_Endpoint  = "https://api.bitgin.app"
)

func GetConfig() *config {
	return defaultConfig
}

var defaultConfig *config

type config struct {
	FrontendEndpoint string
	BackendEndpoint  string
	Key              string
	Secret           string
}

func Initialize(key, secret string) {
	defaultConfig = &config{
		FrontendEndpoint: Frontend_Endpoint,
		BackendEndpoint:  Backend_Endpoint,
		Key:              key,
		Secret:           secret,
	}
}
