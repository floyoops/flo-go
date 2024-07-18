package main

import (
	"fmt"
	wiring "github.com/floyoops/flo-go/backend/internal/infra/di"
	"time"
)

func main() {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02T15:04:05Z07:00")

	fmt.Printf("time: %s", formattedTime)
	app, err := wiring.BuildApp()
	if err != nil {
		panic(err)
	}

	if err := app.Start(8080); err != nil {
		panic(err)
	}
}
