# Prometheus integration


The application is integrated with Prometheus and Grafana to provide metrics observability.
Currently, it is set up to collect metrics from all services, databases and NATS.

### Folder Organization

The `./prometheus` folder is organized as follows:

- `postgresExporter`: Settings related to the PostgresSQL metrics exporter that allow Prometheus to scrape the database metrics
- `redisExporter`: Settings related to the PostgresSQL metrics exporter that allow Prometheus to scrape the database metrics
- `timescaledbExporter`: Settings related to the PostgresSQL metrics exporter that allow Prometheus to scrape the database metrics
- `prometheus.yml`: Prometheus scrape targets configuration

Note that the microservices, ArangoDB and NATS provide endpoints for Prometheus to scrape metrics directly and thus they do not need exporters. 