/*
* @Author: Newbie Coder
* @Date:   2018-09-22 23:28:00
* @Last Modified by:   Newbie Coder
* @Last Modified time: 2018-09-23 02:59:26
*/
package migrations

import (
	"fmt"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/NewbieCoder99/go-night-framework/config"
)

type RoleUser struct {
	Id 			int 	`gorm:"primary_key"`
	User_id		int		`gorm:"unsigned"`
	Role_id		int		`gorm:"unsigned"`
}

func RoleUserMigrate() {

	var TableName string = "role_users"

	db, err := gorm.Open(config.DBConf("driver"), config.DBConf("setup"))

	if err != nil { log.Fatalln(err) }
	fmt.Println("- Migrating Table:",TableName)
	db.AutoMigrate(&RoleUser{})
	fmt.Println("- Migrated Table:",TableName)
	fmt.Println("- Add Foreign Key From table :",TableName)
	db.Model(&RoleUser{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&RoleUser{}).AddForeignKey("role_id", "roles(id)", "CASCADE", "CASCADE")
}