package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/redhatua/tasker/database"
	"github.com/redhatua/tasker/handlers"
)

// TODO: Move to config
const databasePath = "./var/database/tasker.db"

func main() {
	e := echo.New()
	db := database.New(databasePath)
	// Assets
	e.File("/bootstrap.min.css", "public/assets/css/bootstrap.min.css")
	e.File("/jquery.min.js", "public/assets/js/jquery.min.js")
	e.File("/bootstrap.min.js", "public/assets/js/bootstrap.min.js")
	e.File("/vue.js", "public/assets/js/vue.js")
	e.File("/vue-resource.min.js", "public/assets/js/vue-resource.min.js")
	e.File("/", "public/index.html")
	// App
	e.GET("/tasks", handlers.GetTasks(db))
	e.PUT("/tasks", handlers.PutTask(db))
	e.DELETE("/tasks/:id", handlers.DeleteTask(db))
	// TODO: Move to config
	err := e.Start("127.0.0.1:8080")
	if err != nil {
		log.Errorf("Error: %s", err)
	}
}
