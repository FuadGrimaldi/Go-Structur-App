package server

import (
	"context"
	"fmt"
	"go-app/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
)

type Server struct {
	*echo.Echo
	cfg *config.Config
}

func NewServer(cfg *config.Config) *Server {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	return &Server{e, cfg}
}
func (s *Server) Run() {
	go func ()  {
		err := s.Start(fmt.Sprintf(":%s", s.cfg.Port))
		log.Fatal(err)
	} ()
}
func (s *Server) GracefulShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func() {
		if err := s.Shutdown(ctx); err != nil {
			s.Logger.Fatal(err)
		}
	}()
}