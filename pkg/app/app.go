package app

import (
	"github.com/imakiri/playground/core"
	"github.com/imakiri/playground/data"
	"github.com/jackc/pgx/v4"
)

type Service struct {
	db       *pgx.Conn
	log      core.LogService
	config   core.ConfigApp
	configDB core.ConfigDB
}

func NewService(c core.Config) (*Service, error) {
	var s Service
	var err error

	s.config = c.App
	s.configDB = c.DB

	s.db, err = data.Connect(s.configDB)
	if err != nil {
		return nil, err
	}

	return &s, err
}