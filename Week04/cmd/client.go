package main

import (
	"context"
	pb "dev/wire/api/article/v1"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	conn, err := grpc.Dial(viper.GetString("grpc.port"), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("grpc not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewArticleClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetArticle(ctx, &pb.ArticleRequest{Id: 2})
	if err != nil {
		log.Printf("error: %v\n", err)
		return
	}
	log.Printf("article: %d:%s", r.GetId(), r.GetTitle())
}
