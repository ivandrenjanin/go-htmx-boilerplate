package server

import (
	"context"
	"fmt"
	"go-htmx/internal/locator"
	"go-htmx/internal/route"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func Init(l locator.Locator) {
	cfg := l.GetConfig()

	if cfg.App.Env == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.LoadHTMLGlob("internal/templates/**/*.tmpl")
	router.Static("/assets", "./assets")

	// Attach middlewares
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	// Attach routes
	route.StaticPublicHandlers(router, l)
	route.ApiHandlers(router, l)

	rwTime := 10 * time.Second

	port := fmt.Sprintf(":%s", cfg.App.Port)

	s := &http.Server{
		Addr:           port,
		Handler:        router,
		ReadTimeout:    rwTime,
		WriteTimeout:   rwTime,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		// service connections
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)

	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	<-ctx.Done()
	log.Println("timeout of 1 seconds.")
	log.Println("Server exiting")
}
