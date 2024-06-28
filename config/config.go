package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

var (
	confPath string
	conf     Cfg
)

type Cfg struct {
	Port       int    `yaml:"port" port-default:"8080"`
	DBSocket   string `yaml:"db_socket" db_host-default:"localhost:8123"`
	DBName     string `yaml:"db_name"`
	DBUser     string `yaml:"db_user"`
	DBPassword string `yaml:"db_password"`
}

func init() {
	flag.StringVar(&confPath, "config", "", "path to config")
	flag.Parse()

	if confPath == "" {
		confPath = os.Getenv("CONFIG_PATH")
	}

	if confPath == "" {
		confPath = "./config/base.yaml"
	}
}

// Load config
func Load() Cfg {
	if _, err := os.Stat(confPath); os.IsNotExist(err) {
		log.Panicf("config path is empty: %s", confPath)
	}

	if err := cleanenv.ReadConfig(confPath, &conf); err != nil {
		log.Panicf("unexpected err: %v", err)
	}

	return conf
}
