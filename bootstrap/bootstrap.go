package bootstrap

import (
	"github.com/joho/godotenv"
)

func Run() {
	_ = godotenv.Load()
}
