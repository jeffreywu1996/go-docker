# Go-naive-user-service
這是一個簡單的go user service api. 在local 上可以使用 go build 然後go run啟動server 再來client可以使用以下方法操作
## Get user

Request example
```bash
curl --request GET \
  --url http://localhost:8080/user/<user_id>
```

Success example

```bash
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Mon, 26 Jul 2021 15:39:09 GMT
Content-Length: xxx
Connection: close

{
  "user_id": "foobar",
	"age": 30
}
```

## Create user

Request example

```bash
curl --request POST \
  --url http://localhost:8080/user \
  --header 'content-type: application/json' \
  --data '{"user_id": "foobar","age": 30}'
```

Success example

```bash
HTTP/1.1 204 No Content
```

## Delete user

Request example

```bash
curl --request DELETE \
  --url http://localhost:8080/user/<user_id>
```
Success example

```bash
HTTP/1.1 204 No Content
```


## Docker commands
Build image
```
docker build -f Dockerfile -t go-test .
```

Run docker image
```
docker run -it -p 8000:8000 go-test:latest
```
