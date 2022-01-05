package rest

import (
	"github.com/endr-i/promo-back-go/internal/config"
	v1 "github.com/endr-i/promo-back-go/internal/transport/rest/v1"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Services struct {
}

type Server struct {
	config   *config.RestServerConfig
	logger   *logrus.Logger
	router   *chi.Mux
	services *Services
}

func New(cfg *config.RestServerConfig, logger *logrus.Logger, services *Services) *Server {
	return &Server{
		config:   cfg,
		logger:   logger,
		router:   chi.NewRouter(),
		services: services,
	}
}

func (s *Server) Start() error {
	s.configureRouter()

	return http.ListenAndServe(s.config.Address, s.router)
}

func (s *Server) configureRouter() {
	s.router.Get("/ping", handlePing())
	s.router.Route("/v1", v1.ConfigureRouter())
}

func handlePing() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("ping"))
	}
}
