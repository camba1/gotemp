/*
Statements package contains service specific error messages and sql statements that will be encountered while running the
audit service. It is composed of two files:
	- errorstatements.go contains all the error messages. The error messages support multiple languages. Language used
	  depends on the value of the 'language' variable. If not explicitly set, using the SetLanguage function,
	  it defaults to english. Currently the errors are coded directly into string based maps on the package, but they could
	  be loaded from a file or a DB in the future
	- sqlstatements.go contains all the sql statements to be run against the DB. Currently the sql statements are coded
	  directly into constants on the package, but they could be loaded from a file or a DB in the future
*/

package statements
