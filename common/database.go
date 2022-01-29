package common

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hutaiyi.study/gin_vue/model"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	//driverName := "mysql"
	host := viper.GetString("datasource.localhost")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	userName := viper.GetString("datasource.userName")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True", userName, password, host, port, database, charset)

	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		log.Println("connect db err : ", err.Error())
	}

	db.AutoMigrate(&model.User{})
	DB = db
	return db

}

func GetDB() *gorm.DB {
	return DB
}
