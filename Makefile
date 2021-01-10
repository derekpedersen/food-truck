# build binary using local machines criteria
build:
	go build -o bin/food-truck

# build binary targeting windows
build-windows:
	GOOS=windows GOARCH=386 go build -o bin/food-truck.windows

# build binary targeting linux
build-linux:
	GOOS=linux GOARCH=386 go build -o bin/food-truck.linux 

# build binary targeting mac
build-mac:
	GOOS=darwin GOARCH=amd64 go build -o bin/food-truck.mac 

# build all the binaries!
build-all: build build-windows build-linux build-mac

# docker image
docker:
	docker build ./ -t food-truck:latest --no-cache