package mysql

import (
	"fmt"
	"time"

	"github.com/ysfglmzz/e-shop-microservices/catalog/config"
	"github.com/ysfglmzz/e-shop-microservices/catalog/pkg/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormConnect(cfg config.MysqlConfig) *gorm.DB {
	dsn := fmt.Sprintf(constants.MysqlConnectionFormat, cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	gormDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	sqlDb, _ := gormDb.DB()
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour)
	return gormDb
}
