package server

import (
	"context"
	"github.com/endr-i/promo-back-go/internal/config"
	"github.com/endr-i/promo-back-go/internal/transport/rest"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

type Server struct {
	config *config.Config
	logger *logrus.Logger
	rest   *rest.Server
	pg     *pgx.Conn
}

func New(config *config.Config) *Server {
	logger := logrus.New()

	return &Server{
		config: config,
		logger: logger,
		rest:   rest.New(&config.Rest, logger, &rest.Services{}),
	}
}

func (s *Server) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	if err := s.connectDB(); err != nil {
		return err
	}

	if err := s.rest.Start(); err != nil {
		return err
	}

	return nil
}

func (s *Server) Stop() {
	s.pg.Close(context.Background())
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
	pg, err := pgx.Connect(context.Background(), s.config.PG.ConnectionString)
	if err != nil {
		return err
	}

	s.pg = pg

	return nil
}
