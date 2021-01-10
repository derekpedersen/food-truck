# build binary using local machines criteria
build:
	go build -o bin/food-truck

# build binary targeting windows
build-windows:
	go build -o bin/food-truck.windows GOOS=windows GOARCH=386

# build binary targeting linux
build-linux:
	go build -o bin/food-truck.linux GOOS=linux GOARCH=386

# build binary targeting mac
build-mac:
	go build -o bin/food-truck.mac GOOS=mac GOARCH=386

# docker image
docker:
	docker build ./ -t food-truck:latest --no-cache