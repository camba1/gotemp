# Prometheus integration

**Work in Progress**

The application is integrated with Prometheus and Grafana to provide metrics observability.
Currently, it is set up to collect metrics from all databases and NATS metrics when the app is running in Docker. 
The ability to pull metrics from the microservices is coming in the near future. K8s integration will be forth coming as well.

### Folder Organization

The `./prometheus` folder is organized as follows:

- `postgresExporter`: Settings related to the PostgresSQL metrics exporter that allow Prometheus to scrape the database metrics
- `redisExporter`: Settings related to the PostgresSQL metrics exporter that allow Prometheus to scrape the database metrics
- `timescaledbExporter`: Settings related to the PostgresSQL metrics exporter that allow Prometheus to scrape the database metrics
- `prometheus.yml`: Prometheus scrape targets configuration

Note that ArangoDB and NATS provide end point for Prometheus to scrape metrcis directly and thus there is no need for exporters. 