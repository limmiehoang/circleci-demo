package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"syscall"

	"github.com/limmiehoang/circleci-demo/config"
	"github.com/limmiehoang/circleci-demo/system"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

var (
	configPath = flag.String("cfg", "/etc/circleci-demo/demo.toml", "a path to configuration file")
)

func main() {
	flag.Parse()

	_, err := os.Stat(*configPath)
	if err != nil {
		log.Printf("%s: %s\n", syscall.ENOENT.Error(), *configPath)
		os.Exit(int(syscall.ENOENT))
	}

	cfg, err := config.New(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	setup(cfg)

	flag.Set("bind", cfg.Server.Listen)
	goji.Serve()
}

func setup(cfg *config.Config) {
	m := system.NewMiddleware(cfg)

	goji.Use(m.ApplyCors)

	goji.Get("/ping", func(c web.C, w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, "pong")
	})
}
