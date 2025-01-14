package config

import (
	"github.com/joho/godotenv"
	"os"
)

var (
	SessionUserKey = "userId"
	Port           string
	Host           string
	PublicUrl      string
	TelegramUrl    string
	SessionSecret  string
	MongoDBUri     string
)

func init() {
	var err error

	if os.Getenv("GO_ENV") != "production" {
		if err = godotenv.Load(); err != nil {
			panic(err)
		}
	}

	Port = os.Getenv("PORT")
	Host = os.Getenv("HOST")
	PublicUrl = os.Getenv("PUBLIC_URL")
	TelegramUrl = os.Getenv("TELEGRAM_URL")
	SessionSecret = os.Getenv("SESSION_SECRET")
	MongoDBUri = os.Getenv("MONGODB_URI")

	if len(Port) == 0 || len(Host) == 0 || len(PublicUrl) == 0 ||
		len(TelegramUrl) == 0 || len(SessionSecret) == 0 || len(MongoDBUri) == 0 {
		panic("Invalid env variables")
	}
}
