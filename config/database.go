package config

import (
	"github.com/CemHarput/bookStoreGo/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB
const DATABASE_URI = `host=localhost port=5432 user=postgres password=cem123. dbname=bookStoreGo sslmode=disable`

func Connect() error {
    var err error

    Database, err = gorm.Open(postgres.Open(DATABASE_URI), &gorm.Config{
        SkipDefaultTransaction: true,
        PrepareStmt:            true,
    })

    if err != nil {
        panic(err)
    }

    Database.AutoMigrate(&model.Book{})
	Database.AutoMigrate(&model.PurchaseOrder{})

    return nil
}