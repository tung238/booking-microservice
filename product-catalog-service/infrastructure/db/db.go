package db

import (
	"errors"
	"fmt"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	infra "github.com/tung238/booking-microservice/product-catalog-service/infrastructure"

	"github.com/jinzhu/gorm"
)

// DB Global DB connection
var DB *gorm.DB

func init() {
	var err error

	dbConfig := infra.Config.DB
	if infra.Config.DB.Adapter == "mysql" {
		DB, err = gorm.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name))
		// DB = DB.Set("gorm:table_options", "CHARSET=utf8")
	} else if infra.Config.DB.Adapter == "postgres" {
		DB, err = gorm.Open("postgres", fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=disable", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Name))
	} else if infra.Config.DB.Adapter == "sqlite" {
		DB, err = gorm.Open("sqlite3", fmt.Sprintf("%v/%v", os.TempDir(), dbConfig.Name))
	} else {
		panic(errors.New("not supported database adapter"))
	}

	if err == nil {
		if os.Getenv("DEBUG") != "" {
			DB.LogMode(true)
		}

		// l10n.RegisterCallbacks(DB)
		// sorting.RegisterCallbacks(DB)
		// validations.RegisterCallbacks(DB)
		// media.RegisterCallbacks(DB)
		// publish2.RegisterCallbacks(DB)
	} else {
		panic(err)
	}
}
