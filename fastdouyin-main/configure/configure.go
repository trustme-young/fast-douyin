package configure

import (
	"os"

	"github.com/spf13/viper"
)

// // Config 配置数据结构体
// type Config struct {
// 	Mysql struct {
// 		User    string
// 		Host    string
// 		Port    string
// 		Dbname  string
// 		Passwd  string
// 		Charset string
// 	}
// }

// var ConfigData *Config

func InitConfig() {
	workDir, _ := os.Getwd()
	// println("workdir", workDir)
	viper.SetConfigName("configure")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/configure")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

// // GetCOnfigData 返回配置数据方法
// func GetCOnfigData() *Config {
// 	return ConfigData
// }
