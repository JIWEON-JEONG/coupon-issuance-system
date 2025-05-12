package domain

import (
	"campaign-coupon-system/model"
	"context"
	"fmt"
	"gorm.io/gorm"
	"math/rand"
)

type CouponService interface {
	//특정 Campaign 에 해당하는 Coupon들을 모두 조회합니다.
	GetByCampaignId(campaignId int) ([]model.Coupon, error)
	//특정 Campaign 에 사용되는 유니크한 Coupon들을 발급하고 저장합니다.
	GenerateCoupons(ctx context.Context, tx *gorm.DB, campaign model.Campaign) error
}

func NewCouponService(repository CouponRepository) CouponService {
	return &couponService{repository: repository}
}

func (c *couponService) GetByCampaignId(campaignId int) ([]model.Coupon, error) {
	coupons, err := c.repository.FindByCampaignId(campaignId)
	return coupons, err
}

const (
	maxCodeLength = 10
	minCodeLength = 5
	maxAttempts   = 50
)

var hangulRunes = []rune("가나다라마바사아자차카타파하")
var digits = []rune("0123456789")

func (c *couponService) GenerateCoupons(ctx context.Context, tx *gorm.DB, campaign model.Campaign) error {
	coupons := make([]model.Coupon, 0, campaign.AvailableCoupons)
	generatedCodes := make(map[string]struct{})

	for len(coupons) < campaign.AvailableCoupons {
		attempts := 0
		var code string
		for {
			if attempts >= maxAttempts {
				return fmt.Errorf("쿠폰 코드 생성 실패: 중복으로 인해 %d번 시도 후에도 고유한 코드 생성 실패", maxAttempts)
			}
			code = generateCouponCode()
			if _, exists := generatedCodes[code]; !exists {
				break
			}
			attempts++
		}
		generatedCodes[code] = struct{}{}
		coupons = append(coupons, model.Coupon{
			CampaignID:    int(campaign.ID),
			Code:          code,
			AvailableFrom: campaign.StartDateTime,
		})
	}
	err := c.repository.Insert(ctx, tx, coupons)
	if err != nil {
		return err
	}
	return nil
}

// 랜덤 쿠폰 코드 생성
func generateCouponCode() string {
	length := rand.Intn(maxCodeLength-minCodeLength+1) + minCodeLength
	code := make([]rune, length)
	for i := 0; i < length; i++ {
		if rand.Intn(2) == 0 {
			code[i] = hangulRunes[rand.Intn(len(hangulRunes))]
		} else {
			code[i] = digits[rand.Intn(len(digits))]
		}
	}
	return string(code)
}

type couponService struct {
	repository CouponRepository
}
