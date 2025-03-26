package routes

import (
	"github.com/fitnis/patient-service/configs"
	"github.com/fitnis/patient-service/internal/handlers"
	middleware "github.com/fitnis/patient-service/internal/middleware"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

// SetupRoutes configures all the routes for the application
func SetupRoutes(e *echo.Echo, config *configs.Config) {
	// Create JWT config
	jwtConfig := middleware.NewJWTConfig(config)

	// Create handlers
	patientHandler := handlers.NewPatientHandler()

	// Middleware
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())
	e.Use(echomiddleware.CORS())

	// Public routes
	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	// API routes
	api := e.Group("/api/v1")

	// Patient routes
	patients := api.Group("/patients")
	patients.POST("", patientHandler.RegisterPatient, middleware.JWTMiddleware(jwtConfig), middleware.HasRole("Doctor"))
	patients.GET("", patientHandler.GetPatients, middleware.JWTMiddleware(jwtConfig), middleware.HasRole("Doctor"))
	patients.GET("/:id", patientHandler.GetPatient, middleware.JWTMiddleware(jwtConfig))
	patients.PUT("/:id", patientHandler.UpdatePatient, middleware.JWTMiddleware(jwtConfig), middleware.HasRole("Doctor"))
	patients.DELETE("/:id", patientHandler.DeletePatient, middleware.JWTMiddleware(jwtConfig), middleware.HasRole("Doctor"))
}
