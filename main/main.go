package main

import (
	"fmt"

	"github.com/dohyung97022/mysqlfunc"
)

func main() {
	fmt.Print("ps: ")
	var ps string
	fmt.Scanln(&ps)

	sqlStr := "dohyung97022:" + ps + "@tcp(adiy-db.cxdzwqqcqoib.us-east-1.rds.amazonaws.com:3306)/adiy"
	db, err := mysqlfunc.Init(sqlStr)
	if err != nil {
		fmt.Printf("error : %v\n", err)
	}
	queryStr := "SELECT * FROM channels"
	v, err := mysqlfunc.GetData(queryStr, db)
	if err != nil {
		fmt.Printf("error : %v\n", err)
	}
	fmt.Printf("v : %v\n", v)
}
