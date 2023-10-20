package envs

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvs() {
	err := godotenv.Load("./envs/conf.env")
	if err != nil {
		log.Fatal("Error, load envs: ", err)
	}
}
