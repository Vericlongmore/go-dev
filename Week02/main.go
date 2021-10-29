package main

import (
	"go-dev/Week02/service"
	"log"
	"net/http"
)

// TODO: 需要Wrap上抛 因为sql.ErrNoRows是kit基础库抛出的err 所以需要上抛
func main() {
	http.HandleFunc("/auth", service.Auth)
	log.Fatalln(http.ListenAndServe(":8083", nil))
}
