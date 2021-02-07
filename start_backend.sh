#!/usr/bin/bash

go version &> /dev/null

if [ $? -ne 1 ]
  then
    go build -buildmode=plugin -o model/plugins/drone/drone.so model/plugins/drone/Drone.go
    go build -buildmode=plugin -o model/plugins/github/github.so model/plugins/github/Github.go
    screen -dS "backend" go run main.go
  else
    echo "Golang is not installed. Please install golang!"
fi