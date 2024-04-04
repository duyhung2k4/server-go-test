package config

import "log"

func init() {
	log.Println("init")
	loadEnv()
}
