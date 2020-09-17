package service

import (
	"github.com/rafaelthomazi/qa/qa/dao"
	"go.uber.org/zap"
)

type Config struct {
	HTTPPort string
	DAO      dao.Config
	Logger   *zap.Logger
}
