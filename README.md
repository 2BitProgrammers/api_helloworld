# 2bitprogrammers/api_helloworld

This is a simple "Hello World" API example.  It is meant to be used for instructional purposes only.

The API listens on port:  1234

API Enpoints:
* **/status** (GET) - this states whether the app is up and healthy
* There are multiple endpoint which returns "Hello World" (do the same thing):
  * **/hello** (GET)
  * **/hello-world** (GET)
  * **/hello_world** (GET)
  * **/helloworld** (GET)
  * **/helloWorld** (GET)

## Run as Standalone GoLang App
This will run the application with the go application.  It assumes that you have installed your golang environment correctly.

```bash
$ cd src
$ go run main.go

2bitprogrammers/api_helloworld v2005.11a
www.2BitProgrammers.com
Copyright (C) 2020. All Rights Reserved.

2005/12/08 21:30:31 Starting App on Port 1234

CTRL+C
```

## Run within Docker
This will run the components on your local system without using minikube or kubernetes.

### Building the Docker Image
```bash
$ docker build . -t 2bitprogrammers/api_helloworld

Sending build context to Docker daemon  10.75kB
Step 1/11 : FROM golang:alpine AS builder
 ---> b3bc898ad092
Step 2/11 : ENV GO111MODULE=on     CGO_ENABLED=0     GOOS=linux     GOARCH=amd64
 ---> Using cache
 ---> 8462443c0070
Step 3/11 : WORKDIR /build
 ---> Using cache
 ---> 99600623930c
Step 4/11 : COPY $PWD/src/go.mod .
 ---> Using cache
 ---> 4c95b0026d27
Step 5/11 : COPY $PWD/src/main.go .
 ---> Using cache
 ---> bed42fd7c929
Step 6/11 : RUN go mod download
 ---> Using cache
 ---> ad0879d66a55
Step 7/11 : RUN go build -o api_helloworld .
 ---> Using cache
 ---> 5f0ed7a3b18d
Step 8/11 : FROM scratch
 --->
Step 9/11 : WORKDIR /
 ---> Using cache
 ---> a66c59ea194a
Step 10/11 : COPY --from=builder /build/api_helloworld .
 ---> 15af0b05ae0a
Step 11/11 : ENTRYPOINT [ "/api_helloworld" ]
 ---> Running in 2439d4227935
Removing intermediate container 2439d4227935
 ---> 1d3b14d1b9d1
Successfully built 1d3b14d1b9d1
Successfully tagged 2bitprogrammers/api_helloworld:latest
SECURITY WARNING: You are building a Docker image from Windows against a non-Windows Docker host. All files and directories added to build context will have '-rwxr-xr-x' permissions. It is recommended to double check and reset permissions for sensitive files and directories.
```

### Image Status
```bash
$ docker images

REPOSITORY                        TAG          IMAGE ID            CREATED              SIZE
2bitprogrammers/api_helloworld    latest       1d3b14d1b9d1        About a minute ago   6.57MB
```

### Running the Container
```bash
$ docker run --rm --name "api_helloworld" -p 1234:1234 2bitprogrammers/api_helloworld 

2bitprogrammers/api_helloworld v2005.11a
www.2BitProgrammers.com
Copyright (C) 2020. All Rights Reserved.

2005/12/08 21:30:31 Starting App on Port 1234

CTRL+C
```

### Check the Container Status (docker)
```bash
$ docker ps

CONTAINER ID    IMAGE                             COMMAND               CREATED              STATUS              PORTS                    NAMES
6d7546f90be1    2bitprogrammers/api_helloworld    "/api_helloworld"     About a minute ago   Up About a minute   0.0.0.0:1234->1234/tcp   api_helloworld
```

### Check API Status (health check):
```bash
$ curl http://localhost:1234/status

{"date":"2020-07-10T03:17:02.9034438-07:00","statusCode":200,"statusText":"OK","data":"{ \"healthy\": true}","errors":"","request":{"uri":"/status","method":"GET","payload":""}}
```

### Watch Container Logs
```bash
$ docker logs -f 2bitprogrammers/api_helloworld

2bitprogrammers/api_helloworld v2005.11a
www.2BitProgrammers.com
Copyright (C) 2020. All Rights Reserved.

2020/12/08 21:30:31 Starting App on Port 1234
2020/12/08 21:33:14 GET /status 200
2020/12/08 21:33:23 GET /hello 200
2020/12/08 21:33:55 GET /helloworld 200

CTRL+C
```

### Stopping the Container
```bash
$ docker stop api_helloworld
```

## Using the API
For the below examples, we will assume the following:
* Server:  locahost (127.0.0.1)
* Bind Port: 1234
* Method: GET
* URI: /helloworld 
* Body Data:   _N/A_

```bash
$ curl http://127.0.0.1:1234/helloworld

{"date":"2020-12-08T21:33:55.6009605Z","statusCode":200,"statusText":"OK","data":"\"Hello World !!!\"","errors":"","request":{"uri":"/helloworld","method":"GET","payload":""}}
```

