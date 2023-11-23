package models

import(
	"time"
	"go.mongodb.org/mogo-driver/bson/primitive"
)

type User struct{
	ID 
	First_Name 					primitive.ObjectID
	Last_Name 					*string
	Password					*string
	Email 						*string
	Phone						*string
	Token						*string
	Refresh_Token 		        *string
	Created_At					time.Time
	Updated_At		  		 	time.Time 
	User_ID						*string
	UserCart					[]ProductUser
	Address_Details				[]Address
	Order_Status				[]Order
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