# WiseHub-Connect
Bachelorarbeit: Konzeption und Umsetzung eines Dashboards mit Plugin-Architektur / Conception and implementation of a dashboard with plugin-architecture


####HTTP Status
Status | Message
---: | :---
666 | user already exists
667 | invalid email or password
668 | invalid email
669 | invalid data
670 | invalid token


####Ports
Port | Listener
---: | :---
9010 | WiseHubConnect - Backend


# Create Plugins
- go build -buildmode=plugin -o model/plugins/github/github.so model/plugins/github/Github.go 


# Run this Project
To run this project you will need <strong>go, gcc and musl</strong> installed.
clone this repository, build the plugins and run <strong>go build -o main && ./main</strong>. 
You will also need an reverse-proxy (nginx, apache,...) to host your WiseHub-Connect.