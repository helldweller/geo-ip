package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/sync/errgroup"
)

var ctx, cancel = context.WithCancel(context.Background())
var group, groupCtx = errgroup.WithContext(ctx)
var server *http.Server

// var expiration = time.Second * 300

// Run is a like main function
func Run() {
	log.Info("Starting app")

	server = &http.Server{
		Addr:    conf.HTTPListenIPPort,
		Handler: nil,
		// BaseContext: ctx,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}

	group.Go(func() error {
		signalChannel := make(chan os.Signal, 1)
		defer close(signalChannel)
		signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
		select {
		case sig := <-signalChannel:
			log.Errorf("Received signal: %s", sig)
			if err := server.Shutdown(ctx); err != nil {
				log.Errorf("Received an error while shutting down the server: %s", err)
			}
			cancel()
		case <-groupCtx.Done():
			log.Error("Closing signal goroutine")
			if err := server.Shutdown(ctx); err != nil {
				log.Errorf("Received an error while shutting down the server: %s", err)
			}
			return groupCtx.Err()
		}
		return nil
	})

	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {

		fmt.Fprint(res, "OK")

	})

	group.Go(func() error {
		log.Infof("Starting web server on %s", conf.HTTPListenIPPort)
		err := server.ListenAndServe()
		return err
	})

	err := group.Wait()
	if err != nil {
		if errors.Is(err, context.Canceled) {
			log.Error("Context was canceled")
		} else {
			log.Errorf("Received error: %v\n", err)
		}
	} else {
		log.Error("Sucsessfull finished")
	}
}
