test:
	go test -v ./...

clean:
	rm -f ./main

build:
	go build -v main.go

run:
	./main


