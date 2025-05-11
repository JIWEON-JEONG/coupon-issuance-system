package main

import (
	"campaign-coupon-system/configuration"
	"campaign-coupon-system/controller"
	"campaign-coupon-system/repository"
	"campaign-coupon-system/router"
	"campaign-coupon-system/usecase"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
)

// @title API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	loadConfig, err := configuration.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load enviroment variables", err)
	}
	//configuration
	db := configuration.ConnectionMySQL(&loadConfig)
	cache := configuration.ConnectionRedis(&loadConfig)

	startServer(db, cache)
}

func startServer(db *gorm.DB, cache *redis.Client) {
	app := fiber.New()

	// swagger settings
	swaggerCfg := swagger.Config{
		BasePath: "/v1",
		FilePath: "./docs/swagger.yaml",
		Path:     "docs",
	}
	app.Use(swagger.New(swaggerCfg))

	campaignRepo := repository.NewCampaignRepository(db, cache)
	campaignUseCase := usecase.NewCampaignUseCase(campaignRepo)
	campaignController := controller.NewCampaignController(campaignUseCase)
	routes := router.NewRouter(app, campaignController)

	err := routes.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
