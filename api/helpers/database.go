package helpers

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func Connect(driver, user, password, port, host, database string) *gorm.DB {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		database)

	fmt.Println(dataSourceName)

	db, err := gorm.Open(driver, dataSourceName)
	if err != nil {
		panic(err.Error())
	}

	return db
}
