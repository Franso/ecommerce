package main

import (
	"log"
	"os"

	"github.com/franso/ecommerce/controllers"
	"github.com/franso/ecommerce/database"
	"github.com/franso/ecommerce/middleware"
	"github.com/franso/ecommerce/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// initialize port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// initialize app to handle all routes
	// get product data from the "Products" collection and user data from the "Users" collection
	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	// create the router
	router := gin.New()

	// use the gin logger for logging
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	// define more routes apart from the user routes
	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}
