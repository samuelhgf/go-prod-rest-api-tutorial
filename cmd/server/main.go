package main

import (
	"fmt"
	transportHTTP "github.com/samuelhgf/go-prod-rest-api-tutorial/internal/transport/http"
	"net/http"
)

// App - the struct which contains things like pointers
// to database connections
type App struct {}

// Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting Up Our APP")

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}
