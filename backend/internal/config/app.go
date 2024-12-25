package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrtzee/nextjs-go/internal/delivery/http"
	"github.com/mrtzee/nextjs-go/internal/delivery/http/middleware"
	"github.com/mrtzee/nextjs-go/internal/repository"
	"github.com/mrtzee/nextjs-go/internal/usecase"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type AppConfig struct {
	DB     *gorm.DB
	App    *fiber.App
	Log    *logrus.Logger
	Config *viper.Viper
}

func (cfg *AppConfig) Run() {
	// setup repositories
	productRepository := repository.NewProductRepository(cfg.Log)
	// setup use cases
	productUseCase := usecase.NewProductUsecase(productRepository, cfg.Log, cfg.DB)
	// setup controller
	productController := http.NewProductController(&productUseCase, cfg.Log)
	// setup middleware
	authMiddleware := middleware.NewAuth()
	routeConfig := http.Router{
		App:               cfg.App,
		ProductController: productController,
		AuthMiddleware:    authMiddleware,
	}
	routeConfig.Setup()
}
