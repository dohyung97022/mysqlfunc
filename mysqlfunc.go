package mysqlfunc

import (
	"database/sql"
	"errors"

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

// GetColNames to a get all column names
func GetColNames(table string, db *sql.DB) (colNames []string, err error) {
	rows, err := db.Query("SELECT * FROM " + table + " LIMIT 1")
	if err != nil {
		return nil, err
	}
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	for _, col := range columns {
		colNames = append(colNames, col)
	}
	return colNames, nil
}

// GetColNameTypes to a get all column name:type
func GetColNameTypes(table string, db *sql.DB) (map[string]interface{}, error) {
	// "SELECT COLUMN_NAME, DATA_TYPE FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_SCHEMA = 'adiy' AND TABLE_NAME = 'channels';"
	rows, err := db.Query("SELECT COLUMN_NAME, DATA_TYPE FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = '" + table + "';")

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	count := len(columns)
	tableData := make(map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)

		var v1 string
		val := values[0]
		b, ok := val.([]byte)
		if ok {
			v1 = string(b)
		} else {
			return nil, errors.New("title of column is not string. v1 = " + v1)
		}

		var v2 interface{}
		val = values[1]
		b, ok = val.([]byte)
		if ok {
			v2 = string(b)
		} else {
			v2 = val
		}

		tableData[v1] = v2
	}
	return tableData, nil
}
