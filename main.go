package main

import (
	"fmt"

	// "gin-temp-app/database"
	"gin-temp-app/database"
	"gin-temp-app/router"

	// "gin-temp-app/router"

	// "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// init
func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("env not found")
		return
	}
	// client, _, err := database.ConnectDb()
	// if err != nil {
	// 	panic("db is not running")
	// }
	// database.Client = client
}




func main() {

    router.Routing()

	// r.router.ServeHTTP(w,r)

}


func init(){
	database.ConnectDb()

}