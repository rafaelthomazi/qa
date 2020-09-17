package service

import (
	"github.com/rafaelthomazi/qa/qa/dao"
	"go.uber.org/zap"
)

// Config represents the Service parameters
type Config struct {
	HTTPPort string
	DAO      dao.Config
	Logger   *zap.Logger
}
