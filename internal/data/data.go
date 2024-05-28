package data

import (
	"database/sql"
	"demo/internal/data/models"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(InitDB)

type Data struct {
	Db    *gorm.DB
	sqlDB *sql.DB
}

func InitDB() (*Data, error) {
	dsn := "root:root@tcp(localhost:3306)/allen?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&models.UserModel{})
	if err != nil {
		panic("failed to migrate database")
	}

	return &Data{Db: db}, nil
}
