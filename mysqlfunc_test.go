package mysqlfunc

import (
	"fmt"
	"testing"
)

//is this now not uploading?
//why is this uploading?

func TestMain(t *testing.T) {
	fmt.Print("ps: ")
	var ps string
	fmt.Scanln(&ps)

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
