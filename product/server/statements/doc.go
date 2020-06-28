/*
Statements package contains service specific error messages and AQL statements that will be encountered while running the
product service. It is composed of two files:
	- errorstatements.go contains all the error messages. The error messages support multiple languages. Language used
	  depends on the value of the 'language' variable. If not explicitly set, using the SetLanguage function,
	  it defaults to english.
	- sqlstatements.go contains all the AQL statements to be run against the DB. Note that for most simple CRUD operations
	  we simply use the go-driver's default methods to interact with the DB.
*/

package statements
