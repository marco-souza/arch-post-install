ALL: run

run: main.go utils.go
	go run main.go utils.go

build: main.go utils.go
	go build main.go utils.go