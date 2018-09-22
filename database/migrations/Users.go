/*
* @Author: Newbie Coder
* @Date:   2018-09-22 23:28:00
* @Last Modified by:   Newbie Coder
* @Last Modified time: 2018-09-23 02:56:41
*/
package migrations

import (
	"fmt"
	"log"
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/NewbieCoder99/go-night-framework/config"
)

type User struct {
	Id 			int 	`gorm:"primary_key"`
	Username 	string 	`gorm:"type:varchar(191);not null"`
	Email 		string	`gorm:"type:varchar(191);unique;not null"`
	Password 	string 	`gorm:"type:varchar(191);not null"`
	Created_at	time.Time
	Updated_at	time.Time
}

func UserMigrate() {

	var TableName string = "users"

	db, err := gorm.Open(config.DBConf("driver"),config.DBConf("setup"))

	if err != nil { log.Fatalln(err) }
	fmt.Println("- Migrating Table:",TableName)
	db.AutoMigrate(&User{})
	fmt.Println("- Migrated Table:",TableName)
}