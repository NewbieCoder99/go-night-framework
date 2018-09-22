/*
* @Author: Newbie Coder
* @Date:   2018-09-20 10:09:43
* @Last Modified by:   Newbie Coder
* @Last Modified time: 2018-09-22 23:59:15
*/
package main

import(
	"github.com/NewbieCoder99/go-night-framework/config"
	"github.com/NewbieCoder99/go-night-framework/server"
)

func main() {
	config.Init()
	config.GetMode()
	server.Init()
}