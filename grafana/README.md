# Grafana integration

**Work in Progress**

The application is integrated with Prometheus and Grafana to provide metrics observability.
Currently, it is set up to collect metrics from all services, databases and NATS metrics when the app is running in Docker.
K8s integration will be forth coming as well.

### Folder Organization

This folder is organized as follows:

- `Provisioning`: Folder that contains Grafana configuration
  
    - `dashboards`: contains all the dashboards that will be available by default when it first comes up.
    - `datasources`: Definition of the data sources used by Grafana
    
- `docker-compose.env`: Environment variables required to run the service when running the Grafana with docker-compose