package main

import (
	"github.com/ANUMADHAV07/crud-go.git/internal/app"
)


func main() {
	app,err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	app.Logger.Printf("we are running our app")
}