/*
* @Author: Newbie Coder
* @Date:   2018-09-23 01:29:19
* @Last Modified by:   Newbie Coder
* @Last Modified time: 2018-09-23 02:55:38
*/
package main

import(
	"github.com/NewbieCoder99/go-night-framework/database/migrations"
)

func main() {
	migrations.UserMigrate()
	migrations.RoleMigrate()
	migrations.RoleUserMigrate()
}