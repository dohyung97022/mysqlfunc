# mysqlfunc

### go functions for simple mysql

<br />

# How to use

#### Import

go get github.com/dohyung97022/mysqlfunc

```go
import("github.com/dohyung97022/mysqlfunc")
```

#### Initiate params

```go
	id := "insert_id"
	ps := "insert_ps"
	endpoint := "projectname-db.cxdzwqqcqoib.us-east-1.rds.amazonaws.com"
	port := 3306
	schema := "schema_name"

	err := Init(id, ps, endpoint, port, schema)
```

#### Send params to functions

```go
queryStr := "SELECT * FROM my_table"
v, err := mysqlfunc.GetQuery(queryStr)
```

<br />

# functions

<br />

#### GetQuery

most basic way to get any data by query

```go
queryStr := "SELECT * FROM my_table"

v, err := mysqlfunc.GetQuery(queryStr) (map[int]map[string]interface{}, error)
```

#### ExecQuery

use this if you don't care about the return value and just want to execute (just a wrapper for DB.Exec)

```go
err := mysqlfunc.ExecQuery(queryStr) error
```

#### InsertData

Insert data to a table (DataNames and data must be in the same order)

```go
dataNames := []string{"abouts_varchar", "age_int", "birth_date_time", "male_bool"}
var data []interface{}
data = append(data, "Hello, world", 24, time.Now(), true)

err = InsertData("test", dataNames, data) (error)
```

#### GetColNames

Get all column names from a table

```go
v, err := mysqlfunc.GetColNames(table string) (colNames []string, err error)
```

#### GetColNameTypes

Get all column names and types

```go
v, err := mysqlfunc.GetColNameTypes(table string) (map[string]string, error)
```

#### ClearTable

Clears all data from a table, use with caution!  
Takes a boolean to check if you want the auto increment to reset to 1. (false in some cases)

```go
err := mysqlfunc.ClearTable(table string, resetIncrement bool) error
```
