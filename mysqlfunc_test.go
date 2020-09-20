package mysqlfunc

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	ps := ""

	sqlStr := "dohyung97022:" + ps + "@tcp(adiy-db.cxdzwqqcqoib.us-east-1.rds.amazonaws.com:3306)/adiy"
	db, err := Init(sqlStr)
	if err != nil {
		fmt.Printf("error : %v\n", err)
	}
	queryStr := "SELECT * FROM channels"
	_, err = GetData(queryStr, db)
	if err != nil {
		fmt.Printf("error : %v\n", err)
	}
	colNamesArray, err := GetColNames("channels", db)
	if err != nil {
		fmt.Printf("error : %v\n", err)
	}
	fmt.Printf("colNamesArray : %v\n", colNamesArray)

	defer db.Close()
}
