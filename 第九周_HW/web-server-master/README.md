
# web-server
## author T——Lee
开发 web 服务程序


# 框架使用 使用时若未安装go get即可
github.com/gorilla/mux

github.com/unrolled/render

github.com/urfave/negroni

# curl测试：
命令：curl -v http://localhost:7000/web/lg
结果如下：

```
*   Trying 127.0.0.1...
* Connected to localhost (127.0.0.1) port 7000 (#0)
> GET /web/lg HTTP/1.1
> Host: localhost:7000
> User-Agent: curl/7.47.0
> Accept: */*
> 
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=UTF-8
< Date: Tue, 14 Nov 2017 12:10:03 GMT
< Content-Length: 19
< 
{
  "-_-!": "lg"
}
* Connection #0 to host localhost left intact
```



# ab测试：
安装ab 使用命令：sudo apt install apache2-utils

命令： ab -n 100 -c 10 http://localhost:7000/web/lg
结果如下：
```
This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient).....done


Server Software:        
Server Hostname:        localhost
Server Port:            7000

Document Path:          /web/lg
Document Length:        19 bytes

Concurrency Level:      10
Time taken for tests:   0.009 seconds
Complete requests:      100
Failed requests:        0
Total transferred:      14200 bytes
HTML transferred:       1900 bytes
Requests per second:    11721.96 [#/sec] (mean)
Time per request:       0.853 [ms] (mean)
Time per request:       0.085 [ms] (mean, across all concurrent requests)
Transfer rate:          1625.51 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0       0
Processing:     0    1   0.3      1       3
Waiting:        0    1   0.2      1       1
Total:          0    1   0.3      1       3

Percentage of the requests served within a certain time (ms)
  50%      1
  66%      1
  75%      1
  80%      1
  90%      1
  95%      1
  98%      2
  99%      3
 100%      3 (longest request)
```
