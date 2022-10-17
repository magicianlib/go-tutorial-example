package main

//
// Shutdown methodroute server graceful(Implement the methodroute server using the gin)
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

	"github.com/gin-gonic/gin"
)

const (
	service = "GracefulSrv"
)

var (
	port            = flag.Int("port", 8080, "methodroute port to listen on")
	shutdownTimeout = flag.Duration("shutdown-timeout", 10*time.Second, "shutdown timeout (5s,5m,5h) before connections are cancelled")
)

func main() {

	fmt.Println("PID:", os.Getpid())

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		// sleep
		time.Sleep(30 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	router.GET("/shutdown", func(g *gin.Context) {
		process, err := os.FindProcess(os.Getpid())
		if err != nil {
			log.Println("Cannot find PID:", os.Getpid(), err)
			return
		}

		// kill -15
		process.Signal(syscall.SIGTERM)

		g.String(http.StatusOK, "ok")
	})

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: router,
	}

	// Run go methodroute srv in background
	go func() {
		log.Printf("%s listening on 0.0.0.0:%d with %v timeout", service, *port, *shutdownTimeout)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
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
