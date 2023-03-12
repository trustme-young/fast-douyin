package main

import (
	// "database/sql"
	"github.com/gin-gonic/gin"
	"github.com/ikuraoo/fastdouyin/configure"

	// "fmt"
	"github.com/ikuraoo/fastdouyin/service"
	"github.com/ikuraoo/fastdouyin/util"
)

func main() {
	configure.InitConfig()
	util.DB()

	go service.RunMessageServer()

	r := gin.Default()
	util.InitLogger()

	initRouter(r)

	// db, err := sql.Open("mysql", "root:root@/godb?charset=utf8&parseTime=True&loc=Local")
	// defer db.Close()
	// if err != nil {
	// 	fmt.Println("数据库崩了 ")
	// }

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
