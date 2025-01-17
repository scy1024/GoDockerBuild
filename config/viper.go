package config

import (
	"GoDockerBuild/middleware"
	"fmt"
	"github.com/spf13/viper"
)

// 包级别的 Viper 实例变量
var devConfig *viper.Viper //开发测试用的配置文件
var config *viper.Viper    //正式版用的配置文件

func init() {
	// 初始化第一个 Viper 实例
	devConfig = viper.New()
	devConfig.SetConfigName("dev_config")
	devConfig.SetConfigType("json")
	devConfig.AddConfigPath("./config")
	if err := devConfig.ReadInConfig(); err != nil {
		fmt.Printf("Error reading dev_config file, %s\n", err)
	}

	// 初始化第二个 Viper 实例
	config = viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath("./config")
	if err := config.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s\n", err)
	}
}

// 读取线上环境的mysql配置
func GetDevMysqlConfig() middleware.Config {
	host := config.GetString("mysql.host")
	port := config.GetString("mysql.port")
	username := config.GetString("mysql.username")
	password := config.GetString("mysql.password")
	database := config.GetString("mysql.database")
	myConfig := middleware.Config{
		User:   username,
		Pass:   password,
		Addr:   host,
		Port:   port,
		Dbname: database,
	}
	return myConfig
}
func GetMysqlConfig() middleware.Config {
	host := devConfig.GetString("mysql.host")
	port := devConfig.GetString("mysql.port")
	username := devConfig.GetString("mysql.username")
	password := devConfig.GetString("mysql.password")
	database := devConfig.GetString("mysql.database")
	myConfig := middleware.Config{
		User:   username,
		Pass:   password,
		Addr:   host,
		Port:   port,
		Dbname: database,
	}
	return myConfig
}

func Test() {
	// 设置配置文件名和存放路径

	// 读取配置文件
	if err := devConfig.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	// 获取并打印 MySQL 配置信息
	host := devConfig.GetString("mysql.host")
	port := devConfig.GetInt("mysql.port")
	username := devConfig.GetString("mysql.username")
	password := devConfig.GetString("mysql.password")
	database := devConfig.GetString("mysql.database")

	fmt.Println("MySQL Configuration:")
	fmt.Println("Host:", host)
	fmt.Println("Port:", port)
	fmt.Println("Username:", username)
	fmt.Println("Password:", password)
	fmt.Println("Database:", database)
}
