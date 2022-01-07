package server

import (
	"github.com/endr-i/promo-back-go/internal/config"
	"github.com/endr-i/promo-back-go/internal/transport/rest"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"time"
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
	dbConfig, err := pgx.ParseConfig(s.config.PG.ConnectionString)
	if err != nil {
		return err
	}

	dbConfig.ConnectTimeout = time.Duration(s.config.PG.ConnectTimeout) * time.Millisecond
	dbConfig.PreferSimpleProtocol = s.config.PG.PreferSimpleProtocol

	db := stdlib.OpenDB(*dbConfig)
	pg := sqlx.NewDb(db, "pgx")

	pg.SetMaxOpenConns(s.config.PG.MaxOpenConns)
	pg.SetMaxIdleConns(s.config.PG.MaxIdleConns)

	s.pg = pg

	return nil
}
