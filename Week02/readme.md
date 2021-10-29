#### dao层还是区分一下，

```golang

if err != nil {
      if errors.Is(err, sql.ErrNoRows) {
      return nil, fmt.Errorf("sql:%s error:[%w]", sql, code.ErrNotFound)
      }
     return nil, errors.Wrapf(code.ErrDBServer, fmt.Sprintf("query: %s error(%v)", sql, err))
}
```

在biz层根据
errors.Is(err, code.ErrNotFound) 来判断

### 代码运行 

go run main.go