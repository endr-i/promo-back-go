package controller

import (
	"context"
	"github.com/endr-i/promo-back-go/internal/model"
	"github.com/endr-i/promo-back-go/internal/repo"
	"github.com/jmoiron/sqlx"
)

type CouponTypeController interface {
	Create(ctx context.Context, dto repo.CouponTypeDto) (*model.CouponType, error)
}

type couponTypeController struct {
	db *sqlx.DB
}

func (c *couponTypeController) Create(ctx context.Context, dto repo.CouponTypeDto) (*model.CouponType, error) {
	tx, err := c.db.Beginx()
	if err != nil {
		tx.Rollback()

		return nil, err
	}

	couponTypeRepo := repo.NewCouponTypeRepo(tx)

	res, err := couponTypeRepo.Create(ctx, dto)
	if err != nil {
		tx.Rollback()

		return nil, err
	}

	tx.Commit()

	return res, nil
}

func NewCouponTypeController(db *sqlx.DB) CouponTypeController {
	return &couponTypeController{
		db: db,
	}
}
