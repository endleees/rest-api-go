package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/endleees/rest-api-go/internal/app/apiserver"
	"github.com/endleees/rest-api-go/internal/app/test"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()
	config := apiserver.NewConfig()
	_, error := toml.DecodeFile(configPath, config)
	if error != nil {
		log.Fatal(error)
	}
	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
	test.Test()

}
