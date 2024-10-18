default:
	go build main.go parser.go intf.go &&\
		mv ./main ./build/main
