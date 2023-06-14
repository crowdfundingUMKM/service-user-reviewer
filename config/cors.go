package config

import (
	"time"

	"github.com/gin-contrib/cors"
)

func InitCors() cors.Config {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT"}
	config.AllowHeaders = []string{"Origin", "Authorization", "Content-Type"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowWildcard = true
	// hasil permintaan CORS akan disimpan di cache browser selama 12 jam sebelum browser melakukan permintaan baru ke server.
	config.MaxAge = 12 * time.Hour
	config.AllowFiles = true

	return config
}
