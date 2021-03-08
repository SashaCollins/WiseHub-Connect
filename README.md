# WiseHub-Connect
Bachelorarbeit: Konzeption und Umsetzung eines Dashboards mit Plugin-Architektur / Conception and implementation of a dashboard with plugin-architecture

## Backend (Go)
### Requirements
- Go v1.15.8

### Environment variables

| Status   |      Message      |  Example |
|----------|:-------------:|------:|
| MAIL_HOST |  SMTP server address | mx.example.com |
| MAIL_PORT |    SMTP port (TLS)   |   587 |
| MAIL_USERNAME | SMTP name | test@example.com |
| MAIL_PASSWORD | SMTP password | test123!! |
| DB_DRIVER | Database driver | sqlite,mysql,sqlserver,postgres |
| DB_NAME | Database name | wisehub.db |
| DB_USERNAME | Database username| foo |
| DB_PASSWORD | Database password | bar |
| DB_SSL_MODE | Database ssl mode | enable,disable |
| TARGET_URL | Target url for email validation | http://wisehub.localhost |


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
