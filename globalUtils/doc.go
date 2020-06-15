/*
globalUtils package includes generic utilities that are common to all services. This package is referenced in all services
It is compose of the following files:
- audit: Allows the creation of audit messages to be sent out to the audit service
- broker: Handles the sending and receiving messages to the broker
- languageUtils: keeps track of the languages that have implemented in the application. Currently that just refers to the
	service error messages.
- timeUtils: Useful time conversion methods
- dbConnectUtils: Handles connecting to the databases, including setting up retries in case of connection failure
- authUtils: Useful auth based functions

*/
package globalUtils
