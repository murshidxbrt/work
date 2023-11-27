package controllers

import (
	"context"
	"log"
	"net/http"
	"time"
)

func HashPassword (password string) string{

}

func VerifyPassword (userPassword string, givenPassword string) (bool, string) {

}

func Signup () gin.HandlerFunc {

	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON{http.StatusBadRequest, gin.H{"error": err.Error()}}
			return
		}
		validationErr := validation.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
			return
		}
		count, err := UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		if count > 0{
			c.JSON(http.StatusBadRequest, gin.H{"error":" user already exists"})
		}
		count, err  = UserCollection.CountDocuments(ctx, bson.M{"Phone": user.Phone})

		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

	}

}

func Login() gin.HandlerFunc {
	
}

func ProductViewerAdmin() gin.handlerFunc{

}

func searchProduct() gin.HandlerFunc {

}

func searchProductByQuery() gin.HandlerFunc {
	
}