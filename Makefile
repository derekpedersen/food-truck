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

# docker is set to use the pre-compiled food-truck.linux as it's a debian based image
docker:
	docker build ./ -t food-truck:latest --no-cache
	docker run --name food-truck --rm -i -t food-truck