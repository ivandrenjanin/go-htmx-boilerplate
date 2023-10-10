package server

import (
	"context"
	"fmt"
	"go-htmx/internal/locator"
	"go-htmx/internal/route"
	"go-htmx/pkg/session"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

func Init(l locator.Locator) {
	cfg := l.GetConfig()

	if cfg.App.Env == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	ses, err := session.NewSessionMiddleware(l.GetDB(), cfg)

	if err != nil {
		log.Fatalf("session init error: %s\n", err)
	}

	r := gin.Default()

	// Attach middlewares
	r.Use(ses)

	r.Use(csrf.Middleware(csrf.Options{
		IgnoreMethods: []string{"GET", "HEAD", "OPTIONS"},
		Secret:        cfg.App.AccessSecret,
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()
		}}))

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{cfg.App.Host}

	r.Use(cors.New(corsConfig))
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(helmet.Default())
	r.LoadHTMLGlob("internal/templates/**/*.tmpl")
	r.Static("/assets", "./assets")

	if cfg.App.Env != "PRODUCTION" {
		r.GET("/csrf", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"csrf_token": csrf.GetToken(c),
			})
		})
	}

	// Attach routes
	route.StaticPublicHandlers(r, l)
	route.ApiHandlers(r, l)

	rwTime := 10 * time.Second

	port := fmt.Sprintf(":%s", cfg.App.Port)

	s := &http.Server{
		Addr:           port,
		Handler:        r,
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
