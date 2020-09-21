package mysqlfunc

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"

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

	//DB database from func Init
	DB *sql.DB

	err error
)

//Init to initiate db
func Init(id string, ps string, endpoint string, port int, schema string) error {
	DB, err = sql.Open("mysql", id+":"+ps+"@tcp("+endpoint+":"+strconv.Itoa(port)+")/"+schema+"?multiStatements=true")
	if err != nil {
		return err
	}
	err = DB.Ping()
	if err != nil {
		return err
	}
	return nil
}

//GetQuery most basic way to get any data by query
func GetQuery(queryStr string) (map[int]map[string]interface{}, error) {
	rows, err := DB.Query(queryStr)
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
	if len(tableData) == 0 {
		return nil, errors.New("the query has no value")
	}
	return tableData, nil
}

//ExecQuery just executes the query if you don't care about the return value. Return nil if success
func ExecQuery(queryStr string) error {
	_, err = DB.Exec(queryStr)
	if err != nil {
		return err
	}
	return nil
}

// GetColNames to a get all column names
func GetColNames(table string) (colNames []string, err error) {
	rows, err := DB.Query("SELECT * FROM " + table + " LIMIT 1")
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
func GetColNameTypes(table string) (map[string]string, error) {
	rows, err := DB.Query("SELECT COLUMN_NAME, DATA_TYPE FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = '" + table + "';")

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	count := len(columns)
	tableData := make(map[string]string, 0)
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

		var v2 string
		val = values[1]
		b, ok = val.([]byte)
		if ok {
			v2 = string(b)
		} else {
			return nil, errors.New("type of column is not string. v2 = " + v2)
		}

		tableData[v1] = v2
	}
	return tableData, nil
}

// InsertData to put a data to a table
func InsertData(table string, dataNames []string, data []interface{}) error {
	if len(dataNames) != len(data) {
		return errors.New("error : InsertData parameter len(dataNames) does not match len(data)")
	}
	var dataNamesStr strings.Builder
	var dataStr strings.Builder

	for c, v := range dataNames {
		dataNamesStr.WriteString(v)
		dataStr.WriteString("?")
		if c != len(dataNames)-1 {
			dataNamesStr.WriteString(",")
			dataStr.WriteString(",")
		}
	}
	q, err := DB.Prepare("INSERT INTO " + table + "(" + dataNamesStr.String() + ") VALUES(" + dataStr.String() + ")")
	if err != nil {
		return err
	}
	_, err = q.Exec(data...)
	if err != nil {
		return err
	}
	return nil
}

// GetDataOf gets data of colNames from table
func GetDataOf(table string, colNames []string) (map[int]map[string]interface{}, error) {
	var b strings.Builder
	b.WriteString("SELECT ")
	for c, v := range colNames {
		b.WriteString(v)
		if c != len(colNames)-1 {
			b.WriteString(",")
		}
	}
	b.WriteString(" FROM ")
	b.WriteString(table)
	return GetQuery(b.String())
}

// Where query
type Where struct {
	a  string
	is string
	b  interface{}
}

// GetDataOfWhere gets data that matches where. ex(where:="")
//this needs fixing for sure. make a type. like "a" "is" "b"
//  add '' into everyting. like string int datetime all sorts of things!
func GetDataOfWhere(table string, colNames []string, where []Where) (map[int]map[string]interface{}, error) {
	var b strings.Builder
	b.WriteString("SELECT ")
	for c, v := range colNames {
		b.WriteString(v)
		if c != len(colNames)-1 {
			b.WriteString(",")
		}
	}
	b.WriteString(" FROM ")
	b.WriteString(table)
	b.WriteString(" WHERE ")
	for c, v := range where {
		b.WriteString(v.a)
		b.WriteString(" ")
		b.WriteString(v.is)
		b.WriteString(" '")
		switch v := v.b.(type) {
		case int:
			fmt.Println("int:", v)
		case float64:
			fmt.Println("float64:", v)
		default:
			fmt.Println("unknown")
		}
		b.WriteString(v.b.(string))
		b.WriteString("'")
		if c != len(where)-1 {
			b.WriteString(" AND ")
		}
	}
	fmt.Printf(b.String())
	return GetQuery(b.String())
}

// ClearTable Clears all data from a table. Use with caution!
func ClearTable(table string, resetIncrement bool) error {
	var q strings.Builder
	q.WriteString("DELETE FROM ")
	q.WriteString(table)
	if resetIncrement {
		q.WriteString("; ALTER TABLE test AUTO_INCREMENT = 0;")
	}
	return ExecQuery(q.String())
}
