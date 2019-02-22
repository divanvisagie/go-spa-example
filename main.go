package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/divanvisagie/go-spa-example/statik"
	"github.com/labstack/echo"
	"github.com/rakyll/statik/fs"
	"gopkg.in/go-playground/validator.v8"
)

var e = echo.New()

// Todo is a representation of a todo object for JSON
type Todo struct {
	Description string `json:"description"`
}

// Validate that he object is sufficient
func (t *Todo) Validate() error {
	if t.Description == "" {
		return fmt.Errorf("Item is missing it's description")
	}
	return nil
}

// ErrorJSON returns an error text in json
type ErrorJSON struct {
	Message string `json:"message"`
}

type todoController struct {
	get  func(c echo.Context) error
	post func(c echo.Context) error
}

// CustomValidator is used for validation, dah
type CustomValidator struct {
	validator *validator.Validate
}

func newTodoController() *todoController {
	todos := []*Todo{
		&Todo{
			Description: "Default todo",
		},
	}

	return &todoController{
		get: func(c echo.Context) error {
			e.Logger.Printf("Handle get")
			return c.JSON(http.StatusOK, todos)
		},
		post: func(c echo.Context) error {
			e.Logger.Printf("Handle post")
			t := &Todo{}
			if err := c.Bind(t); err != nil {
				e.Logger.Printf("Could not parse request")
				return c.JSON(http.StatusBadRequest, err)
			}

			if err := t.Validate(); err != nil {
				e.Logger.Printf("Validation failed")
				return c.JSON(http.StatusBadRequest, ErrorJSON{Message: err.Error()})
			}

			todos = append(todos, t)

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

	e.GET("/*", echo.WrapHandler(http.StripPrefix("/", staticHandler)))

	e.GET("/todo", tc.get)
	e.POST("/todo", tc.post)

	e.Logger.Fatal(e.Start(addr))
}
