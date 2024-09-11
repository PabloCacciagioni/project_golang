package routes

import (
	"time"

	"github.com/PabloCacciagioni/project_golang/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var (
	store   *session.Store
	USER_ID string = "user_id"
)

func SetupRoutes(app *fiber.App) {
	store = session.New(session.Config{
		CookieHTTPOnly: true,
		Expiration:     time.Hour * 1,
	})
	db := database.ConnectDb()

	// Here we setup our database in the context
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	// Then we just get the routes
	todoApp := app.Group("/todo", AuthMiddleware)
	todoApp.Get("/list", ListTodos)
	todoApp.Get("/create", AddTodo)
	todoApp.Post("/create", AddTodo)
	todoApp.Get("/edit/:id", UpdateTodo)
	todoApp.Post("/edit/:id", UpdateTodo)
	todoApp.Delete("/delete/:id", DeleteTodo)
	todoApp.Post("/logout", HandleLogout)

	restApp := app.Group("/", MaybeAuthMiddleware)
	restApp.Get("/", HandleViewHome)
	restApp.Get("/login", HandleViewLogin)
	restApp.Post("/login", HandleViewLogin)
	restApp.Get("/register", HandleViewRegister)
	restApp.Post("/register", HandleViewRegister)
	restApp.Get("/status", GetStatus)
	restApp.Get("/todos/:id", GetTodo)

}
