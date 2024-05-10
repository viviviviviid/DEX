package main

import (
	"fmt"

	"github.com/m/go/db"
	"github.com/m/go/rest"
)

func main() {
	// defer db.close() 함수 추가시
	fmt.Println("Start")
	db.InitDB()
	rest.Start()
}
