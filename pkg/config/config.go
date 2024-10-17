package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	MongoURI    string
	MongoDB     string
	KafkaBroker string
	KafkaTopic  string
	Port        string
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error cargando archivo .env")
	}

	MongoURI = getEnv("MONGO_URI")
	KafkaBroker = getEnv("KAFKA_BROKER")
	MongoDB = getEnv("MONGO_DB")
	KafkaTopic = getEnv("KAFKA_TOPIC")
	Port = getEnv("PORT")
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Variable de entorno %s no definida", key)
	}
	return value
}
