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
	mysqlfunc.GetData()
}
