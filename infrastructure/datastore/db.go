package datastore

import (
	"fmt"
    "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Database() *gorm.DB {
//	DBMS := "sqlserver://sa:p@$$w0rd@localhost:1433?database=GolangLearn"
    DBMS := "app:mitraumat@tcp(127.0.0.1:3306)/golanglearn?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(DBMS), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	return db
}
