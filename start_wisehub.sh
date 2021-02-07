#!/usr/bin/bash

go version &> /dev/null
#screen &> /dev/null

if [ $? -ne 1 ]
  then
    go build -buildmode=plugin -o model/plugins/drone/drone.so model/plugins/drone/Drone.go
    go build -buildmode=plugin -o model/plugins/github/github.so model/plugins/github/Github.go

    screen -dmS "WiseHubBackend" go run main.go
  else
    echo "Golang is not installed. Please install golang!"
fi

npm version &> /dev/null

if [ $? -ne 1 ]
  then
    # shellcheck disable=SC2164
    cd ./view
    echo fs.inotify.max_user_watches=524288 | sudo tee -a /etc/sysctl.conf && sudo sysctl -p
    npm i
    screen -dmS "WiseHubFrontend" npm run serve
  else
    echo "nodejs and npm are not installed. Please install nodejs and npm!"
fi
