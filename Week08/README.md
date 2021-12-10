###问题1、命令:
```
redis-benchmark -d 10 -t get,set
```
###SET
|-|执行次数和耗时|每秒请求次数|
|----|----|----|
|10|100000 requests completed in 0.62 seconds |160771.70 requests per second|
|20|100000 requests completed in 0.64 seconds|157232.70 requests per second|
|50|100000 requests completed in 0.64 seconds|156985.86 requests per second|
|100|100000 requests completed in 0.61 seconds|162866.44 requests per second|
|200|100000 requests completed in 0.61 seconds|164744.64 requests per second|
|1k|100000 requests completed in 0.62 seconds|162074.56 requests per second|
|5k|100000 requests completed in 0.67 seconds|149031.30 requests per second|
###GET
|-|执行次数和耗时|每秒请求次数|
|----|----|----|
|10|100000 requests completed in 0.60 seconds|167504.19 requests per second|
|20|100000 requests completed in 0.62 seconds|161812.31 requests per second|
|50|100000 requests completed in 0.59 seconds|168634.06 requests per second|
|100|100000 requests completed in 0.60 seconds|167224.08 requests per second|
|200|100000 requests completed in 0.58 seconds|171821.30 requests per second|
|1k|100000 requests completed in 0.63 seconds|158730.16 requests per second|
|5k|100000 requests completed in 0.67 seconds|149476.83 requests per second|

##问题2、

```golang
go run main.go
```