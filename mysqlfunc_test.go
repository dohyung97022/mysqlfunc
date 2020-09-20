package mysqlfunc

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	//always delete password before save!!!!
	ps := ""
	sqlStr := "dohyung97022:" + ps + "@tcp(adiy-db.cxdzwqqcqoib.us-east-1.rds.amazonaws.com:3306)/adiy"
	err := Init(sqlStr)
	if err != nil {
		fmt.Printf("error : %v\n", err)
	}
	defer DB.Close()

	v, err := GetColNameTypes("channels")
	if err != nil {
		fmt.Printf("error : %v\n", err)
	}
	fmt.Printf("v : %v\n", v)
}
