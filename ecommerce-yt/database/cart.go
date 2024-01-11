package controllers

import (
	"context"

	"github.com/murshidxbrt/ecommerce-yt/models"
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
}

func RemoveCartIteam(){

}

func BuyIteamFromcart(){

}

func InstantBuyer(){

}

