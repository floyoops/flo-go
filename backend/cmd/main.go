package main

import (
	"fmt"
	"github.com/floyoops/flo-go/internal"
	"time"
)

func main() {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02T15:04:05Z07:00")

	fmt.Printf("time: %s", formattedTime)
	if err := internal.NewApp().Start(8080); err != nil {
		panic(err)
	}
}
