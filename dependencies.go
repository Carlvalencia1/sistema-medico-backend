package main

import (
	"log"
	"smartvitals/src/core"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

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
	engine := gin.Default()
	engine.SetTrustedProxies([]string{"127.0.0.1"}) // Añadido por seguridad
	return &Dependencies{
		engine: engine,
	}
}

func (d *Dependencies) Run() error {
	// Añadido para cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error cargando archivo .env: ", err)
	}

	database := core.NewDatabase()

	// --- Actualización para RabbitMQ ---
	rabbitMQ, err := core.NewRabbitMQ()
	if err != nil {
		log.Fatalf("❌ Error crítico con RabbitMQ: %v", err) // Cambiado de log.Printf a log.Fatalf
	}
	defer rabbitMQ.Close() // Asegura que se cierre la conexión al finalizar

	// --- CASOS MÉDICOS (Mantenido exactamente igual) ---
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

	// --- Configuración de dependencias para pacientes (Mantenido igual) ---
	patientsDatabase := patientsInfrastructure.NewMySQL(database.Conn)
	patientsProducer := patientsInfrastructure.NewProducer(rabbitMQ)

	getAllPatientsUseCase := patientsUseCases.NewGetAllPatientsUseCase(patientsDatabase)
	getAllPatientsController := patientsControllers.NewGetAllController(getAllPatientsUseCase)

	createPatientsUseCase := patientsUseCases.NewCreatePatientsUseCase(
		patientsDatabase,
		patientsProducer,
	)
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

	// --- USUARIOS (Mantenido exactamente igual) ---
	userDataBase := usersInfraestructure.NewMysql(database.Conn)
	createUser := usersUseCases.NewSaveUser(userDataBase)
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
		getUsersController,
		updateUserController,
		deleteUserController,
		getUserByUsernameController,
		getUserByIdController,
	)

	// --- ESP32 (Mantenido exactamente igual) ---
	esp32Database := esp32Adapter.NewMySQL(database.Conn)
	createEsp32UseCase := esp32UseCases.NewSaveEsp32(esp32Database)
	createEsp32Controller := esp32Controllers.NewCreateEsp32Controller(createEsp32UseCase)
	getEsp32ByUsernameUseCase := esp32UseCases.NewGetEsp32ByOwnerIDUseCase(esp32Database)
	getEsp32ByUsernameController := esp32Controllers.NewGetEsp32ByPropietarioController(getEsp32ByUsernameUseCase)
	deleteEsp32UseCase := esp32UseCases.NewDeleteEsp32UseCase(esp32Database)
	deleteEsp32Controller := esp32Controllers.NewDeleteEsp32Controller(deleteEsp32UseCase)
	sp32Routes := esp32Infrastructure.NewEsp32Routes(d.engine, createEsp32Controller, getEsp32ByUsernameController, deleteEsp32Controller)

	// --- CORS Configuration ---
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://54.84.210.136",      // Your frontend origin
		"http://100.28.173.85:8080", // Your backend origin
	}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"}
	config.AllowHeaders = []string{
		"Origin",
		"Content-Length",
		"Content-Type",
		"Authorization",
		"Access-Control-Allow-Origin",
		"Access-Control-Allow-Headers",
		"Access-Control-Allow-Methods",
		"Accept",
		"X-Requested-With",
	}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true // Enable credentials
	config.MaxAge = 12 * time.Hour

	// Apply CORS middleware
	d.engine.Use(cors.New(config))

	// Add OPTIONS handler for preflight requests
	d.engine.OPTIONS("/*path", func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Status(204)
	})

	// --- RUTAS (Mantenido exactamente igual) ---
	medicalCasesRoutes.SetupRoutes()
	patientsRoutes.SetupRoutes()
	userRoutes.SetupRoutes()
	sp32Routes.Run()

	return d.engine.Run(":8080")
}
