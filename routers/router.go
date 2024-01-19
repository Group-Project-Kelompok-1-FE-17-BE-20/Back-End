package router

import (
	"Laptop/middlewares"
	//_projectRepo "Laptop/controllers/project/data"
	//_projectHandler "Laptop/controllers/project/handler"
	//_projectService "Laptop/controllers/project/service"
	//_taskRepo "Laptop/controllers/task/data"
	//_taskHandler "Laptop/controllers/task/handler"
	//_taskService "Laptop/controllers/task/service"
	_userRepo "Laptop/features/user/data"
	_userHandler "Laptop/features/user/handler"
	_userService "Laptop/features/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	userRepo := _userRepo.New(db)
	userService := _userService.New(userRepo)
	userHandlerAPI := _userHandler.New(userService)

	e.POST("/login", userHandlerAPI.Login)
	e.POST("/users", userHandlerAPI.CreateUser)
	e.GET("/users", userHandlerAPI.GetAllUser)
	e.GET("/users/:user_id", userHandlerAPI.GetUserById, middlewares.JWTMiddleware())
	e.PUT("/users/:user_id", userHandlerAPI.UpdateUserById, middlewares.JWTMiddleware())
	e.DELETE("/users/:user_id", userHandlerAPI.DeleteUserById, middlewares.JWTMiddleware())

	// toko
	// e.GET("/users/:user_id/store-attributes", userHandlerAPI.GetTokoById) // get data toko
	// e.GET("/users/:user_id", userHandlerAPI.GetUserById, middlewares.JWTMiddleware())
	// e.PUT("/users/:user_id", userHandlerAPI.UpdateUserById, middlewares.JWTMiddleware())
	// e.DELETE("/users/:user_id", userHandlerAPI.DeleteUserById, middlewares.JWTMiddleware())

	// 	projectRepo := _projectRepo.New(db)
	// 	projectService := _projectService.New(projectRepo)
	// 	projectHandlerAPI := _projectHandler.New(projectService)
	// 	e.POST("/projects", projectHandlerAPI.CreateProject, middlewares.JWTMiddleware())
	// 	e.GET("/projects", projectHandlerAPI.GetAllProject, middlewares.JWTMiddleware())
	// 	e.GET("/projects/:project_id", projectHandlerAPI.GetProjectById, middlewares.JWTMiddleware())
	// 	e.PUT("/projects/:project_id", projectHandlerAPI.UpdateProjectById, middlewares.JWTMiddleware())
	// 	e.DELETE("/projects/:project_id", projectHandlerAPI.DeleteProjectById, middlewares.JWTMiddleware())

	// 	taskRepo := _taskRepo.New(db)
	// 	taskService := _taskService.New(taskRepo, projectRepo)
	// 	taskHandlerAPI := _taskHandler.New(taskService)
	// 	e.POST("/tasks", taskHandlerAPI.CreateTask, middlewares.JWTMiddleware())
	// 	e.GET("/tasks", taskHandlerAPI.GetAllTask, middlewares.JWTMiddleware()) // masih kurang
	// 	e.GET("/tasks/:task_id", taskHandlerAPI.GetTaskById, middlewares.JWTMiddleware())
	// 	e.PUT("/tasks/:task_id", taskHandlerAPI.UpdateTaskById, middlewares.JWTMiddleware())
	// 	e.DELETE("/tasks/:task_id", taskHandlerAPI.DeleteTaskById, middlewares.JWTMiddleware())

}
