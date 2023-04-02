package controller

import (
	"context"
	"github.com/endr-i/promo-back-go/internal/model"
	"github.com/endr-i/promo-back-go/internal/repo"
	"github.com/jmoiron/sqlx"
)

type CouponController interface {
	Create(ctx context.Context, dto repo.CouponDto) (*model.Coupon, error)
}

type couponController struct {
	db *sqlx.DB
}

func (s *couponController) Create(ctx context.Context, dto repo.CouponDto) (*model.Coupon, error) {
	tx, err := s.db.Beginx()
	if err != nil {
		return nil, err
	}

	couponRepo := repo.NewCouponRepo(tx)

	coupon, err := couponRepo.Create(ctx, dto)
	if err != nil {
		return nil, err
	}

	return coupon, nil
}

func NewCouponController(db *sqlx.DB) CouponController {
	return &couponController{db: db}
}
