package main

import (
	"my-go-app/pkg/repository/hello"
	"fmt"
	"my-go-app/routes"
)

func main() {
	t := hello.Hello()
	fmt.Println(t)
	routes.Routes()
}


