package main

import (
	"context"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// Student gorm 严格区分model大小写
type Student struct {
	gorm.Model
	Name  string
	Age   uint
	Grade string
}

func initTable(db *gorm.DB) {
	//定义原生ddl
	createTable := "" +
		"CREATE TABLE IF NOT EXISTS `students` (" +
		"`id` INTEGER PRIMARY KEY AUTOINCREMENT," +
		"`name` TEXT," +
		"`age` INTEGER," +
		"`grade` TEXT" +
		")"
	db.Exec(createTable)
}

func saveBatch(db *gorm.DB) {
	stubs := []Student{
		{Name: "david"},
		{Name: "merry"},
		{Name: "peter"},
	}
	_ = gorm.G[Student](db).CreateInBatches(context.Background(), &stubs, len(stubs))
}

func main() {
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
	err := db.AutoMigrate(&Student{})
	if err != nil {
		fmt.Println(err)
		return
	}
	//initTable(db)
	ctx := context.Background()
	//insert data
	saveBatch(db)
	stu := Student{Name: "张三", Age: 20, Grade: "三年级"}
	gorm.G[Student](db).Create(ctx, &stu)
	stubs := gorm.G[Student](db).Where("age > ?", 18)
	fmt.Println(stubs)
	result, _ := gorm.G[Student](db).Where("name = ?", "张三").Update(ctx, "grade", "三年级")
	fmt.Println(result)
	deleteRes, _ := gorm.G[Student](db).Where("age < ?", 15).Delete(ctx)
	fmt.Println(deleteRes)
	//db.
	//result := db.Save(&stu)
	//stu, _ := gorm.G[Student](db).Where("name=?", "David").First(ctx)
	//fmt.Println(result)
	//Drop table
	//gorm.G[Student](db).Exec(ctx, `DROP TABLE IF EXISTS students`)
}
