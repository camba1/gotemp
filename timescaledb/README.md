## TimeScaleDB usage

TimescaleDB time series database stores all historical audit data.
When a service modifies data, it sends a message to the pub/sub broker.
This message is then picked up by the audit service and stored in TimescaleDB

There are two folders in this directory:

- `timescaleDB`: Persists container data to the host. All data in the container is deleted if this folder is deleted.
- `timesacleInit`: Contains scripts to create tables, hypertables and seed data that runs the first time the DB is initialized

Due to the way the official TimescaleDB image is built, in order to populate the DB with our scripts, it is necessary to:

- Drop the `timescaleDB` folder so that the database is re-initialized
- Rebuild the custom timescaleDB image using the Dockerfile in the directory. This will copy our scripts to the correct location in the image
- Start container 
 