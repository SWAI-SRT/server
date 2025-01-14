package main

import (
	"log"

	"swai/config"
	"swai/controller"
	"swai/middleware"
	"swai/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// @title Swagger API Documentation
// @version 1.0
// @description This is a sample server for a Fiber application.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("구성을 로드하지 못했습니다. %v", err)
    }

    db, err := config.InitDB(&cfg)
    if err != nil {
        log.Fatalf("데이터베이스 초기화 실패: %v", err)
    }

    app := fiber.New()

    app.Use(logger.New())

    authService := service.NewAuthService(db, cfg.JWTSecret)
    imageService := service.NewImageService(&cfg)
    authController := controller.NewAuthController(authService, imageService)

    mapService := service.NewMapService(db)
    mapController := controller.NewMapController(mapService)

    reportsService := service.NewReportsService(db)
    reportsController := controller.NewReportsController(reportsService, imageService, mapService)

    imageController := controller.NewImageController(imageService)

    app.Get("/health", func(c *fiber.Ctx) error {
        return c.SendString("OK")
    })

    api := app.Group("/auth")
    api.Post("/signup", authController.Signup)
    api.Post("/signin", authController.Signin)
    api.Get("/refresh", authController.Refresh)

    api.Use(middleware.JWTMiddleware(cfg.JWTSecret))
    api.Get("/me", authController.GetProfile)
    api.Patch("/me", authController.EditProfile)
    api.Post("/logout", authController.Logout)
    api.Delete("/me", authController.DeleteAccount)

    reports := app.Group("/reports")
    reports.Use(middleware.JWTMiddleware(cfg.JWTSecret))
    reports.Post("/", reportsController.CreateReport)
    reports.Get("/", reportsController.FindAllReports)
    reports.Get("/by-user", reportsController.FindReportByUserId)
    reports.Get("/:reportId", reportsController.FindReport)

    mapGroup := app.Group("/map")
    mapGroup.Use(middleware.JWTMiddleware(cfg.JWTSecret))
    mapGroup.Post("/", mapController.CreateMarker)
    mapGroup.Get("/", mapController.FindAllMarker)
    mapGroup.Get("/:markerId", mapController.FindMarker)

    image := app.Group("/image")
    image.Use(middleware.JWTMiddleware(cfg.JWTSecret))
    image.Post("/", imageController.UploadImage)

    log.Fatal(app.Listen(":8080"))
}