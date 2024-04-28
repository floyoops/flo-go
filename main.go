package main

import (
	"fmt"
	"github.com/floyoops/flo-go/pkg/app/render"
	"github.com/floyoops/flo-go/pkg/app/router"
	"github.com/labstack/echo/v4"
	"html/template"
	"os"
)

func main() {
	rootPath, _ := os.Getwd()
	fmt.Printf("rootPath by os.Getwd(): %s", rootPath)
	e := echo.New()
	renderer := &render.TemplateRenderer{
		Templates: template.Must(template.ParseGlob(rootPath + "/public/*.html")),
	}
	e.Renderer = renderer
	router.Init(e)
	if err := e.Start(":8080"); err != nil {
		fmt.Printf("Error on start server %s", err.Error())
	}
}
