package routes

import (
	"github.com/WeDias/golang-test-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupAnalyticsRoutes(api fiber.Router) {
	apiProduct := api.Group("/analytics")
	apiProduct.Get("/best-seller", controllers.BestSeller)
}
