package main

import (
	"chatapp/internal/config"
	"log"
	"net/http"
	"os"
)

const portNumber = "8080"

func main() {

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := config.AppConfig{
		InfoLog:  infoLog,
		ErrorLog: errorLog,
	}
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)

	}

}
