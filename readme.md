# Super tiny clean architecture example with Golang

```
$ curl -i -X GET localhost:8080
HTTP/1.1 200 OK
Date: Mon, 01 Jul 2019 01:19:07 GMT
Content-Length: 0

$ curl -i -X PUT localhost:8080 -d helloworld
HTTP/1.1 200 OK
Date: Mon, 01 Jul 2019 01:19:28 GMT
Content-Length: 0

$ curl -i -X GET localhost:8080
HTTP/1.1 200 OK
Date: Mon, 01 Jul 2019 01:19:30 GMT
Content-Length: 10
Content-Type: text/plain; charset=utf-8

helloworld%
```

# references

- https://www.pospome.work/entry/2017/10/11/023848
