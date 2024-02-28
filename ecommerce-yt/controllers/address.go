package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/murshidxbrt/ecommerce-yt/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Query("id")
		if userID == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error": "Invalid code"})
			c.Abort()
			return
		}

		address, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			c.IndentedJSON(500, "Internal Server Error")
			return
		}

		var address models.Address

		address.AddressID = primitive.NewObjectID()

		if err := c.BindJSON(&address); err != nil {
			c.IndentedJSON(http.StatusNotAcceptable, err.Error())
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		matchFilter := bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: address}}}}
		unwind := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$address"}}}}
		group := bson.D{{Key: "$group", Value: bson.D{{Key: "_id", Value: "$address_id"}, {Key: "count", Value: bson.D{{Key: "$sum", Value: 1}}}}}}

		pointCursor, err := UserCollection.Aggregate(ctx, mongo.Pipeline{matchFilter, unwind, group})
		if err != nil {
			c.IndentedJSON(500, "Internal server error")
			return
		}

		var addressInfo []bson.M
		if err := pointCursor.All(ctx, &addressInfo); err != nil {
			panic(err)
		}

		var size int32
		for _, addressNo := range addressInfo {
			count := addressNo["count"]
			size = count.(int32)
		}

		if size < 2 {
			filter := bson.D{primitive.E{Key: "_id", Value: address}}
			update := bson.D{{Key: "$push", Value: bson.D{{Key: "address", Value: address}}}}
			_, err := UserCollection.UpdateOne(ctx, filter, update)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			c.IndentedJSON(400, "Not Allowed")
		}
	}
}

	

func EditHomeAddress() gin.HandlerFunc{
	return func(c *gin.Context) {
		user_id := c.Query("id")
		if user_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"Error":"Invalid"})
			c.Abort()
			return
		}
		usert_id, err := primitive.ObjectIDFromhex(user_id)
		if err != nil {
			c.IndentedJSON(500, "Internal server error")
		}
		var editaddress models.Address
		if err := c.BindJSON(&editaddress); err != nil{
			c.IndentedJSON(http.StatusBadRequest, err.Error())
		}
		var ctx, cancel = context.WithTimeout(context.Background(),100 * time.Second)
		defer cancel()
		filter := bson.D{primitive.E{key:"_id", Value: usert_id}}
		update := bson.D{{Key:"$set", Value:bson.D{primitive.E{Key:"address.0.house_name", Value: editaddress.House},{Key:"address.0.street_name", Value: editaddress.Street},{Key:"address.0.city_name", Value: Editaddress.City},{Key:"address.0.pin_code", Value: editadress.Pincode}}}}
		_, err = UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.IndentedJSON(500, "Something went wrong")
		}
		defer cancel()
		ctx.Done()
		c.IndentedJSON(200,"successfully updated the home address")
	}

}

func EditWorkAddress() gin.HandlerFunc{
	return func(c *gin.Context) {
		user_id := c.Query("id")
		if user_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"Error":"Invalid"})
			c.Abort()
			return
		}
		usert_id, err := primitive.ObjectIDFromhex(user_id)
		if err != nil {
			c.IndentedJSON(500, "Internal server error")
		}
		var editaddress models.Address
		if err := c.BindJSON(&editaddress); err != nil{
			c.IndentedJSON(http.StatusBadRequest, err.Error())
		}
		var ctx, cancel = context.WithTimeout(context.Background(),100 * time.Second)
		defer cancel()
		filter := bson.D{primitive.E{key:"_id", Value: usert_id}}
		update := bson.D{{Key:"$set", Value:bson.D{primitive.E{Key:"address.1.house_name", Value: editaddress.House},{Key:"address.1.street_name", Value: editaddress.Street},{Key:"address.1.city_name", Value: Editaddress.City},{Key:"address.1.pin_code", Value: editadress.Pincode}}}}
		_,err = UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.IndentedJSON(500, "something went wrong")
			return
		}
		defer cancel()
		ctx.Done()
		c.IndentedJSON(200, "Sucessfully updated the work address")

	}

}

func DeleteAddress() gin.HandlerFunc{
	return func (c *gin.Context) {
		user_id := c.Query("id")

		if user_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound,  gin.H{"Error":"Invalid Search Index"})
			c.Abort()
			return
		}

		addresses := make (make[]models.Address, 0)
		usert_id, err := primitive.ObjectIDFromhex(user_id)
		if err != nil {
			c.IndentedJSON(500, "Internal server error")
		}
		var ctx, cancel = context.WithTimeout(context.Background(),100 * time.Second)
		defer cancel()
		filter := bson.D{primitive.E(key:"_id", Value: user_id)}
		update := bson.D{{key:"$set", value: bson.D{primitive.E{key:"address", Value: AddAddress}}}}
		_, err = UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.IndentedJSON(400, "wrong command")
			return
		}
		defer func()
		ctx.Done()
		c.IndentedJSON(200, "Sucessfully Deleted")
	}

}