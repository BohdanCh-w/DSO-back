package api

import (
	"log"
	"net/http"
	"os"

	"github.com/BohdanCh-w/DSO-back/config"
	"github.com/julienschmidt/httprouter"
)

type Options struct {
	Shutdown  chan os.Signal
	AppConfig config.AppConfig
	Logger    *log.Logger
}

func API(opts Options) http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/lab1", lab1_func1(opts.Logger, opts.AppConfig.SaveLocation))

	return router
}
