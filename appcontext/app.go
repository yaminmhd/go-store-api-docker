package appcontext

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/yaminmhd/go-hardware-store/config"
	"github.com/yaminmhd/go-hardware-store/model"
)

type applicationContext struct {
	db *gorm.DB
}

var applicationcontext *applicationContext

func initDB() *gorm.DB {
	db, err := gorm.Open("mysql", config.Database().ConnectionURL())
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Initiate() {
	db := initDB()
	applicationcontext = &applicationContext{
		db: db,
	}
}

func GetDB() *gorm.DB {
	return applicationcontext.db
}

func CreateTables() {
	db := GetDB()
	transaction := db.Begin()
	err := transaction.Debug().DropTableIfExists(&model.Product{}).Error
	if err != nil {
		transaction.Rollback()
		log.Fatal(err)
	}
	err = transaction.Debug().CreateTable(&model.Product{}).Error
	if err != nil {
		transaction.Rollback()
		log.Fatal(err)
	}
	err = transaction.Commit().Error
	if err != nil {
		transaction.Rollback()
		log.Fatal(err)
	}
}
