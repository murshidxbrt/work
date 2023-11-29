package controllers

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"go.mogodb.org/mongo-driver/bson"
	"go.mogodb.org/mongo-driver/bson/primitive"
	"go.mogodb.org/mongo-driver/mongo"
)

var UserCollection *mongo.Collection =database.UserData(database.Client, "Users")
var ProductCollection *mongo.Collection = database.ProductData(database.Client, "Products")
var Validate = Validator.New()

func HashPassword (password string) string{
	bytes, err := bycrypt.GetIteamFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	} 
	return string(bytes)

}

func VerifyPassword (userPassword string, givenPassword string) (bool, string) {
	err := bycrypt.CompareHashAndPassword([]byte(givenPassword), []byte(userPassword))
	valid := true
	msg := ""

	if err != nil {
		msg = "Login or Password is incorrect"
		valid = false
	}
	return valid, msg
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
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error" : "this phone no. is already in use"})
			return
		}
		password := HashPassword("user.Password")
		user.Password = &password

		user.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User = user.ID.hex()
		token, refreshtocken,  _ := generate.TockenGenerator(*user.Email, *user.First_Name, *user.Last_Name, *user.User_Id)
		user.Token = &token
		user.Refresh_Token = &refreshtocken
		user.UserCart = make([]models.ProductUser, 0)
		user.Address_Details = make([]models.Address, 0)
		user.Order_Status = make([]models.Order, 0)
		_, inserterr := UserCollection.InsertOne(ctx, user)
		 if inserterr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"the user did not get created"})
			return
		 }
		 defer cancel()
		 c.JSON(http.StatusCreated, "Successfully Signed in !")


	}

}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
			var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
			defer cancel()

			var user models.User
			if err := c.BindJSON(&user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H {"error": err})
				return
			}

			err := UserCollection.FindOne(ctx ,bson.M{"email": user.Email}).Decode(&founduser)
			defer cancel()

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "login or password incorrect"})
				return
		
			}
			PasswordIsValid, msg := VerifyPassword(*user.Password, *founduser.Password)

			defer cancel()

			if !PasswordIsValid{
				c.JSON(http.StatusInternalServerError, gin.H{"error":msg})
				fmt.Println(msg)
				return
			}
			token, refreshToken, _ := generate.TockenGenerator(*founduser.Email, *founduser.First_Name, *founduser.Last_Name, *founduser.User_ID)
			defer cancel()

			generate.UpdateAllTockens(token, refreshToken, *founduser.User_ID)
			c.JSON(http.Statusfound, founduser)

		}
	}
	

func ProductViewerAdmin() gin.handlerFunc{

}

func searchProduct() gin.HandlerFunc {

}

func searchProductByQuery() gin.HandlerFunc {
	
}