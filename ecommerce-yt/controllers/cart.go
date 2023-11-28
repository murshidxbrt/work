package controllers

import (
	"time"
	"context"
	"errors"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	
)

	type Appliction struct {
		prodCollection *mongo.Collection
		userCollection *mongo.Collection
	}
	func NewAppliction(prodCollection , userCollection *mongo.Collection) *Appliction{
		return &Application{
			prodCollection: prodCollection,
			userCollection: userCollection
		}
	}

func (App *Appliction) AddCart() gin.Handler {
	return func (c *gin.Context) {
		productQueryID := c.Query("id")
		if productQueryID == "" {
			log.Println("product id is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))
			return
		}

		userQueryID := c.Query("userID")
		if userQueryID == "" {
			log.Println("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)

		if err != nil {
			log.Println(err)
			c.AbortWithError(http.StatusInternalServerError,)
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), S*time.Second)
		defer cancel()

		err = database.AddProductToCart(ctx, app.prodCollection, app.userCollection, productID, productQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
		c.IndentedJSON(200, "Successfully added to the cart")	
	}

}

func RemoteIteam() gin.HandlerFunc {

}

func GetIteamFromCart() gin.Handler {

}

 func BuyFromcart() gin.HandlerFunc {

}

func instantBuy() gin.HandlerFunc {

}