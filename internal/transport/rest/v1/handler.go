package v1

import (
	"encoding/json"
	"github.com/endr-i/promo-back-go/internal/controller"
	"github.com/endr-i/promo-back-go/internal/repo"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
)

type Controllers struct {
	controller.CouponController
	controller.CouponTypeController
}

// ConfigureRouter todo: split router, unify & make dto objects with proto
func ConfigureRouter(controllers *Controllers) func(r chi.Router) {
	return func(r chi.Router) {
		r.Route("/coupon", func(rCoupon chi.Router) {
			rCoupon.Post("/", func(writer http.ResponseWriter, request *http.Request) {
				bytes, err := io.ReadAll(request.Body)
				if err != nil {
					http.Error(writer, err.Error(), http.StatusBadRequest)
					return
				}

				var dto repo.CouponDto
				err = json.Unmarshal(bytes, &dto)
				if err != nil {
					http.Error(writer, err.Error(), http.StatusBadRequest)
					return
				}

				coupon, err := controllers.CouponController.Create(request.Context(), dto)
				if err != nil {
					http.Error(writer, err.Error(), http.StatusBadRequest)
					return
				}

				response, err := json.Marshal(coupon)
				if err != nil {
					http.Error(writer, err.Error(), http.StatusBadRequest)
					return
				}

				writer.Write(response)
			})
		})

		r.Route("/coupon-type", func(rCouponType chi.Router) {
			rCouponType.Post("/", func(writer http.ResponseWriter, request *http.Request) {
				bytes, err := io.ReadAll(request.Body)
				if err != nil {
					http.Error(writer, err.Error(), http.StatusBadRequest)
					return
				}

				var dto repo.CouponTypeDto
				err = json.Unmarshal(bytes, &dto)
				if err != nil {
					http.Error(writer, err.Error(), http.StatusBadRequest)
					return
				}

				coupon, err := controllers.CouponTypeController.Create(request.Context(), dto)
				if err != nil {
					http.Error(writer, err.Error(), http.StatusBadRequest)
					return
				}

				response, err := json.Marshal(coupon)
				if err != nil {
					http.Error(writer, err.Error(), http.StatusBadRequest)
					return
				}

				writer.Write(response)
			})
		})
	}
}
