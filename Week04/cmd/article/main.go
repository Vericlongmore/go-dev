package main

import (
	"dev/wire/cmd/wire"
	"github.com/spf13/viper"
	"log"
)

func main() {

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	db, err := wire.InitDB()
	if err != nil {
		log.Printf("Init Server error:%v\n", err)
		return
	}

	srv, err := wire.InitializeServer(db)

	if err != nil {
		log.Printf("Init Server error:%v\n", err)
		return
	}

	log.Println("Start Server")
	if err = srv.Run(); err != nil {
		log.Printf("Run Server error:%v\n", err)
		return
	}
}
