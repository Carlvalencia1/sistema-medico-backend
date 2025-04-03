package main

import (
	"smartvitals/src/core"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	medicalcasesUseCases "smartvitals/src/feautures/cases/application"
	medicalcasesInfrastructure "smartvitals/src/feautures/cases/infrastructure"
	medicalcasesControllers "smartvitals/src/feautures/cases/infrastructure/controllers"

	patientsUseCases "smartvitals/src/feautures/patients/application"
	patientsInfrastructure "smartvitals/src/feautures/patients/infraestructure"
	patientsControllers "smartvitals/src/feautures/patients/infraestructure/controllers"

	usersUseCases "smartvitals/src/feautures/users/application"
	usersInfraestructure "smartvitals/src/feautures/users/infraestructure"
	usersControllers "smartvitals/src/feautures/users/infraestructure/controllers"

	esp32UseCases "smartvitals/src/feautures/esp32/application"
	esp32Infrastructure "smartvitals/src/feautures/esp32/infraestructure"
	esp32Adapter "smartvitals/src/feautures/esp32/infraestructure/adapters"
	esp32Controllers "smartvitals/src/feautures/esp32/infraestructure/controllers"
)

type Dependencies struct {
	engine *gin.Engine
}

func NewDependencies() *Dependencies {
	return &Dependencies{
		engine: gin.Default(),
	}
}

