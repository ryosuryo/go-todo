package main

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/ichtrojan/go-todo/routes"
	"github.com/ichtrojan/thoth"
	"github.com/joho/godotenv"
)

func main() {
	/* Variable untuk logging errors
	 * untuk file log akan disimpan di logs/error.log
	 */
	logger, _ := thoth.Init("log")

	if err := godotenv.Load(); err != nil {
		logger.Log(errors.New("no .env file found"))
		log.Fatal("No .env file found")
	}

	port, exist := os.LookupEnv("PORT")

	if !exist {
		logger.Log(errors.New("PORT not set in .env"))
		log.Fatal("PORT not set in .env")
	}

	// Fungsi ini untuk membuka http server dengan port yang ada di file .env
	err := http.ListenAndServe(":"+port, routes.Init())

	if err != nil {
		logger.Log(err)
		log.Fatal(err)
	}
}
