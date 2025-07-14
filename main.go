package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/ANUMADHAV07/crud-go.git/internal/app"
	"github.com/ANUMADHAV07/crud-go.git/internal/routes"
)

func main() {
	var port int

	flag.IntVar(&port, "port", 8080, "go backend server running")
	flag.Parse()

	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	defer app.DB.Close()
	r := routes.SetupRoutes(app)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf("we are running our app %d\n", port)

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}

}
