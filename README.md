# Food Truck

This was written as a job application project for [Redfin](https://redfin.com).

## Delivery

One of the benefits of [Golang](https://golang.org) is you can target different platforms when compiling, so I was able to provide a pre-compiled binary for the following Operation Systems.

| Operating System | file |
| --- | --- |
| **Windows** | `./bin/food-truck.windows` |
| **Mac** | `./bin/food-truck.mac` |
| **Linux** | `./bin/food-truck.linux` |

### Docker

This project contains a [Dockerfile](https://docs.docker.com/engine/reference/builder/) which allows the user to build a docker image using a [debian](https://www.debian.org/) based container and linux x86 binary.

#### Execution

If you have the necessary prereqs for using a `Makefile` you can just run:

```
    make docker
```

If you only have Docker prereqs installed you can run:

```
	docker build ./ -t food-truck:latest --no-cache
	docker run --name food-truck --rm -i -t food-truck
```

### Local Compilation

If you want to compile this program locally you will need follow these steps:

1. [Install Golang following these instructions](https://golang.org/doc/install)
2. 

### Anonymity

In an effort to comply with requirement that this project be as annonymous as possible I wrote this project from the `/usr/local/go` folder. 

`sudo chown -R $(whoami) /usr/local`

