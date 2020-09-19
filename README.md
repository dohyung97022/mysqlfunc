# mysqlfunc

go functions for simple mysql

Make mysql parameters.

```go
queryStr := "username:password@tcp(post-aws-rds-database-endpoint:3306)/post-schema-name"
sqlStr := "SELECT * FROM my_table"
```

Send parameters to functions.

```go
mysqlfunc.GetData(queryStr string, sqlStr string) (map[int]map[string]interface{}, error)
```

functions list.
