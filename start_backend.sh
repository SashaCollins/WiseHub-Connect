#!/usr/bin/bash

go build -buildmode=plugin -o model/plugins/drone/drone.so model/plugins/drone/Drone.go
go build -buildmode=plugin -o model/plugins/github/github.so model/plugins/github/Github.go


go run main.go

