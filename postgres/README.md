## Postgres DB usage

The purpose of this folder is two fold:
1. Persist the data created by the PostgreSql container used by some services.
2. Seed the database with initial structure and data when the container is created for the first time

As such, we have two folders to handle those tasks:
- `postgresDB`: Volume mounted to the container to share the data with the host
- `postgresinit`: Volume mounted to the container. Contains the DB schemas and data to be used to seed the app

*Note* that the scripts in the postgresinit folder are only run if the DB is initializing for the first time.
Therefore, to reinitialize the data in the container, the `postgresDB` folder and all its contents must be dropped so that they can be recreated automatically.
Dropping the postgerDB folder will **delete all data** existing in the database instance.

Currently, the system uses multiple databases in the same postgresDB instance, but that can be changed by configuring the environment variables associated with each service.
 
 At this point the system uses the following files to seed data:
 
- `postgres_Datasetup.sql`: Setups the data for the promotion service in the default postgres DB
- `userdb.sh`: Creates the `appuser` DB, connects to the newly created DB and calls `userSchema_sql.txt`
- `userSchame_sql.txt`: Called by `userdb.sh` to create `appuser` table and seed it with data. the .txt extension is needed to avoid Postgres trying to run the file and creating the data in the default postgres database
 