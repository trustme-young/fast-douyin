package main

import (
	"github.com/ikuraoo/fastdouyin/configure"
	"github.com/ikuraoo/fastdouyin/entity"
	"github.com/ikuraoo/fastdouyin/util"
	"github.com/spf13/viper"
)

func main() {
	configure.InitConfig()
	host := viper.GetString("datasource.host")
	println("host:", host)
	var Fav entity.Favourite
	util.DB().First(&Fav)
	println(&Fav)
}
