package mysqlfunc

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	//always delete password before save!!!!
	ps := ""

	sqlStr := "dohyung97022:" + ps + "@tcp(adiy-db.cxdzwqqcqoib.us-east-1.rds.amazonaws.com:3306)/adiy"
	db, err := Init(sqlStr)
	if err != nil {
		fmt.Printf("error : %v\n", err)
	}
	defer db.Close()

	v, err := GetColNameTypes("channels", db)
	if err != nil {
		fmt.Printf("error : %v\n", err)
	}
	fmt.Printf("v : %v\n", v)
}
