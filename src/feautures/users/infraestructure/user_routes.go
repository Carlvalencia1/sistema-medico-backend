package infraestructure

import (
	"github.com/gin-gonic/gin"
	"smartvitals/src/feautures/users/infraestructure/controllers"
)

	type UserRoutes struct {
		engine              *gin.Engine
		createUserController *controllers.CreateUserController
		loginController   *controllers.LogInController
		getUsersController  *controllers.GetUsersController 
		updateUserController *controllers.UpdateUserController
		deleteUserController *controllers.DeleteUserController
		getByUsernameController *controllers.GetByUsernameController
		getUserByIdController *controllers.GetUserByIdController
		
	
	}

	func NerUserRoutes (engine *gin.Engine, createUserController *controllers.CreateUserController, loginController *controllers.LogInController, getUsersController *controllers.GetUsersController, updateUserController *controllers.UpdateUserController, deleteUserController *controllers.DeleteUserController, getByUsernameController *controllers.GetByUsernameController, getUserByIdController *controllers.GetUserByIdController) *UserRoutes {
		return &UserRoutes{
			engine: engine,
			createUserController: createUserController,
			loginController: loginController,
			getUsersController: getUsersController,
			updateUserController: updateUserController,
			deleteUserController: deleteUserController,
			getByUsernameController: getByUsernameController,
			getUserByIdController: getUserByIdController,
		}
		
	}

	func (r *UserRoutes) SetupRoutes() {
		UserRoutes := r.engine.Group("/users")
		{
			UserRoutes.POST("/", r.createUserController.Run)
			UserRoutes.POST("/login", r.loginController.Run)
			UserRoutes.GET("/", r.getUsersController.Run)
			UserRoutes.GET("username/:username", r.getByUsernameController.Run)
			UserRoutes.PUT("/:id", r.updateUserController.Run)
			UserRoutes.DELETE("/:id", r.deleteUserController.Run)
			UserRoutes.GET("/:id", r.getUserByIdController.Run)
		}
	}
	func (r *UserRoutes) Run() error {
		if err := r.engine.Run(); err != nil {
            return err
        }
        return nil
    }