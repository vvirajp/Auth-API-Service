package config

import (
	"time"

	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
)

func LoadAllConfigs() {
	LoadApp()
	InitializeCache()
}

// FiberConfig func for configuration Fiber app.
func FiberConfig() fiber.Config {

	// Return Fiber configuration.
	return fiber.Config{
		ReadTimeout:                  60 * time.Second,
		WriteTimeout:                 60 * time.Second,
		IdleTimeout:                  90 * time.Second,
		JSONEncoder:                  json.Marshal,
		JSONDecoder:                  json.Unmarshal,
		CaseSensitive:                true,
		BodyLimit:                    104857600,
		StrictRouting:                true,
		ServerHeader:                 "Fiber",
		AppName:                      "Go Fiber v1.0.0",
		ReadBufferSize:               4096000,
		WriteBufferSize:              4096000,
		DisablePreParseMultipartForm: true,
	}
}
