package main

import (
	"log"
	"net/http"

	_ "github.com/divanvisagie/go-spa-example/statik"
	"github.com/labstack/echo"
	"github.com/rakyll/statik/fs"
)

// Todo is a representation of a todo object for JSON
type Todo struct {
	Description string `json:"description"`
}

type todoController struct {
	get  func(c echo.Context) error
	post func(c echo.Context) error
}

func newTodoController() *todoController {
	todos := []*Todo{
		&Todo{
			Description: "Default todo",
		},
	}

	return &todoController{
		get: func(c echo.Context) error {
			return c.JSON(http.StatusOK, todos)
		},
		post: func(c echo.Context) error {
			return c.JSON(http.StatusOK, todos)
		},
	}
}

func main() {
	addr := ":8080"

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	staticHandler := http.FileServer(statikFS)
	tc := newTodoController()

	e := echo.New()
	e.GET("/*", echo.WrapHandler(http.StripPrefix("/", staticHandler)))
	e.GET("/todo", tc.get)
	e.POST("/todos", tc.post)

	e.Logger.Fatal(e.Start(addr))
}
