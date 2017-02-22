package main

import (
	"log"
	"net/http"

	"github.com/nazarnovak/mind2/app"
)

func main() {
	app := &app.App{}
	err := app.Init()
	if err != nil {
		log.Fatalln(err)
	}

	defer app.DB.Close()
	defer app.Glue.Release()

	http.Handle("/", app.Server())

	log.Printf("Listening on :%s", app.Config.Port)
	err = http.ListenAndServe(":" + app.Config.Port, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
