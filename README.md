# mysqlfunc
### go functions for simple mysql
***
## How to use
### Import

```go
go get github.com/dohyung97022/mysqlfunc
import("github.com/dohyung97022/mysqlfunc")
```

Initiate mysql parameters

```go
sqlStr := "username:password@tcp(post-aws-rds-database-endpoint:3306)/post-schema-name"
db, err := mysqlfunc.Init(sqlStr)
```

Send parameters to functions

```go
queryStr := "SELECT * FROM my_table"
mysqlfunc.GetData(queryStr, db) (map[int]map[string]interface{}, error)
```

Dont forget to close db at the end.

```go
defer db.Close()
```

functions list.
