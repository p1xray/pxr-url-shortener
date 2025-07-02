package main

import (
	"fmt"
	"github.com/p1xray/pxr-url-shortener/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Printf("%#v\n", cfg)
}
