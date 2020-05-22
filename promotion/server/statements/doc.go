/*
Statements package contains service specific error messages and sql statements that will be encountered while running the
promotion service. It is composed of two files:
	- errorstatements.go contains all the error messages. The error messages support multiple languages. Language used
	  depends on the value of the 'language' variable. If not explicitly set, using the SetLanguage function,
	  it defaults to english.
	- sqlstatements.go contains all the sql statements to be run against the DB
*/

package statements
