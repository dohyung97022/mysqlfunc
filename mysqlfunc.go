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
	Name          string
	DataType      int
	DataTypeParam int
}

var (
	//DataTypeInt of type INT
	DataTypeInt = 0

	//DataTypeVarChar of type VarChar
	DataTypeVarChar = 1
)

//Init to initiate db
func Init(sqlStr string) (*sql.DB, error) {
	db, err := sql.Open("mysql", sqlStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

//GetData to get a map[int][string]interface
func GetData(queryStr string, db *sql.DB) (map[int]map[string]interface{}, error) {

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

//InsertData to a table
func InsertData() {

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

//go get github.com/dohyung97022/mysqlfunc
//go get -u github.com/dohyung97022/mysqlfunc
