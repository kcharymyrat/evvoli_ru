package api

import "log"

type ApiConfig struct {
	Port int
	Env  string
}

type ApiApplication struct {
	Config ApiConfig
	Logger *log.Logger
}
