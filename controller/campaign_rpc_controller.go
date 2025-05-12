package controller

import (
	proto "campaign-coupon-system/controller/gen/campaign/v1"
	"campaign-coupon-system/model"
	"campaign-coupon-system/usecase"
	"context"
	"time"
)

type CampaignRpcController struct {
	proto.UnimplementedCampaignServiceServer
	useCase usecase.CampaignUseCase
}

func NewCampaignRpcController(uc usecase.CampaignUseCase) *CampaignRpcController {
	return &CampaignRpcController{useCase: uc}
}

func (s *CampaignRpcController) CreateCampaign(ctx context.Context, req *proto.CreateCampaignRequest) (*proto.CreateCampaignResponse, error) {
	startTime, err := time.Parse(time.RFC3339, req.StartDateTime)
	if err != nil {
		return nil, err
	}

	campaign := model.Campaign{
		AvailableCoupons: int(req.AvailableCouponCount),
		StartDateTime:    startTime.UTC(),
	}

	created, err := s.useCase.CreateCampaign(ctx, campaign)
	if err != nil {
		return nil, err
	}

	return &proto.CreateCampaignResponse{
		CampaignId: uint32(created.ID),
	}, nil
}

func (s *CampaignRpcController) IssueCoupon(ctx context.Context, req *proto.IssueCouponRequest) (*proto.IssueCouponResponse, error) {
	code, err := s.useCase.IssueCoupon(int(req.CampaignId), int(req.UserId))
	if err != nil {
		return nil, err
	}
	return &proto.IssueCouponResponse{Code: code}, nil
}

func (s *CampaignRpcController) GetCampaign(ctx context.Context, req *proto.CampaignRequest) (*proto.CampaignResponse, error) {
	dto, err := s.useCase.GetCampaignInfo(int(req.CampaignId))
	if err != nil {
		return nil, err
	}

	issuedCodes := make([]string, len(dto.IssuedCodes))
	copy(issuedCodes, dto.IssuedCodes)

	return &proto.CampaignResponse{
		CampaignId:    uint32(dto.CampaignID),
		IssuedCodes:   issuedCodes,
		StartDateTime: dto.StartDateTime.Format(time.RFC3339),
	}, nil
}
