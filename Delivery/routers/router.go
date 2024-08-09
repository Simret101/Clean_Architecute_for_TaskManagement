package routers

import (
	"task/Delivery/controllers"
	"task/Infrastructure"
	"task/Repositories"
	"task/Usecases"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// To Initialize services, repositories, use cases, and controllers
	taskRepo := Repositories.NewTaskRepository()
	userRepo := Repositories.NewUserRepository()

	jwtService := Infrastructure.NewJWTService("your-secret-key", 24*time.Hour)
	passwordService := Infrastructure.NewPasswordService()

	taskUseCase := Usecases.TaskUseCase{TaskRepo: taskRepo}
	userUseCase := Usecases.UserUseCase{
		UserRepo:        userRepo,
		JWTService:      jwtService,
		PasswordService: passwordService,
	}
	taskController := controllers.TaskController{TaskUseCase: taskUseCase}
	userController := controllers.UserController{UserUseCase: userUseCase}

	// Public routes
	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)

	// Protected routes
	protectedRoutes := r.Group("/tasks")
	protectedRoutes.Use(Infrastructure.AuthMiddleware(jwtService))
	{
		protectedRoutes.GET("/", taskController.GetAllTasks)
		protectedRoutes.GET("/:id", taskController.GetTaskByID)
		protectedRoutes.POST("/", taskController.CreateTask)
		protectedRoutes.PUT("/:id", taskController.UpdateTask)
		protectedRoutes.DELETE("/:id", taskController.DeleteTask)

	}

	return r
}
