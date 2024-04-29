package main

import "github.com/floyoops/flo-go/pkg/app"

func main() {
	if err := app.NewApp().Start(8080); err != nil {
		panic(err)
	}
}
