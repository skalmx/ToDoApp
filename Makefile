build:
	go build -o ./bin cmd/app/main.go

run:	build
	./bin

