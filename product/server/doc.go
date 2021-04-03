/*
Product server: Main package to instantiate the product service.
This package is composed of the following files:
- productServer: Main entry point of service. It starts the service, reads environment variables,
	connects to the database, initializes the broker and sets up the authentication wrapper
- handler:Handles all the requests made to the service from outside clients
- validation: Handles data validation routines and sends audit messages to the broker for storage
- monitoring: contains utilities related to monitoring the application

A note about retrieving data from the database: Since we are dealing with a multi-model No-SQL DB, there is a possibility
that there will be fields in the DB that are not in the protobuf definition used to interact with data. Any such such
additional field will be pulled into the extraFields proto struct and passed along to the client.
As of this writing, this is a read only field and any changes made to it will be discarded before saving the record


*/
package main
