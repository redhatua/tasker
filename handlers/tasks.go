package handlers

import (
	"database/sql"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/redhatua/tasker/models"
	"net/http"
	"strconv"
)

type Handler map[string]interface{}

func GetTasks(db *sql.DB) echo.HandlerFunc {

	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetTasks(db))
	}

}

func PutTask(db *sql.DB) echo.HandlerFunc {

	return func(c echo.Context) error {
		var task models.Task
		err := c.Bind(&task)
		if err != nil {
			log.Fatalf("Error binding: ", err)
		}
		id, err2 := models.PutTask(db, task.Name)
		if err2 != nil {
			return err2
		} else {
			return c.JSON(http.StatusCreated, Handler{
				"created": id,
			})
		}
	}

}

func DeleteTask(db *sql.DB) echo.HandlerFunc {

	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		_, err := models.DeleteTask(db, id)
		if err != nil {
			log.Fatalf("Error deleting: ", err)
		} else {
			return c.JSON(http.StatusOK, Handler{
				"deleted": id,
			})
		}

		return c.JSON(http.StatusOK, Handler{
			"deleted": 123,
		})
	}

}
