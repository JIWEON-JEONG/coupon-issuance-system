package controller

import (
	"campaign-coupon-system/model"
	"campaign-coupon-system/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type CampaignController struct {
	campaignUseCase usecase.CampaignUseCase
}

func NewCampaignController(campaignUseCase usecase.CampaignUseCase) *CampaignController {
	return &CampaignController{campaignUseCase: campaignUseCase}
}

func (c *CampaignController) GetCampaign(ctx *fiber.Ctx) error {
	campaignIdStr := ctx.Params("campaignId")
	campaignIdInt, err := strconv.Atoi(campaignIdStr)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid Campaign ID"})
	}
	campaignInfo, err := c.campaignUseCase.GetCampaignInfo(campaignIdInt)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return ctx.Status(http.StatusOK).JSON(campaignInfo)
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

	createdCampaign, err := c.campaignUseCase.CreateCampaign(ctx.UserContext(), campaign)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return ctx.Status(http.StatusOK).JSON(createdCampaign)
}

func (c *CampaignController) IssueCoupon(ctx *fiber.Ctx) error {
	var request IssueCouponDto
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid Request Body"})
	}
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	campaignIdStr := ctx.Params("campaignId")
	campaignIdInt, err := strconv.Atoi(campaignIdStr)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid Campaign ID"})
	}

	code, err := c.campaignUseCase.IssueCoupon(campaignIdInt, request.UserID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	if code == "" {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"message": "SoldOut Coupon."})
	}
	return ctx.Status(http.StatusOK).JSON(code)
}
