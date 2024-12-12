package config

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var appCache *cache.Cache

// InitializeCache initializes the global cache instance
func InitializeCache() {
	appCache = cache.New(15*time.Minute, 30*time.Minute)
}

// GetCache returns the global cache instance
func GetCache() *cache.Cache {
	if appCache == nil {
		InitializeCache()
	}
	return appCache
}
