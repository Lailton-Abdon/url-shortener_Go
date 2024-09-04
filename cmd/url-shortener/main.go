package main

import (
	"fmt"

	"github.com/Lailton-Abdon/url-shortener_Go/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)
}
