package initialize

import (
	"log"
	"os"
)

func Run() {
	env := os.Getenv("ENV")
	if env == "" {
        log.Fatal("ENV variable is not set")
    }
	LoadConfig(env)

	r := InitRouter()

	r.Run(":8002")
}