package storage

import (
	"context"
	"fmt"
	"github.com/gaursagarMT/starter/src/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLStorage struct {
	client *gorm.DB
}

func (m MySQLStorage) CreateUser() {
	panic("implement me")
}

func (m MySQLStorage) RemoveUser() {
	panic("implement me")
}

func (m MySQLStorage) UpdateUser() {
	panic("implement me")
}

func (m MySQLStorage) DeleteUser() {
	panic("implement me")
}

func (m MySQLStorage) CreatePlace() {
	panic("implement me")
}

func (m MySQLStorage) CreateTables(ctx context.Context) error {
	panic("implement me")
}

func GetMySQLStorage(ctx context.Context, config config.MySQLConfig) (IStorage, error) {
	mysqlDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Host,
		config.Password,
		config.Host,
		config.Port,
		config.Database)
	mysqlClient, _ := gorm.Open(mysql.New(mysql.Config{
		DSN: mysqlDSN,
	}), &gorm.Config{})
	return &MySQLStorage{client: mysqlClient}, nil
}
