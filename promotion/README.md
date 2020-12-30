## Promotion Service

This folder contains code related to the promotion service. It is organized as follows:

- `Client`: Contains a client service that calls the product service to perform multiple operations
- `Proto`: Proto buffer messages and services definitions
- `Server`: Product service

Additionally, this folder contains the following files:

- `DockerfileCli`: Build the image for the promotion client
- `Dockerfile`: Build the image for the promotion service
- `docker-compose.env`: Environment variables required to run the service when running the service
- `docker-compose-cli.env`: Environment variable required to run the client when running the service 

Note that these docker files expect to be built using the root folder of the whole project as the context
(i.e.: do not build the docker build directly in this folder)