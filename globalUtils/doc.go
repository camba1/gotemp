/*
globalUtils package includes generic utilities that are common to all services. This package is referenced in all services
It is compose of the following files:
- audit: Allows the creation of audit messages to be sent out to the audit service
- AuthUtils: Useful auth based functions
- broker: Handles the sending and receiving messages to the broker
- cache: Useful functions to cache information in the store
- dbConnectUtils: Handles connecting to the databases, including setting up retries in case of connection failure
- languageUtils: keeps track of the languages that have implemented in the application. Currently that just refers to the
	service error messages.
- timeUtils: Useful time conversion methods
*/
package globalUtils
