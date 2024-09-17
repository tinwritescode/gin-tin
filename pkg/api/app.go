package api

import (
	"github.com/tinwritescode/gin-tin/pkg/config"
	"github.com/tinwritescode/gin-tin/pkg/handler"
	"github.com/tinwritescode/gin-tin/pkg/repository"
	"github.com/tinwritescode/gin-tin/pkg/service"

	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine
	cfg    *config.Config
}

func NewApp(cfg *config.Config) (*App, error) {
	repo := repository.NewBookRepository()
	svc := service.NewBookService(repo)
	h := handler.NewHandler(svc)

	router := gin.New()
	h.SetupRoutes(router)

	return &App{
		router: router,
		cfg:    cfg,
	}, nil
}

func (a *App) Run() error {
	return a.router.Run(a.cfg.ServerAddress)
}
