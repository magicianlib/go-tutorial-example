package main

//
// Shutdown http server graceful(Implement the http server using the std)
//
//

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	service = "GracefulSrv"
)

var (
	port            = flag.Int("port", 8080, "http port to listen on")
	shutdownTimeout = flag.Duration("shutdown-timeout", 10*time.Second, "shutdown timeout (5s,5m,5h) before connections are cancelled")
)

func main() {
	flag.Parse()

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// sleep
		time.Sleep(30 * time.Second)
		w.Write([]byte("Welcome Go Http Server"))
		w.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		process, err := os.FindProcess(os.Getpid())
		if err != nil {
			log.Println("Cannot find PID:", os.Getpid(), err)
			return
		}

		// kill -15
		process.Signal(syscall.SIGTERM)

		w.Write([]byte("ok"))
		w.WriteHeader(http.StatusOK)
	})

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: mux,
	}

	// Run go http srv in background
	go func() {
		log.Printf("%s listening on 0.0.0.0:%d with %v timeout", service, *port, *shutdownTimeout)
		if err := srv.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of *shutdownTimeout seconds.
	quit := make(chan os.Signal)

	// kill -2 is syscall.SIGINT
	// kill (-15, or no param) default send syscall.SIGTERM
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	// signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	log.Printf("%s shutting down ...\n", service)

	ctx, cancel := context.WithTimeout(context.Background(), *shutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("Server exiting")
}
