build:
	GOOS=js GOARCH=wasm go build -o main.wasm
exec:
	goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))'