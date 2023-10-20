package main

import (
	"users-services/envs"
	"users-services/internal/app"
)

func main() {
	envs.LoadEnvs()
	app.Run()
}
