package config

import (
	"github.com/spf13/viper"
	"log"
)

func Do()  {
	var configFile="config"
	viper.SetConfigName(configFile)

	viper.AddConfigPath("./demo/config/")
	err:=viper.ReadInConfig()
	if err!=nil{
		log.Fatalln(err.Error())
	}

	name:=viper.GetString("user.name")
	age:=viper.GetInt("user.age")

	println(name,age)
}