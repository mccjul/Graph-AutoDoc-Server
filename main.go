package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/mccjul/Graph-AutoDoc-Server/handler"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.ERROR)
	e.Use(middleware.Logger())

	// Database connection
	db, err := mgo.Dial("localhost:32768")
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Initialize handler
	h := &handler.Handler{DB: db}

	// APP Routes
	e.POST("/app", h.CreateApp)
	e.GET("/app:id", h.GetApps)
	e.PATCH("/app:id", h.PatchApp)
	e.DELETE("/app:id", h.RemoveApp)

	// DOC Routes
	e.GET("/doc:id", h.GetDoc)
	e.PATCH("/doc", h.PatchDoc)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
