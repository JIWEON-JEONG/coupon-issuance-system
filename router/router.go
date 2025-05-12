package router

import (
	"campaign-coupon-system/controller"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(router *fiber.App, campaignController *controller.CampaignController) *fiber.App {
	router.Get("/health-check", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
	router.Get("/v1/campaigns/:campaignId", campaignController.GetCampaign)
	router.Post("/v1/campaigns", campaignController.CreateCampaign)
	router.Post("/v1/campaigns/:campaignId/coupons/issue", campaignController.IssueCoupon)
	return router
}
