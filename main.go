package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/locona/gin-spatium/router"
	"github.com/spf13/viper"
)

var (
	env  = flag.String("e", "local", "env mode")
	mode = flag.String("m", "debug", "gin mode")
)

func init() {
	flag.Parse()
	viper.AddConfigPath("./config")
	switch *env {
	case "dev":
		viper.SetConfigName("dev")
	case "test":
		viper.SetConfigName("test")
	case "staging":
		viper.SetConfigName("staging")
	case "prod":
		viper.SetConfigName("prod")
	default:
		viper.SetConfigName("dev")
	}
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("config error")
	}
}

func main() {
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())
	router.V1(engine)
	engine.Run(fmt.Sprint(":", viper.GetString("port")))
}
