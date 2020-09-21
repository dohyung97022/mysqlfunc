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
sqlStr := "username:password@tcp(post-aws-rds-database-endpoint:3306)/post-schema-name"
err := mysqlfunc.Init(sqlStr)
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

use this if you don't care about the return value and just want to execute
this is just a wrapper for DB.Exec

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
