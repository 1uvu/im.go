package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"im/internal/api/routers"
	"im/internal/pkg/logger"
	"im/pkg/config"
)

type API struct {
}

func NewAPI() *API {
	return &API{}
}

func (api *API) Run() {
	conf := config.GetConfig().API

	gin.SetMode(conf.RunMode)

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", conf.ListenPort),
		Handler: routers.GetRouter(),
	}

	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-quit

		logger.Info("server is shutting down")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		httpServer.SetKeepAlivesEnabled(false)
		if err := httpServer.Shutdown(ctx); err != nil {
			logger.Errorf("shutdown server got error: %s", err.Error())
		}

		close(done)
	}()

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Errorf("listen server got error: %s", err.Error())
		}
	}()

	<-done
	logger.Info("server closed")
}
