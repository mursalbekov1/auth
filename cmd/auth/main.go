package main

import (
	"auth/internal/config"
	"fmt"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println((cfg))

	// TODO: инициалиризовать логгер

	// TODO: инициалиризовать app3
}
