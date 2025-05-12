package domain

import (
	"campaign-coupon-system/model"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// Mock CouponRepository for testing
type mockCouponRepository struct {
	mock.Mock
}

func (m *mockCouponRepository) Insert(ctx context.Context, tx *gorm.DB, coupons []model.Coupon) error {
	args := m.Called(ctx, tx, coupons)
	return args.Error(0)
}

func (m *mockCouponRepository) FindCouponDtoByCampaignIdOrNil(campaignId int) ([]model.CouponDto, error) {
	args := m.Called(campaignId)
	// Get(0)로 가져온 []*model.CouponDto를 []model.CouponDto로 변환
	results := args.Get(0).([]*model.CouponDto)
	couponDtos := make([]model.CouponDto, len(results))
	for i, result := range results {
		if result != nil {
			couponDtos[i] = *result // 포인터 역참조하여 값 복사
		}
	}
	return couponDtos, args.Error(1)
}

func (m *mockCouponRepository) InsertIssuedCoupon(issuedCoupon model.IssuedCoupon) error {
	args := m.Called(issuedCoupon)
	return args.Error(0)
}

func (m *mockCouponRepository) FindIssuedCouponByCampaignId(campaignId int) ([]model.IssuedCoupon, error) {
	args := m.Called(campaignId)
	return args.Get(0).([]model.IssuedCoupon), args.Error(1)
}

// @description 캠페인내의 생성된 쿠폰의 코드값은 모두 고유해야합니다.
func TestCouponService_GenerateCoupons_Uniqueness(t *testing.T) {
	// Arrange
	mockRepo := new(mockCouponRepository)
	service := NewCouponService(mockRepo)
	ctx := context.Background()
	tx := &gorm.DB{}
	// 많은 수의 쿠폰 생성 시도
	availableCoupons := 1000
	campaign := model.Campaign{
		ID:               1,
		AvailableCoupons: availableCoupons,
		StartDateTime:    time.Now().UTC(),
	}
	mockRepo.On("Insert", ctx, tx, mock.AnythingOfType("[]model.Coupon")).Return(nil).Once()

	// Act
	err := service.GenerateCoupons(ctx, tx, campaign)

	// Assert
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)

	args := mockRepo.Calls[0].Arguments.Get(2).([]model.Coupon)
	assert.Len(t, args, availableCoupons)

	generatedCodes := make(map[string]struct{})
	for _, coupon := range args {
		if _, exists := generatedCodes[coupon.Code]; exists {
			assert.Fail(t, "중복된 쿠폰 코드가 생성되었습니다: %s", coupon.Code)
		}
		generatedCodes[coupon.Code] = struct{}{}
		assert.NotEmpty(t, coupon.Code)
		assert.Equal(t, 1, coupon.CampaignID)
		assert.Equal(t, campaign.StartDateTime, coupon.AvailableFrom)
	}
}
