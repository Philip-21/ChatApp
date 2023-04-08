package handlers

import (
	"chatapp/internal/config"
	"net/http"
)

type AppConfig struct {
	App *config.AppConfig
}

var Repo *AppConfig

func (m *AppConfig) Home(w http.ResponseWriter, r *http.Request) {}
