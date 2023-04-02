package server

import (
	"github.com/endr-i/promo-back-go/internal/config"
	"github.com/endr-i/promo-back-go/internal/connection/pg"
	"github.com/endr-i/promo-back-go/internal/transport/rest"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Server struct {
	config *config.Config
	logger *logrus.Logger
	rest   *rest.Server
	pg     *sqlx.DB
}

func New(config *config.Config) *Server {
	logger := logrus.New()

	return &Server{
		config: config,
		logger: logger,
		rest:   nil,
	}
}

func (s *Server) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	if err := s.connectDB(); err != nil {
		return err
	}

	if err := s.configureTransport(); err != nil {
		return err
	}

	if err := s.rest.Start(); err != nil {
		return err
	}

	s.logger.Info("stating server")

	return nil
}

func (s *Server) Stop() {
	s.pg.Close()
	s.logger.Info("stopping server")
}

func (s *Server) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *Server) connectDB() error {
	db, err := pg.NewDB(s.config.PG)

	if err != nil {
		return err
	}

	s.pg = db

	return nil
}

func (s *Server) configureTransport() error {
	s.rest = rest.New(&s.config.Rest, s.logger, s.pg)

	return nil
}
