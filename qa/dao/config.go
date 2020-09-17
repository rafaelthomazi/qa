package dao

import "go.uber.org/zap"

type Config struct {
	URI      string
	Database string
	Logger   *zap.Logger
}
