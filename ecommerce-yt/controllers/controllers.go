package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
t
	"github.com/murshidxbrt/work/ecommerce-yt/database"
	"github.com/murshidxbrt/work/ecommerce-yt/models"
	 generate "github.com/murshidxbrt/work/ecommerce-yt/tockens"
	"github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
	"go.mogodb.org/mongo-driver/bson"
	"go.mogodb.org/mongo-driver/bson/primitive"
	"go.mogodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
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
			var founduser moodels.User
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
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var product models.Product
		defer cancel()
		if err := c.BindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		product.Product_ID = primitive.NewObjectID()
		_,anyerr := ProductCollection.InsertOne(ctx,products)
		if anyerr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"not inserted" })
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, "successfully added")
	}
}

func SearchProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var productlist []models.Product
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		cursor, err := ProductCollection.Find(ctx, bson.D{{}})
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "some thing went wrong, please try after some time")
			return
		}

		err = cursor.All(ctx, &productlist)

		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		defer cursor.Close()

		if err := cursor.err(); err != nil{
			log.Println(err)
			c.IndentedJSON(400, "invalid")
			return
		}
		defer cancel()
		c.IndentedJSON(200, productlist)
	}

}

func SearchProductByQuery() gin.HandlerFunc {
	return func(c *gin.Context){
		var searchProduct []models.Product
		queryParam := c.Query("name ")

		if queryParam == "" {
			log.Println("query is empty")
			c.Header("Content-Type", "application/json")
			c.JSON(HTTP.StatusNotFound, gin.H{"Error": "Invalid search 	index"})
			c.Abort()
			return 

		}

		var ctx, cancel = context.WithTimeOut(context.Background(), 100*time.Second)
		defer cancel()

		searchquerydb, err :=ProductCollection.Find(ctx, bson.M{"product_name" : bson.M{"$regex" : queryParam}})

		if err != nil {
			c.IndentedJSON(404, "something went wrong while fetching the data")
			return
		}

		err = searchquerydb.All(ctx, &searchProduct)
		if err != nil {
			log.Println(err)
			c.IndentJSON(400, "invalid")
			return
		}
		defer searchquerydb.Close(ctx)

		if err = searchquerydb.Err(); err != nil {
			log.Println(err)
			c.IndentedJSON(400, "invalid request")
			return
		}

		defer cancel()
		c.IndentedJSON(200, searchProduct)
	}
	
}