package main

import (
	"e-commerce/config"
)

func main(){
	server := config.NewServer()
	server.Run()
}
