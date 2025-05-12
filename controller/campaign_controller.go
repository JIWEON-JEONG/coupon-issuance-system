package controller

import (
	"campaign-coupon-system/model"
	"campaign-coupon-system/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type CampaignController struct {
	campaignUseCase usecase.CampaignUseCase
}

func NewCampaignController(campaignUseCase usecase.CampaignUseCase) *CampaignController {
	return &CampaignController{campaignUseCase: campaignUseCase}
}

func (c *CampaignController) CreateCampaign(ctx *fiber.Ctx) error {
	var request CreateCampaignDto
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid Request Body"})
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	utcStartDateTime := request.StartDateTime.UTC()

	campaign := model.Campaign{
		AvailableCoupons: request.AvailableCoupons,
		StartDateTime:    utcStartDateTime,
	}

	createdCampaign, err := c.campaignUseCase.CreateCampaign(campaign)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	return ctx.Status(http.StatusOK).JSON(createdCampaign)
}
