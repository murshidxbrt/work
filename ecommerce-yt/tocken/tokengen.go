package tocken

import(
	jwt "github.com/dgrijalva/jwt-go"
)


type SignedDetails struct {
	Email string 
	First_Name string
	Last_Name string
	Uid string
	jwt.StandardClaims

}
var SECRET_KEY

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

	tocken, :=  jwtNewWithClaims(jwt.SingingMethodHS256, claims).SingnedString([]byte(SECRET_KEY))

}

func ValidateTocken(){

}

func UpdateAllTockens(){

}
