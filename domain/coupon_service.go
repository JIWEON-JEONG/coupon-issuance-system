package domain

import (
	"campaign-coupon-system/model"
	"context"
	"fmt"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"time"
)

type CouponService interface {
	//GetByCampaignId(campaignId int) ([]model.Coupon, error)
	//특정 Campaign 에 사용되는 유니크한 Coupon들을 발급하고 저장합니다.
	GenerateCoupons(ctx context.Context, tx *gorm.DB, campaign model.Campaign) error
	IssueCoupon(campaignId int, userId int) (string, error)
	GetIssueCoupons(campaignId int) ([]model.IssuedCoupon, error)
}

func NewCouponService(repository CouponRepository) CouponService {
	return &couponService{repository: repository}
}

func (c *couponService) GetIssueCoupons(campaignId int) ([]model.IssuedCoupon, error) {
	return c.repository.FindIssuedCouponByCampaignId(campaignId)
}

func (c *couponService) IssueCoupon(campaignId int, userId int) (string, error) {
	couponDtos, err := c.repository.FindCouponDtoByCampaignIdOrNil(campaignId)
	if err != nil {
		return "", fmt.Errorf("failed to find available coupon: %w", err)
	}
	if couponDtos == nil {
		return "", nil
	}

	nowUTC := time.Now().UTC()
	couponDto := couponDtos[0]

	if couponDto.AvailableFrom.After(nowUTC) {
		return "", fmt.Errorf("coupon is not yet available. Available from: %s", couponDto.AvailableFrom.String())
	}

	issuedCoupon := model.IssuedCoupon{
		CampaignID:    campaignId,
		Code:          couponDto.Code,
		UserID:        userId,
		AvailableFrom: couponDto.AvailableFrom,
	}
	// 비동기로 히스토리 저장
	go func() {
		if err := c.repository.InsertIssuedCoupon(issuedCoupon); err != nil {
			log.Printf("Error saving coupon history (async): %+v, error=%v", issuedCoupon, err)
		}
	}()
	return couponDto.Code, nil
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
