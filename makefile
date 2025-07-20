default: build

PREFIX=/home/dmitri/local/bin # Adjust to the prefix you use.

build:
	go build main.go parser.go intf.go &&\
		mv ./main ./build/main

install:
	install -m 755 ./build/main $(PREFIX)

test:
	go run test.go parser.go intf.go
