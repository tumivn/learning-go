package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
)

func main() {

	e := echo.New()

	e.GET("/", hello)
	e.GET("/users/:id", getUser)

	e.GET("/todos", getTodos)

	e.Logger.Fatal(e.Start(":8080"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func getTodos(c echo.Context) error {
	t := getTodoManager()
	t.m.Lock()
	defer t.m.Unlock()
	return c.JSON(http.StatusOK, t.todos)
}

type TodoManager struct {
	todos []Todo
	m     sync.Mutex
}

// Singleton
var todoManager *TodoManager

// get TodoManager Singleton
func getTodoManager() TodoManager {
	if todoManager == nil {
		todoManager = &TodoManager{
			todos: []Todo{
				{
					ID:          "1",
					Title:       "Todo 1",
					Description: "Description 1",
					Done:        false,
				},
			},
			m: sync.Mutex{},
		}
	}
	return *todoManager
}

type Todo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
