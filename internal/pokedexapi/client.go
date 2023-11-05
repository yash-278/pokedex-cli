package pokedexapi

import (
	"net/http"
	"time"

	"github.com/yash-278/pokedexcli/internal/pokedex_cache"
)

// Client -
type Client struct {
	cache      pokedex_cache.Cache
	httpClient http.Client
}

// NewClient -
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokedex_cache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
