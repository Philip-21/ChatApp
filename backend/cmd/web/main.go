package main

import (
	"chatapp/internal/config"
	"log"
	"os"
)

func main() {

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := config.AppConfig{
		InfoLog:  infoLog,
		ErrorLog: errorLog,
	}

}
