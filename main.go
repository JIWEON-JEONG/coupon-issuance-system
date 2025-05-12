package main

import (
	"campaign-coupon-system/configuration"
	"campaign-coupon-system/controller"
	"campaign-coupon-system/domain"
	"campaign-coupon-system/repository"
	"campaign-coupon-system/router"
	"campaign-coupon-system/usecase"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
)

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

	campaignRepo := repository.NewCampaignRepository(db, cache)
	couponRepo := repository.NewCouponRepository(db, cache)
	campaignService := domain.NewCampaignService(campaignRepo)
	couponService := domain.NewCouponService(couponRepo)
	campaignUseCase := usecase.NewCampaignUseCase(db, campaignService, couponService)
	campaignController := controller.NewCampaignController(campaignUseCase)
	routes := router.NewRouter(app, campaignController)

	err := routes.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
