package main

import (
	"context"
	"errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Account struct {
	gorm.Model
	Name    string
	Balance uint
}

type Transaction struct {
	gorm.Model
	FromAccountId uint
	ToAccountId   uint
	Amount        uint
}

func initDB() *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)
	db, _ := gorm.Open(sqlite.Open("study/task03/test.db"), &gorm.Config{Logger: newLogger})
	db.AutoMigrate(&Account{}, &Transaction{})
	return db
}

func initData(db *gorm.DB) {
	ac1 := Account{Name: "Alice", Balance: 200}
	ac2 := Account{Name: "Bob", Balance: 300}
	ctx := context.Background()
	gorm.G[Account](db).Create(ctx, &ac1)
	gorm.G[Account](db).Create(ctx, &ac2)
}
func transferMoney(from string, to string, money uint, db *gorm.DB) {
	db.Transaction(func(tx *gorm.DB) error {
		ctx := context.Background()
		ac1, _ := gorm.G[Account](tx).Where("name = ?", from).First(ctx)
		ac2, _ := gorm.G[Account](tx).Where("name = ?", to).First(ctx)
		if ac1.Balance < money {
			log.Printf("Trasfer Money from %s to %s Failed ", from, to)
			return errors.New("transfer Money Failed")
		}
		ac1.Balance -= money
		ac2.Balance += money
		transaction := Transaction{FromAccountId: ac1.ID, ToAccountId: ac2.ID, Amount: money}
		gorm.G[Transaction](tx).Create(ctx, &transaction)
		gorm.G[Account](tx).Where("name = ?", from).Update(ctx, "balance", ac1.Balance)
		gorm.G[Account](tx).Where("name = ?", to).Update(ctx, "balance", ac2.Balance)
		return nil
	})

}
func main() {
	db := initDB()
	//initData(db)
	//完成一个方法让Bob对Alice转账
	transferMoney("Bob", "Alice", 100, db)
}
