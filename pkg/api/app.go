package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tinwritescode/gin-tin/pkg/config"
	"github.com/tinwritescode/gin-tin/pkg/handler"
	"github.com/tinwritescode/gin-tin/pkg/model"
	"github.com/tinwritescode/gin-tin/pkg/repository"
	"github.com/tinwritescode/gin-tin/pkg/service"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type App struct {
	router *gin.Engine
	cfg    *config.Config
}

func NewApp(cfg *config.Config) (*App, error) {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to the database:", err)
	}

	db.AutoMigrate(&model.Book{})
	repo := repository.NewBookRepository(db)
	svc := service.NewBookService(repo, db)
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
