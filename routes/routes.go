package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zanamira43/todo/controllers"
	"github.com/zanamira43/todo/middleware"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/", controllers.Home)

	// auth routes
	api.Post("/register", controllers.Register)
	api.Post("/login", controllers.Login)

	// todos end-points route
	api.Get("/todos", controllers.GetAllTodos)
	api.Post("/todos", controllers.CreateTodo)
	api.Get("/todos/:id", controllers.GetSingleTodo)
	api.Put("/todos/:id", controllers.UpdateTodo)
	api.Delete("/todos/:id", controllers.DeleteTodo)

	// todo date routes end point
	api.Get("/dates", controllers.AllTodoDates)
	api.Post("/dates", controllers.CreateTodoDate)
	api.Get("/dates/:id", controllers.GetSingleTodoDate)
	api.Put("/dates/:id", controllers.UpdateTodoDate)
	api.Delete("/dates/:id", controllers.DeleteTodoDate)
	api.Post("/search/date", controllers.GetSingleTodoDateByDate)

	// middleware ware route
	api.Use(middleware.IsAuthenticated)

	api.Get("/user", controllers.User)
	api.Post("/logout", controllers.Logout)

	// update user info & password route
	api.Put("/users/info", controllers.UpdateInfo)
	api.Put("/users/password", controllers.UpdatePassword)

	// user routes
	api.Get("/users", controllers.AllUsers)
	api.Post("/users", controllers.CreateUser)
	api.Get("/users/:id", controllers.GetUser)
	api.Put("/users/:id", controllers.UpdateUser)
	api.Delete("/users/:id", controllers.DeleteUser)

	// user role routes
	api.Get("/roles", controllers.AllRoles)
	api.Post("/roles", controllers.CreateRole)
	api.Get("/roles/:id", controllers.GetRole)
	api.Put("/roles/:id", controllers.UpdateRole)
	api.Delete("/roles/:id", controllers.DeleteRole)

	// user permission routes
	api.Get("/permissions", controllers.AllPermission)
	api.Post("/permissions", controllers.CreatePermission)

}
