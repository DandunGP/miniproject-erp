package main

import (
	"erp/config"
	m "erp/middleware"
	"erp/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	m.LogMiddleware(e)
	e.Logger.Fatal(e.Start(":8000"))
}
