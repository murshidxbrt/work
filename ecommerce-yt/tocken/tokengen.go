package tocken

import(
	"log"
	"os"
	"time"
	"github.com/murshidxbrt/ecommerce-yt/database"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/form3tech-oss/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
)


type SignedDetails struct {
	Email string 
	First_Name string
	Last_Name string
	Uid string
	jwt.StandardClaims

}

var UserDta *mongo.Collectiion = database.userData(database.Client, "Users")

var SECRET_KEY = os.Getenv("SECRET_KEY")

func TockenGenerator(email string, firstName string, lastName string, uid string)(signedtocken string, singnedrefreshtoken string,  err error){
	claims := &SignedDetails{
		Email: email,
		First_Name: firstName,
		Last_Name: lastName,
		Uid: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshclaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},

 	}

	tocken, err :=  jwtNewWithClaims(jwt.SingingMethodHS256, claims).SingnedString([]byte(SECRET_KEY))
		if err != nil {
			return "", "", err
		}

		refreshtoken, err := jwtNewWithClaims(jwt.SingingMethodHS384, refreshclaims).SingnedString([]byte(SECRET_KEY))
		if err != nil {
			log.Panic(err)
		}
		return token, refreshtoken, err

}

func ValidateTocken(signedtocken string)(claims *SignedDetails, msg string){
	token, err := jwt.ParseWithClaims(signedtocken, &SignedDetails{}, func(tocken *jwt.Tocken)(interface{}, error){
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		msg =  err.Error()
		return
	}

	claims, ok := tocken.Claims.(*SignedDetails)
	if !ok {
		msg = "the token in invalid"
		return
	}

	claims.ExpiresAt < time.Now().Add().Local().Unix(){
			msg = "tocken already expired"
			return
	}
	return claims, msg

}

func UpdateAllTockens(signedtocken string, storedfreshtoken string,userid string){
	var ctx, cancel = context.WithTimeout(context.Background(),100*time.Second)

	var updateobj primitive.D

}
