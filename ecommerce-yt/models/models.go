package models

import (
	"time"

	"go.mongodb.org/mogo-driver/bson/primitive"
)

type User struct{
	ID 							primitive.ObjectID				`jason:"_id"bson:"_id `
	First_Name 					*string 						`jason:"First_Name" `			
	Last_Name 					*string							`jason:"Last_Name" `
	Password					*string							`jason:"Password" `
	Email 						*string							`jason:"Email" `
	Phone						*string							`jason:"Phone"' `
	Token						*string							`jason:"Token" `
	Refresh_Token 		        *string							`jason:"Refresh_Token"`
	Created_At					time.Time						`jason:"Created_At"`
	Updated_At		  		 	time.Time 						`jason:"Updated_At"`
	User_ID						*string							`jason:"User_id"`
	UserCart					[]ProductUser					`jason:"user_cart" bson:"usercart"`
	Address_Details				[]Address						`jason:"adress" bson:"address"` 
	Order_Status				[]Order							`jason:"orders"bson:"orders"`
}

type Product struct{
	Adress_Id					primitive.ObjectID
	Product_Name 				*string
	Price 						*uint64
	Rating						*uint8
	Image 						*string
}
type ProductUser struct{
	Prioduct_ID					primitive.ObjectID
	Product_Name 				*string
	Price 						*uint64
	rating						*uint8	
	Image 						*string

}
type Address struct {
	Address_Id					*string
	House                       *string
	Street                      *string	
	City						*string
	Pincode 					*string

}
type Order struct{
	Order_Id					primitive.ObjectID
	Order_Cart					[]ProductUser
	Order_At 					time.Time
	Price						int
	Discount					*int
	Payment_Method				Payment

}
type Payment struct {
	Digital                     bool
	COD                         bool
}