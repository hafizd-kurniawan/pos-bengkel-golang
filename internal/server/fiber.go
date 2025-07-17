package server

import (
	"boilerplate/config"
	"boilerplate/internal/delivery/http/routes"
	"boilerplate/internal/middleware"
	"boilerplate/internal/repository"
	"boilerplate/internal/usecase"
	"boilerplate/internal/wrapper/handler"
	repo_wrapper "boilerplate/internal/wrapper/repository"
	usecase_wrapper "boilerplate/internal/wrapper/usecase"
	"boilerplate/pkg/infra/db"
	"fmt"
	"log"
	"time"
	
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/sirupsen/logrus"
)

func Run(conf *config.Config, dbList *db.DatabaseList, appLoger *logrus.Logger) {

	//* Initial Engine
	engine := html.New("./views", ".html")

	//* Initial Fiber App
	app := fiber.New(fiber.Config{
		AppName:      conf.App.Name,
		ServerHeader: "Go Fiber",
		Views:        engine,
		BodyLimit:    conf.App.BodyLimit * 1024 * 1024,
	})

	//* Initial Data Middleware
	middleware.InitMiddlewareConfig(app, dbList, conf, appLoger)

	//* General Middleware
	middleware.CORSMiddleware()
	middleware.DefaultLimitterMiddleware()
	//middleware.RecoverMiddleware()

	//* Initial New Architecture (Repository -> Usecase -> Handler)
	
	// Initialize new repository manager
	repoManager := repository.NewRepositoryManager(dbList.DatabaseApp)
	
	// Initialize new usecase manager
	usecaseManager := usecase.NewUsecaseManager(repoManager)
	
	// Setup new routes
	routes.SetupFoundationRoutes(app, usecaseManager)
	routes.SetupCustomerRoutes(app, usecaseManager)
	routes.SetupInventoryRoutes(app, usecaseManager)
	routes.SetupServiceRoutes(app, usecaseManager)
	routes.SetupFinancialRoutes(app, usecaseManager)
	
	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"message": "POS Bengkel API is running",
			"timestamp": time.Now(),
		})
	})

	//* Initial Old Wrapper (for compatibility)

	if dbList.DatabaseApp == nil {
		log.Println("is nil")
	}

	repo := repo_wrapper.NewRepository(conf, dbList, appLoger)
	usecase := usecase_wrapper.NewUsecase(repo, conf, dbList, appLoger)
	handler := handler.NewHandler(usecase, conf, appLoger)

	//* Root Endpoint
	app.Get("/", handler.General.Root.GetRoot)

	//* Api Endpoint
	// api := app.Group(conf.App.Endpoint)

	//* General Routes
	//generalEncyrption.NewRoutes(api, handler)

	//* Core Routes
	
	//* CMS Routes
	// cmsWorkOfType.NewRoutes(api, handler)

	//* Not found
	app.All("*", handler.General.NotFound.GetNotFound)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", conf.App.Port)))
}
