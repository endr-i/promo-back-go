package model

type CouponType struct {
	Id     string `db:"id"`
	Name   string `db:"name"`
	Chance int    `db:"chance"`
}
