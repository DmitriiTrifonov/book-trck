package main

import (
	"github.com/DmitriiTrifonov/book-trck/internal/config"
	"github.com/kelseyhightower/envconfig"
	"log"
)

func main() {
	var cfg config.Config
	err := envconfig.Process("book_trck", &cfg)
	if err != nil {
		panic(err)
	}
	log.Println(cfg)
}
