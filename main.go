package main

import (
	"UserAuth/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":80", r))
	fmt.Println("Listening at port 4000 ...")
}
