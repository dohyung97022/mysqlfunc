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
v, err := mysqlfunc.GetData(queryStr)
```

<br />

# functions

<br />

#### GetData

Simple function to get data from mysql

```go
v, err := mysqlfunc.GetData(queryStr) (map[int]map[string]interface{}, error)

// v = map[0:map[id:1 name:Jhon comment:Thank] 1:map[id:2 name:Sam comment:You]]
```

#### GetColNames

Get all column names from a table

```go
v, err := mysqlfunc.GetColNames(table string) (colNames []string, err error)

// v = [id name comment]
```

#### GetColNameTypes

Get all column names and types

```go
v, err := mysqlfunc.GetColNameTypes(table string) (map[string]interface{}, error)

// v = [id:int name:varchar comment:varchar]
```
