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

func (app *Appliction) RemoveIteam() gin.HandlerFunc {

	return func(c *gin.Context) {
		productQueryID := c.Query("id")
		if productQueryID == "" {
			log.Println("product id is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))
			return
		}
		userQueryID := c.Query("user_ID")
		if userQueryID == "" {
			log.Println("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return
	    }
		productID, err := primitive.ObjectIDFromHex(productQueryID) 

		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return 

		}
		
		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

		defer cancel()

		err = database.RemoveCartIteam(ctx, app.prodCollection, app.userCollection, productID, productQueryID)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}
		c.IndentedJSON(200, "Successfully removed iteam from cart")


	}

}

func GetIteamFromCart() gin.Handler{

	return func(c *gin.Context){
		user_id := c.Query("id")

		if user_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error": "Invalid id"})
			c.Abort()
			return
		}

		usert_id, _ := primitive.ObjectIDFromHex(user_id)

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var filledcart models.User
		err := UserCollection.FindOne(ctx, bson.D{primitive.E{key: "_id", Value : usert_id}}).Decode(&filledcart)

		if err != nil {
			log.Println("err")
			c.IndentedJSON(500, "not found")
			return
		}

		filter_match := bson.D{{key:"$match", Value: bson.D{primitive.E{key: "id", Value:usert_id}}}}
		unwind := bson.D{{key:"$unwind", Value: bson.D{primitive.E{key:"path", Value:"$usercart"}}}}
		grouping := bson.D{{key:"$group", Value: bson.D{primitive.E{key:"_id", Value: }}
	
	   
	}

}

 func (app *Appliction) BuyFromcart() gin.HandlerFunc {

	return func(c *gin.Context) {
		userQueryID := c.Query("id")

		if userQueryID == ""{
			log.Panicln("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("UserID is empty"))
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.second)

		defer cancel()

		err := database.BuyIteamFromCart(ctx, app.userCollection ,userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)

		}
		c.IndentedJSON("sucessfully placed the order")
		
	}

}

func (app *Appliction) InstantBuy() gin.HandlerFunc {

	return func (c *gin.Context) {

		productQueryID := c.Query("id")
		if productQueryID == "" {
			log.Println("product id is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))
			return
		}
		userQueryID := c.Query("user_ID")
		if userQueryID == "" {
			log.Println("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return
	    }
		productID, err := primitive.ObjectIDFromHex(productQueryID) 

		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return 

		}
		
		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

		defer cancel()

		err = database.InstantBuyer(ctx, app.prodCollection, app.userCollection, productID, userQueryID)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}

		c.IndentedJSON(200, "sucessfully placed the order")
	}

}