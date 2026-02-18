package main

import (
	"fmt"
	"log"
	"net/http"

	"api/config"
	"api/src/router"
)

func main() {

	config.Carregar()

	fmt.Printf("RODANDO A API")

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":5001", r))
}
