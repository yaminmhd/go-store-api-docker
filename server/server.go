package server

import (
	"context"
	"github.com/urfave/negroni"
	"github.com/yaminmhd/go-hardware-store/config"
	"github.com/yaminmhd/go-hardware-store/log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func listenServer(apiServer *http.Server) {
	err := apiServer.ListenAndServe()
	if err != http.ErrServerClosed {
		log.Log.Fatal(err.Error())
	}
}

func waitForShutdown(apiServer *http.Server) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig,
		syscall.SIGINT,
		syscall.SIGTERM)
	_ = <-sig
	log.Log.Info("API server shutting down")
	// Finish all apis being served and shutdown gracefully
	apiServer.Shutdown(context.Background())
	log.Log.Info("API server shutdown complete")
}

func StartServer() {
	log.Log.Info("Starting Hardware store")
	dep := Init()
	router := Router(dep)
	handlerFunc := router.ServeHTTP

	n := negroni.New()
	n.UseHandlerFunc(handlerFunc)
	portInfo := ":" + strconv.Itoa(config.Port())
	server := &http.Server{Addr: portInfo, Handler: n}
	go listenServer(server)

	waitForShutdown(server)
}
