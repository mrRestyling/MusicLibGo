package config

import "os"

func Host() string {
	host := os.Getenv("HOST_SONG")
	if host == "" {
		host = "0.0.0.0"

	}
	return host
}

func Port() string {
	port := os.Getenv("PORT_SONG")
	if port == "" {
		port = "8080"
	}
	return port
}
