package models

import (
	"time"

	"go.mongodb.org/mogo-driver/bson/primitive"
)

type User struct{
	ID 							primitive.ObjectID				`jason:"_id"bson:"_id `				
	First_Name 					*string 						`jason:"First_Name" 			validate:"required,min=2,max=30" `
	Last_Name 					*string							`jason:"Last_Name"              validate:"required,min=2,max=30" `
	Password					*string							`jason:"Password"               validate:"required,min=6" `
	Email 						*string							`jason:"Email"                  validate:"email, required"`
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
	Adress_Id					primitive.ObjectID				`bson:"_id" `
	Product_Name 				*string							`json:"product_name"`
	Price 						*uint64							`jason:"price`
	Rating						*uint8							`jason:"rating`
	Image 						*string							`json:"image"`
}
type ProductUser struct{
	Prioduct_ID					primitive.ObjectID				`bson:"_id"`
	Product_Name 				*string							`json:"product_name" bson :"product_name"`
	Price 						*int							`json:"price" bson:"price"`
	Rating						*uint							`json:"rating" bson:"rating"`
	Image 						*string							`json:"image" bson:"image"`

}
type Address struct {
	Address_Id					primitive.ObjectID              `bson:"_id"`
	House                       *string							`json:"house_name" bson:"house_name"`
	Street                      *string							`json:"street_name" bson:"street_name"`		
	City						*string							`json:"city_name" bson:"city_name"`
	Pincode 					*string							`json:"pinc_ode" bson:"pin_code"`

}
type Order struct{
	Order_Id					primitive.ObjectID				`bson:"_id"`
	Order_Cart					[]ProductUser					`json:"order_list" bson:"order_list"`
	Order_At 					time.Time						`json:"ordered_at" bson:"ordered_at"`
	Price						int								`json:"total_price" bson:"total_price"`
	Discount					*int							`json:"discount" bson:"discount"`
	Payment_Method				Payment							`json:"payment_method" bson:"payment`

}
type Payment struct {
	Digital                     bool
	COD                         bool
}
