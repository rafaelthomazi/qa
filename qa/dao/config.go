package dao

import "go.uber.org/zap"

// Config represents the DAO parameters
type Config struct {
	URI      string
	Database string
	Logger   *zap.Logger
}
