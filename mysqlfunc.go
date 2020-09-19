package mysqlfunc

import (
	"database/sql"
	// Need this to connect to mysql
	_ "github.com/go-sql-driver/mysql"
)

//Column type to add or modify column in tables
//Check mysqlfunc.DataType for suggestions of dataType
//DataTypeParam is for stuff like VARCHAR(dataTypeParam)
type Column struct {
	name          string
	dataType      int
	dataTypeParam int
}

//DataTypeInt of type INT
var DataTypeInt = 0

//DataTypeVarChar of type VarChar
var DataTypeVarChar = 1

//GetData to get a map[int][string]interface
func GetData(queryStr string, sqlStr string) (map[int]map[string]interface{}, error) {
	db, err := sql.Open("mysql", sqlStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query(queryStr)
	if err != nil {
		return nil, err
	}
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	count := len(columns)
	tableData := make(map[int]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	a := 0
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData[a] = entry
		a++
	}
	return tableData, nil
}

// cmd
// git config --global user.email "dohyung97022@gmail.com"
// git config --global user.name "doe"

// add
// git add .
// git commit -m "commit_text"
// git push origin master

// sqlStr := "id:password@tcp(adiy-db.cxdzwqqcqoib.us-east-1.rds.amazonaws.com:3306)/adiy"
// queryStr := "SELECT * FROM channels"
