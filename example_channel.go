package main

// import (
// 	"fmt"
// 	"time"
// )

// type ShopData struct {
// 	Name string
// }

// type ProductData struct {
// 	Name string
// }

// type CartData struct {
// 	Shop    *ShopData
// 	Product *ProductData
// }

// func PlayChannel() {
// 	timeStart := time.Now()
// 	shop := LoadShopData()
// 	product := LoadProductData()
// 	cart := BuildCart(shop, product)
// 	fmt.Println(time.Since(timeStart).Seconds(), cart.Product.Name, "-", cart.Shop.Name)
// }

// func LoadShopData() *ShopData {
// 	result := new(ShopData)
// 	result.Name = apicallShop()
// 	return result
// }

// func LoadProductData() *ProductData {
// 	result := new(ProductData)
// 	result.Name = apicallProduct()
// 	return result
// }

// func apicallShop() string {
// 	<-time.Tick(time.Second)
// 	return "ini namanya shop"
// }

// func apicallProduct() string {
// 	<-time.Tick(time.Second)
// 	return "ini namanya product"
// }

// func BuildCart(shop *ShopData, product *ProductData) *CartData {
// 	return &CartData{
// 		Product: product,
// 		Shop:    shop,
// 	}
// }
