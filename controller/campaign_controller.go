package controller

import (
	"campaign-coupon-system/domain"
	"campaign-coupon-system/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type CampaignController struct {
	campaignUseCase domain.CampaignUseCase
}

func NewCampaignController(campaignUseCase domain.CampaignUseCase) *CampaignController {
	return &CampaignController{campaignUseCase: campaignUseCase}
}

// CreateCampaign method to create a new campaign.
// @Description Create a new campaign.
// @Summary create a new campaign
// @Tags Campaign
// @Accept json
// @Produce json
// @Param availableCoupons body integer true "availableCoupons"
// @Param startDateTime body string true "startDateTime"
// @Success 200 {object} model.CreateCampaignDto
// @Router /v1/campaigns [post]
func (c *CampaignController) CreateCampaign(ctx *fiber.Ctx) error {
	var request model.CreateCampaignDto
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request body"})
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
