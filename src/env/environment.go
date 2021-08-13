package env

import (
	"context"
	"github.com/gaursagarMT/starter/src/config"
	"gorm.io/gorm"
)

type Environment interface {
	GetConfig() config.Config
}

type EnvironmentImpl struct {
	MySQLClient gorm.Option
	Config      config.Config
}

func (e EnvironmentImpl) GetMySQLClient() gorm.Option {
	return e.MySQLClient
}

func (e EnvironmentImpl) GetConfig() config.Config {
	return e.Config
}

func Init(ctx context.Context, config config.Config) (Environment, error) {
	// mysqlDSN := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

	return &EnvironmentImpl{
		Config: config,
	}, nil
}
