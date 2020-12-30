## Audit Service

This folder contains code related to the audit service. It is organized as follows:

- `Server`: Audit service code

Additionally, this folder contains the following files:

- `Dockerfile`: Build the image for the audit service
- `docker-compose.env`: Environment variable required to run the service when running the service



Note that these docker files expect to be built using the root folder of the whole project as the context
(i.e.: do not build the docker build directly in this folder)