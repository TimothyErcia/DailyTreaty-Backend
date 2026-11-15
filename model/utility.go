package model

type Config struct {
	Host          string
	Port          string
	ConnectionURL string
}

type GenericResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}
