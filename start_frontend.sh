#!/usr/bin/bash

npm version &> /dev/null

if [ $? -ne 1 ]
  then
    # shellcheck disable=SC2164
    cd ./view
    echo fs.inotify.max_user_watches=524288 | sudo tee -a /etc/sysctl.conf && sudo sysctl -p
    npm i
    screen -dS "frontend" npm run serve
  else
    echo "nodejs and npm are not installed. Please install nodejs and npm!"
fi