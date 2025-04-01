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
		
	
	}

	func NerUserRoutes (engine *gin.Engine, createUserController *controllers.CreateUserController, loginController *controllers.LogInController, getUsersController *controllers.GetUsersController, updateUserController *controllers.UpdateUserController, deleteUserController *controllers.DeleteUserController, getByUsernameController *controllers.GetByUsernameController) *UserRoutes {
		return &UserRoutes{
			engine: engine,
			createUserController: createUserController,
			loginController: loginController,
			getUsersController: getUsersController,
			updateUserController: updateUserController,
			deleteUserController: deleteUserController,
			getByUsernameController: getByUsernameController,
		}
		
	}