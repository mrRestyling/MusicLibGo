package config

import "os"

func Host() string {
	host := os.Getenv("HOST_SONG")
	if host == "" {
		host = "localhost"
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
