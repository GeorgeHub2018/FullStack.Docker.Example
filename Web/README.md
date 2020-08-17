# This is test application that shows how to put Web site to the Docker Image and run it.

### Files:
  - **js/service.js** - JavaScript file that allow to make Api request to Server side
  - **docker.build.bat** - builds Docker Image from Dockerfile and increase version by 1
  - **docker.containers.bat** - shows list of all Docker Containers
  - **docker.images.bat** - shows list of Docker Images
  - **docker.run.bat** - runs last compiled Docker Image
  - **Dockerfile** - Docker file that describes what will be inside Docker Image (in our case this is only js/service.js and index.html files)
  - **index.html** - main HTML file in this Web project
  - **README.md** - file with readme information
  - **ver** - file with version (version is automaticaly increasing when you build new Docker Image)

### To run this test application from zero you need:
  - [Docker installation](https://docs.docker.com/docker-for-windows/install/)
  - [FullStack.Docker.Example.Service installation](https://github.com/GeorgeHub2018/FullStack.Docker.Example/tree/master/Service)


### Order of work:
  - Install Docker
  - Make FullStack.Docker.Example.Service Docker Image and run it
  - Make Docker Image: run **docker.build.bat** file
  - Run Docker Image: run **docker.run.bat** file
  - In Browser open url: [localhost:38081](http://localhost:38081) and you will see message "Hello, World!"
