package main

import (
	"fmt"

	"github.com/fergiee7/Practica06.git/cmd/config"
)

func main() {
	fmt.Println("Variables de entorno: ", config.USERNAME, config.PASSWORD)
}
