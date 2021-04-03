/*
User server: Main package to instantiate the user and authentication service.
This package is composed of the following files:
- userServer: Main entry point of service. It starts the service, connects to the database, subscribes to the Broker
	and sets up the authentication wrapper
- handler:Handles all the request made to the service from outside clients
- plugins: Indicates which micro plugins we are using in th app
- tokenService: Handles coding/decoding and validation of JWT tokens
- validation: Handles data validation routines and sends audit messages to the broker for storage
- monitoring: contains utilities related to monitoring the application
*/
package main
