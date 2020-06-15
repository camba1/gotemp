##ArangoDB Usage

ArangoDB, a multi-model database, stores all master data. this will 
allow fast hierachy lookups and storage. 
There are three folders in this directory and they are all mounted to the
ArangoDB container:
- apps_db_system: Contains any ArangoDB Foxx framework micro services.
- arangodbinit: Contains the DB schemas and data to be used to seed the app. Scripts to create initialization are written in Javascript.
- db: Persists container data to the host. All data in the container is deleted if this folder is deleted.

*Note* that the scripts in the arangodbinit folder are only run if the DB is initializing for the first time.
Therefore, to reinitialize the data in the container, the `db` folder and all its contents must be dropped so that they can be recreated automatically.
Dropping the db folder will **delete all data** existing in the database instance.

Currently, the system uses multiple databases in the same ArangoDB instance, but that can be changed by configuring the environment variables associated with each service.
 
 At this point the system uses the following files to seed data:
 
 - customer_schema: Setups the database, tables and data for the customer service
 - product_schema: Setups the database, tables and data for the product service
 