func (d *Dependencies) Run() error {
	database := core.NewDatabase()
/*
	// Inicializar RabbitMQ (puede devolver nil si hay error)
	var rabbitMQ *core.RabbitMQ
	rabbitMQTmp, err := core.NewRabbitMQ()
	if err != nil {
		log.Printf("Advertencia: No se pudo conectar a RabbitMQ: %v", err)
		// Continuamos con rabbitMQ = nil
	} else {
		rabbitMQ = rabbitMQTmp
		// No uses defer rabbitMQ.Close() aquí, cerrará la conexión
		// antes de que se use
	} */
	// Configuración de dependencias para casos médicos
	medicalCaseDatabase := medicalcasesInfrastructure.NewMySQL(database.Conn)
	getAllMedicalCaseUseCase := medicalcasesUseCases.NewGetAllUseCase(medicalCaseDatabase)
	getAllMedicalCaseController := medicalcasesControllers.NewGetAllController(getAllMedicalCaseUseCase)
	createMedicalCaseUseCase := medicalcasesUseCases.NewCreateMedicalCaseUseCase(medicalCaseDatabase)
	createMedicalCaseController := medicalcasesControllers.NewCreateMedicalCaseController(createMedicalCaseUseCase)
	updateMedicalCaseUseCase := medicalcasesUseCases.NewUpdateMedicalCaseUseCase(medicalCaseDatabase)
	updateMedicalCaseController := medicalcasesControllers.NewUpdateMedicalCaseController(updateMedicalCaseUseCase)
	deleteMedicalCaseUseCase := medicalcasesUseCases.NewDeleteMedicalCaseUseCase(medicalCaseDatabase)
	deleteMedicalCaseController := medicalcasesControllers.NewDeleteMedicalCaseController(deleteMedicalCaseUseCase)
	medicalCasesRoutes := medicalcasesInfrastructure.NewMedicalCaseRoutes(
		d.engine,
		createMedicalCaseController,
		getAllMedicalCaseController,
		updateMedicalCaseController,
		deleteMedicalCaseController,
	)

	/*findActiveCajaByEsp32UseCase := cajasUseCases.NewFindActiveCajaByEsp32UseCase(cajasDatabase)
     // Crear el productor (funcionará incluso con rabbitMQ = nil)
naranjaProducer := naranjasInfrastructure.NewProducer(rabbitMQ)

	naranjaDatabase := naranjasInfrastructure.NewMySQL(database.Conn)
	createNaranjaUseCase := naranjasUseCases.NewCreateNaranjaUseCase(
		naranjaDatabase,
		naranjaProducer,
		findActiveCajaByEsp32UseCase, */
	
	// Configuración de dependencias para pacientes
	patientsDatabase := patientsInfrastructure.NewMySQL(database.Conn)
	getAllPatientsUseCase := patientsUseCases.NewGetAllPatientsUseCase(patientsDatabase)
	getAllPatientsController := patientsControllers.NewGetAllController(getAllPatientsUseCase)
	createPatientsUseCase := patientsUseCases.NewCreatePatientsUseCase(patientsDatabase)
	createPatientsController := patientsControllers.NewCreatePatientsController(createPatientsUseCase)
	updatePatientsUseCase := patientsUseCases.NewUpdatePatientsUseCase(patientsDatabase)
	updatePatientsController := patientsControllers.NewUpdatePatientsController(updatePatientsUseCase)
	deletePatientsUseCase := patientsUseCases.NewDeletePatientsUseCase(patientsDatabase)
	deletePatientsController := patientsControllers.NewDeletePatientsController(deletePatientsUseCase)
	patientsRoutes := patientsInfrastructure.NewPatientsRoutes(
		d.engine,
		createPatientsController,  
		getAllPatientsController,  
		updatePatientsController,
		deletePatientsController,
	)

	//configuracion de dependencias para usuarios
	userDataBase := usersInfraestructure.NewMysql(database.Conn)
	createUser :=  usersUseCases.NewSaveUser(userDataBase)
	logInUser := usersUseCases.NewLogInUseCase(userDataBase)
	createUserController := usersControllers.NewCreateUserController(createUser)
	logInController := usersControllers.NewLoginController(logInUser)
	userUpdate := usersUseCases.NewUpdateUserUseCase(userDataBase)
	updateUserController := usersControllers.NewUpdateUserController(userUpdate)
	deleteUserUseCase := usersUseCases.NewDeleteUserUsecase(userDataBase)
	deleteUserController := usersControllers.NewDeleteUserController(deleteUserUseCase)
	getUsersUseCase := usersUseCases.NewGetUsersUseCase(userDataBase)
	getUsersController := usersControllers.NewGetUsersController(getUsersUseCase)
	getUserByIdUseCase := usersUseCases.NewGetUserByIDUseCase(userDataBase)
	getUserByIdController := usersControllers.NewGetUserByIdController(getUserByIdUseCase)
	getUserByUsernameUseCase := usersUseCases.NewGetUserByUsernameUseCase(userDataBase)
	getUserByUsernameController := usersControllers.NewGetByUsernameController(getUserByUsernameUseCase)
	userRoutes := usersInfraestructure.NewUserRoutes(
		d.engine,
		createUserController,
		logInController,
		getUsersController,        // Este debe ir antes
		updateUserController,
		deleteUserController,
		getUserByUsernameController,
		getUserByIdController,
	)

	esp32Database := esp32Adapter. NewMySQL(database.Conn)
	createEsp32UseCase := esp32UseCases.NewSaveEsp32(esp32Database)
	createEsp32Controller := esp32Controllers.NewCreateEsp32Controller(createEsp32UseCase)
	getEsp32ByUsernameUseCase := esp32UseCases.NewGetEsp32ByOwnerIDUseCase(esp32Database)
	getEsp32ByUsernameController := esp32Controllers.NewGetEsp32ByPropietarioController(getEsp32ByUsernameUseCase)
	deleteEsp32UseCase := esp32UseCases.NewDeleteEsp32UseCase(esp32Database)
	deleteEsp32Controller := esp32Controllers.NewDeleteEsp32Controller(deleteEsp32UseCase)
	sp32Routes := esp32Infrastructure.NewEsp32Routes(d.engine, createEsp32Controller, getEsp32ByUsernameController, deleteEsp32Controller)

	// Configuración de CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	d.engine.Use(cors.New(config))

	// Configuración de rutas
	medicalCasesRoutes.SetupRoutes()
	patientsRoutes.SetupRoutes()
	userRoutes.SetupRoutes()
	sp32Routes.Run()

	return d.engine.Run(":8080")
}
