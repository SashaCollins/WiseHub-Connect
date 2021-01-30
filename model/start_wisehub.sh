#!/usr/bin/bash

go build -buildmode=plugin -o plugins/drone/drone.so plugins/drone/Drone.go
go build -buildmode=plugin -o plugins/github/github.so plugins/github/Github.go
go build -buildmode=plugin -o plugins/heroku/heroku.so plugins/heroku/Heroku.go

go run wisehubConnect.go
