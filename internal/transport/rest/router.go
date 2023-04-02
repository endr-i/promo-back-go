package rest

import (
	"fmt"
	"github.com/endr-i/promo-back-go/internal/config"
	"github.com/endr-i/promo-back-go/internal/controller"
	v1 "github.com/endr-i/promo-back-go/internal/transport/rest/v1"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	config      *config.RestServerConfig
	logger      *logrus.Logger
	router      *chi.Mux
	controllers *v1.Controllers
}

func New(cfg *config.RestServerConfig, logger *logrus.Logger, db *sqlx.DB) *Server {
	return &Server{
		config: cfg,
		logger: logger,
		router: chi.NewRouter(),
		controllers: &v1.Controllers{
			CouponController:     controller.NewCouponController(db),
			CouponTypeController: controller.NewCouponTypeController(db),
		},
	}
}

func (s *Server) Start() error {
	s.configureRouter()
	fmt.Println("here", s.config.Address)

	return http.ListenAndServe(s.config.Address, s.router)
}

func (s *Server) configureRouter() {
	s.router.Get("/ping", handlePing())
	s.router.Route("/v1", v1.ConfigureRouter(s.controllers))
}

func handlePing() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("ping"))
	}
}
