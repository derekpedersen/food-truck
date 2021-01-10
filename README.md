# Food Truck

This was written as a job application project for [Redfin](https://redfin.com).

## Delivery

I've pre-compiled this project to run on Windows, Mac, or Linux distro that x86 compatible. 

### Windows

`./bin/food-truck.windows`

### Mac

`./bin/food-truck.mac`

### Linux

`./bin/food-truck.linux`

### Docker

This project contains a [Dockerfile](https://docs.docker.com/engine/reference/builder/) which allows the user to build a docker image using a [debian](https://www.debian.org/) based container.

### Local Compilation

#### Golang

This project was built using [Golang](https://golang.org). One of the benefits of [Golang](https://golang.org) is you can target different platforms when compiling, so I was able to provide a pre-compiled binary for the following Operation Systems.

If you wish to compile this program locally you will need to follow the steps

1. [Install Golang following these instructions](https://golang.org/doc/install)

### Anonymity

In an effort to comply with requirement that this project be as annonymous as possible I wrote this project from the `/usr/local/go` folder. 

`sudo chown -R $(whoami) /usr/local`

