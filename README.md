### Running Steps (current - v1)

> Assuming you have go + curl installed already

<br />


1. Clone and move to repo
```
  git clone https://github.com/ShubhamTiwary914/tcp-http-1.1-server.git
  cd tcp-http-1.1-server
```

2. Run the Web Server
```
  go run main.go
```

3. Make Client -> Server Call (curl)
```
  curl -X GET -H "Content-Type: application/json" -d '{"name": "User"}' http://localhost:8080/
```


<br />

### To be Added (v2)
- Chunked Transfer Encoding
- Response Compression
- Connection Pooling
- Rate limit
- Connection Timeouts
