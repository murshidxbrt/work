package controllers

import(
	
)

func AddAddress() gin.HandlerFunc{

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