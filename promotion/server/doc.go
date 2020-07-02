/*
Promotion server: main package to instantiate the promotion service.
This package is composed of the following files:
- promotionServer: Main entry point of service. It starts the service, reads environment variables,
	connects to the database, initializes the broker, initializes the cache and sets up the authentication wrapper
- handler:Handles all the requests made to the service from outside clients
- lookup: Handles lookup of values in other services. As of this writing this is limited to the customer name pulled
	from the customer service and caches it as needed in the redis store
- plugins; Imports the go-micor plugins used by the service.
- validation: Handles data validation routines and sends audit messages to the broker for storage


*/
package main
