package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/BohdanCh-w/DSO-back/cmd/dso-back/api"
	"github.com/BohdanCh-w/DSO-back/config"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	exitCode := 0

	mainLogger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	if err := appMain(mainLogger); err != nil {
		mainLogger.Fatalf("dso-back: error %s", err)

		exitCode = 1
	}

	os.Exit(exitCode)
}

func appMain(mainLogger *log.Logger) error {
	var appConfig config.AppConfig

	if err := envconfig.Process("", &appConfig); err != nil {
		return err
	}

	if err := appConfig.Validate(); err != nil {
		return err
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	webAPI := http.Server{
		Addr: fmt.Sprintf("%s:%d", appConfig.BindIP, appConfig.BindPort),
		Handler: api.API(
			api.Options{
				Shutdown:  shutdown,
				AppConfig: appConfig,
				Logger:    mainLogger,
			},
		),
	}

	serverErrors := make(chan error, 1)
	go func() {
		mainLogger.Printf("listening API address %s", webAPI.Addr)
		serverErrors <- webAPI.ListenAndServe()
	}()

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)
	case sig := <-shutdown:
		mainLogger.Printf("start shutdown after signal %s", sig.String())

		return nil
	}
}
