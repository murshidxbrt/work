package controllers

import (
	"context"
	"log"

	"github.com/murshidxbrt/work/ecommerce-yt/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
var (

	ErrCantFindProduct = errors.New("can't find product")
	ErrCantDecodeProduct = errors.New("can't find product")
	ErrIdIsNotValid = errors.New("this useer is not valid")
	ErrCantUpdateUser = errors.New("cannoot add this product to the cart")
	ErrCantRemoveIteamCart = errors.New("cannoot remove this product from the cart")
	ErrCantGetIteam = errors.New("was unable to get the iteam from the cart")
	ErrCantBuyCartIteam = errors.New("cannot update the Purchase")

)
func AddProductToCart(ctx context.Context, prodCollection, userCollection *mongo.Collection, productID primitive.ObjectID, userID string) error {
	searchfromdb, err := prodCollection.Find(ctx, bson.M{"_id": productID})
	if err != nil {
		log.Println(err)
		return ErrCantFindProduct
	}
	var productCart []models.ProductUser
	err = searchfromdb.All(ctx, &productcart)
	if err != nil {
		log.Println(err)
		return ErrCantDecodeProduct
	}

	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println(err)
		return ErrUserIdIsNotValid
	}

	filter :=bson.D{primitive.E{key: "_id", value: id}}
	update := bson.D{{key: "$push", Value: bson.D{primitive.E{Key:"usercart", Value: bson.D{{Key:"$each", Value:productCart}}}}}}

	_, err := userCollection.UpdateOne(ctx, filter, update)
	if err!=nil {
		return ErrCantUpdateUser
	}
	return nil
}

func RemoveCartIteam(ctx context.Context, prodCollection, userCollection *mongo.Collection, productID primitive.ObjectID, userID string) error {
	id , err:= primitive.ObjectIDFromHex(userID) 
	if err!=nil {
		log.Println(err)
		return ErrUserIdIsNotValid
	}
	filter := bson.D(primitive.E{Key:"_id", Value: id})
	update := bson.M{"$pull": bson.M{"usercart":bson.M{"_id":productID}}}
	_, err = UpdateMany(ctx, filter, update )
	if err!=nil {
		return ErrCantRemoveIteamCart
	}
	return nil
}

func BuyIteamFromcart(ctx context.Context, userCollection *mongo.Collection,   userID string) error{
	// fetch the cart of the user
	// find the cart total 
	// create an order with the iteams
	// added order to the user collection
	// added iteams in the cart to order list
	// empty up the cart


	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println(err)
		return ErrUserIdIsNotValid
	}

	var getcartitems models.User
	var ordercart models.Order

	ordercart.Order_ID = primitive.NewObjectID()
	ordercart.Ordered_At = time.Now()
	ordercart.Order_Cart = make([]models.ProductUser, 0)
	ordercart.Payment_Method.COD = true

	unwind := bson.D{{Key: "$unwind", Value:bson.D{primitive.E{Key:"path", Value"$usercart"}}}}
	grouping := bson.D{{Key:"$group", Value:bson.D{primitive.E{Key:"_id", Value:"$id"}, {Key:"total", Value: bson.D{primitive.E{Key:"$sum", Value:"$usercart.price"}}}}}}
	currentresults, err := userCollectioon.Aggregate(ctx, mongo.Pipeline{unwind , grouping}) 
	ctx.Done()
	if err != nil {
		panic(err)
	}
	var getusercart []bson.M
	if err = currentresult.All(ctx, &getusercart); err != nil {
		panic(err)

	}
	var total_price int32

	for _, user_iteam := range grtusercart{
		price := user_iteam["total"]
		total_price = price.(int32)
	}
	ordercart.Price = int(total_price)

	filter := bson.D{primitive.E{Key:"_id", Value:id}}
	update := bson.D{{Key:"$push", Value:bson.D{primitive.E{Key:"orders", Value:ordercart}}}}
	_, err = userCollection.UpdateMany(ctx, filter, update)
	if err != nil {
		log.Println(err)
	}
	userCollection.FindOne(ctx, bson.D{primitive.E{Key:"_id", Value:id }}.Decode(&getcartitems))
	if err!=nil {
		log.Println(err)
	}

	filter2 := bson.D{primitive.E{Key:"id", Value:id}}
	update2 := bson.M{"$push": bson.M{"orders.$[].order_list:": bson.M{"$each":getcartiteams.UserCart}}}
	_, err = userCollection.UpdateOne(ctx, filter2, update2)
	if err != nil {
		log.Println(err)
	}

	usercart_empty := make([]models.ProductUser, 0)
	filter3 := bson.D{primitive.E{Key:"_id", Value: id}}
	update3 := bson.D{{Key:"$set", Value:bson.D{primitive.E{Key:"usercart", Value: usercart_empty}}}}
	_,err = userCollection.UpdateOne(ctx, filter3, update3)
	if err ! =nil{
		return ErrCantBuyCartItem
	}
	return nil

}

func InstantBuyer(ctx context.Context, prodCollection, userCollection *mongo.Collection, productID primitive.ObjectID, UserID string) error {
	id, err := primitive.ObjectIDFromHex(UserID)

	if err != nil{
		log.Println(err)
		return ErrUserIdNotValid
	}

	var product_details models.ProductUser
	var order_detail models.Order

	orders_details.Order_ID = primitive.NewObjectID()
	orders_details.Ordered_At = time.Now()
	order_detail.Order_cart = make([]models.ProductUser, 0)
	order_detail.Payment_Method.COD = true
	err = prodCollection.FindOne(ctx, bson.D{primitive.E{Key: "_id", Value: productID}}).Decode(&product_details)
	if err != nil {
		log.Println(err)
	}
	orders_detail.Price = product_details.Price

	filter := bson.D{primitive.E{Key:"_id", Value: id}}
	update := bson.D{{Key:"push", Value:bson.D{primitive.E{Key:"orders", Value:orders_details }}}}
	_, err = userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println(err)
	}

	filter2 := bson.D{primitive.E{Key:"_id", Value: id }}
	update2 := bson.M{"$push" :bson.M{"orders.$[].order_list":product_details} }

	_, err = userCollection.UpdateOne(ctx, filter2, update2)
	if err != nil {
		log.Println(err)
	}
	return nil

}

