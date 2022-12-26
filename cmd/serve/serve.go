package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// e.File("", "./build_wasm/wasm_run.html")
	e.Static("", "./build_wasm")
	e.Static("imgs", "./imgs")
	e.Logger.Fatal(e.Start("127.0.0.1:8080"))
}
