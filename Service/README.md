# This is test application that shows how to build Golang BackEnd, put it to the Docker Image and run it.

### Files:
  - **app.build.bat** - builds Golang application for Linux
  - **docker.build.bat** - builds Docker Image from Dockerfile and increase version by 1
  - **docker.containers.bat** - shows list of all Docker Containers
  - **docker.images.bat** - shows list of Docker Images
  - **docker.run.bat** - runs last compiled Docker Image
  - **Dockerfile** - Docker file that describes what will be inside Docker Image (in our case this is only FullStack.Service.exe file)
  - **main.go** - main Golang project file 
  - **README.md** - file with readme information
  - **ver** - file with version (version is automaticaly increasing when you build new Docker Image)

### To run this test application from zero you need:
  - [Golang installation](https://golang.org/doc/install)
  - [Docker installation](https://docs.docker.com/docker-for-windows/install/)

### Order of work:
  - Install Golang
  - Install Docker
  - Build project: run **app.build.bat** file
  - Make Docker Image: run **docker.build.bat** file
  - Run Docker Image: run **docker.run.bat** file
  - In Browser open url: [localhost:38080](http://localhost:38080) and you will see text "Hello, !"
