package system

import (
	"net/http"

	"github.com/limmiehoang/circleci-demo/config"
	"github.com/rs/cors"
	"github.com/zenazn/goji/web"
)

type Middleware struct {
	cfg *config.Config
}

func NewMiddleware(cfg *config.Config) Middleware {
	return Middleware{
		cfg: cfg,
	}
}

func (m Middleware) ApplyCors(c *web.C, h http.Handler) http.Handler {
	cors := cors.New(cors.Options{
		AllowedOrigins:   m.cfg.Server.WhiteList,
		AllowedHeaders:   []string{"Accept", "Content-Type"},
		AllowCredentials: true,
	})

	return cors.Handler(h)
}
