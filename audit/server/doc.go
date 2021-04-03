/*
Audit server: Main package to handle the creation of historical audit data. Anytime a service interacts with data in their
individual services, a message is send to the pub/sub broker  with the globalUtils.AuditTopic topic. This message is
automatically picked up by this service and stored in the time series database for safe guarding.
Service is currently composed of two files:
- auditServer: Main entry point of service. It starts the service, connects to the database and subscribes to the Broker.
- Handler: Retrieves messages from broker and stores them in the time series DB. The messages from the broker are composed of:
	- Header: Contains information indicating, among other things:
		- Service name and service function that performed the operation
		- The type of operation performed (insert, update, etc..)
	- Message payload: Byte slice containing the actual data that was affected (e.g. inserted record)
- monitoring: contains utilities related to monitoring the application
Note that the header of the message is decode and stored in multiple columns,  but the payload is stored directly as is.
*/
package main
