# A minimal WASM example in Go

Start server with
```bash
go run server.go --race
```

Compile WASM package with:
```
GOOS=js GOARCH=wasm go build -o main.wasm wasm/main.go
```

Go to: [http://localhost:8080/](http://localhost:8080/).
