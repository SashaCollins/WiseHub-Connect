# WiseHub-Connect
Bachelorarbeit: Konzeption und Umsetzung eines Dashboards mit Plugin-Architektur / Conception and implementation of a dashboard with plugin-architecture

## Backend (Go)
### Requirements
- Go v1.15.8

### Environment variables
Status | Message
---: | :---
MAIL_HOST | 
MAIL_PORT |
MAIL_USERNAME |
MAIL_PASSWORD |
DB_DRIVER |
DB_NAME |
DB_USERNAME |
DB_PASSWORD | 
DB_SSL_MODE |
TARGET_URL |

#### Create Plugins
- go build -buildmode=plugin -o model/plugins/github/github.so model/plugins/github/Github.go 

#### Run this Project
To run this project you will need <strong>go, gcc or musl</strong> installed.
Create the plugins by run the command above. Clone this repository, build the plugins and run <strong>go build -o main && ./main</strong>.
You will also need a reverse-proxy (nginx, apache,...) to host your WiseHub-Connect.

### Status Codes and Ports
#### HTTP Status
Status | Message
---: | :---
666 | user already exists
667 | invalid email or password
668 | invalid email
669 | invalid data
670 | invalid token

#### Ports
Port | Listener
---: | :---
9010 | WiseHubConnect - Backend

## Frontend (Vuejs)
The frontend is in the view directory.
