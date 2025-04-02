package infraestructure

import (
	"smartvitals/src/feautures/users/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
    engine                  *gin.Engine
    createUserController    *controllers.CreateUserController
    loginController         *controllers.LogInController
    getUsersController      *controllers.GetUsersController
    updateUserController    *controllers.UpdateUserController
    deleteUserController    *controllers.DeleteUserController
    getByUsernameController *controllers.GetByUsernameController
    getUserByIdController   *controllers.GetUserByIdController
}

func NewUserRoutes(
    engine *gin.Engine,
    createUserController *controllers.CreateUserController,
    loginController *controllers.LogInController,
    getUsersController *controllers.GetUsersController,
    updateUserController *controllers.UpdateUserController,
    deleteUserController *controllers.DeleteUserController,
    getByUsernameController *controllers.GetByUsernameController,
    getUserByIdController *controllers.GetUserByIdController,
) *UserRoutes {
    return &UserRoutes{
        engine:                  engine,
        createUserController:    createUserController,
        loginController:         loginController,
        getUsersController:      getUsersController,
        updateUserController:    updateUserController,
        deleteUserController:    deleteUserController,
        getByUsernameController: getByUsernameController,
        getUserByIdController:   getUserByIdController,
    }
}

func (r *UserRoutes) SetupRoutes() {
    userRoutes := r.engine.Group("/users")
    {
        userRoutes.POST("/", r.createUserController.Run)
        userRoutes.POST("/login", r.loginController.Run)
        userRoutes.GET("/", r.getUsersController.Run)
        userRoutes.GET("/username/:username", r.getByUsernameController.Run)
        userRoutes.PUT("/:id", r.updateUserController.Run)
        userRoutes.DELETE("/:id", r.deleteUserController.Run)
        userRoutes.GET("/:id", r.getUserByIdController.Run)
    }
}

func (r *UserRoutes) Run() error {
    if err := r.engine.Run(); err != nil {
        return err
    }
    return nil
}