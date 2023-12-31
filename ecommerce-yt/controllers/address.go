package controllers

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddAddress() gin.HandlerFunc{

	return func( c *gin.Context){
		user_id := c.Query("id")
		if user_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error": "Invalid code"})
			c.Abort()
			return
		}
		address, err := ObjectIDFromHex(user_id)
		if err != nil {
			c.IndentedJSON(500, "Internal Server Error")
		}

		var address models.Address

		address.Address_id = primitive.NewObjectID()

		if err = c.BindJSON(&addresses); err != nil {
			c.IndentedJSON(http.StatusNotAcceptable, err.Error())
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		match_filter := bson.D{{key:"$match", Value: bson.D{primitive.E{Key:"_id", Vlue: address}}}}
		unwind := bson.D{{Key:"$unwind", Value:bson.D{primitive.E{key:"path", Value: "$address"}}}}
		groupe := bson.D{{Key:"$group", Value:bson.D{primitive.E{key:"_id", Value:"$address_id"},{key:"count", Value: bson.D{primitive.E{key:"$sum", Value: 1}}}}
		pontcourser, err :=UserCollection.Aggregate(ctx, mongo.Pipeline{match_filter,unwind, group})


	}
}

func EditHomeAddress() gin.HandlerFunc{

}

func EditWorkAddress() gin.HandlerFunc{

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