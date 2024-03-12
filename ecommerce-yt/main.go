package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/murshidxbrt/work/ecommerce-yt/controllers"
	"github.com/murshidxbrt/work/ecommerce-yt/database"
	"github.com/murshidxbrt/work/ecommerce-yt/middleware"
	"github.com/murshidxbrt/work/ecommerce-yt/routes"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := controllers.NewAppliction(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addcart", app.AddToCart())
	router.GET("/removeiteam", app.RemoveIteam())
	router.GET("/cartcheckout", app.ByFromCart())
	router.GET("/instantrbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))

}
