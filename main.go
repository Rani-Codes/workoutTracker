package main

import (
	"github.com/Rani-Codes/workoutTracker/internal/app"
)

func main() {
	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	app.Logger.Println("App is running")
}
