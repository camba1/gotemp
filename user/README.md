## User Service

This folder contains code related to the user service. It is organized as follows:

- `Client`: Contains a client service that calls the user service to perform multiple operations
- `Proto`: Proto buffer messages and services definitions
- `Server`: User service 

Additionally, this folder contains the following files:

- `DockerfileCli`: Build the image for the user client  
- `Dockerfile`: Build the image for the user service
- `docker-compose.env`: Environment variable required to run the service when running the service with docker-compose

Note that these docker files expect to be built using the root folder of the whole project as the context 
(i.e.: do not run the docker build directly in this folder)

