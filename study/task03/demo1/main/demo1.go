package main

import (
	"context"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("../test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	db.AutoMigrate(&Product{})
	//pdts := []Product{
	//	{Code: "2", Price: 200},
	//	{Code: "3", Price: 100},
	//	{Code: "4", Price: 500},
	//}
	//Create
	//err = gorm.G[Product](db).Create(ctx, &Product{Code: "1", Price: 100})
	//if err != nil {
	//	return
	//}
	//insert batch
	//err = gorm.G[Product](db).CreateInBatches(ctx, &pdts, len(pdts))
	//Read
	product1, err1 := gorm.G[Product](db).Where("id=?", 3).First(ctx)
	if err1 != nil {
		return
	}
	//Update
	fmt.Println(product1)
}
