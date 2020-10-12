package statements

import (
	"fmt"
	"goTemp/globalUtils"
	"log"
)

//language: Language in which the application messages will be presented
var language = globalUtils.LangEN

//SetLanguage: Populates the language variable to set in which language the message will be displayed
func SetLanguage(newLanguage globalUtils.Languages) {
	language = newLanguage
}

//UserErr defines User audit specific error messages
type UserErr string

//errTxtEn: Error messages in english map
var errTxtEn = map[string]UserErr{
	"internalError":          "Internal error. Error: %v\n",
	"insertError":            "Unable to create user. Error: %v\n",
	"updateError":            "Unable to update user. Error: %v \n",
	"deleteError":            "Unable to delete user %v. Error: %v\n",
	"deleteRowNotFoundError": "row with id %d not found. Unable to delete the row\n",
	"selectReadError":        "Unable to get rows from the DB. Error: %v \n",
	"selectScanError":        "Unable to read the user rows returned from the Db. Error: %v\n",
	"selectRowReadError":     "Unable to get row from the DB. Error: %v \n",
}

//errTxtES: Error Messages in spanish
var errTxtES = map[string]UserErr{
	"internalError":          "Error interno. Error: %v\n",
	"insertError":            "No se pudo crear la auditoria. Error: %v\n",
	"updateError":            "No se pudo actualizar la auditoria. Error: %v \n",
	"deleteError":            "No se pudo borrar la auditoria %v. Error: %v\n",
	"deleteRowNotFoundError": "La auditoria %d no se pudo encontrar. No se pudo borrar el usuario",
	"selectReadError":        "No su pudo leer datos de la base de datos. Error: %v \n",
	"selectScanError":        "No se pudo leer los datos recibidos de la base de datos. Error: %v\n",
	"selectRowReadError":     "No se pudo leer la auditoria de la base de datos. Error: %v \n",
}

//getSqlTxt: Returns the error message based on the error map key and the selected language
func (ge *UserErr) getSqlTxt(errKey string, myLanguage globalUtils.Languages) string {
	var returnstr string
	switch myLanguage {
	case globalUtils.LangEN:
		returnstr = string(errTxtEn[errKey])
	case globalUtils.LangES:
		returnstr = string(errTxtES[errKey])
	case globalUtils.LangFR:
		log.Fatalf("%s language not implemented for audit", myLanguage)
	default:
		log.Fatalf("%s language not implemented for audit", myLanguage)
	}
	return returnstr
}

/*
The following functions return the appropriate error to the client. it also interpolates the passed arguments
to the returned message
*/

//internalError returns relevant error in the selected language
func (ge *UserErr) internalError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("internalError", language), err)
}

//InsertError returns relevant error in the selected language
func (ge *UserErr) InsertError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("insertError", language), err)
}

//UpdateError returns relevant error in the selected language
func (ge *UserErr) UpdateError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("updateError", language), err)
}

//DeleteError returns relevant error in the selected language
func (ge *UserErr) DeleteError(Id int64, err error) string {
	return fmt.Sprintf(ge.getSqlTxt("deleteError", language), Id, err)
}

//DeleteRowNotFoundError  returns relevant error in the selected language
func (ge *UserErr) DeleteRowNotFoundError(id int64) string {
	return fmt.Sprintf(ge.getSqlTxt("selectRowReadError", language), id)
}

//SelectReadError returns relevant error in the selected language
func (ge *UserErr) SelectReadError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("selectReadError", language), err)
}

//SelectScanError  returns relevant error in the selected language
func (ge *UserErr) SelectScanError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("selectScanError", language), err)
}

//SelectRowReadError returns relevant error in the selected language
func (ge *UserErr) SelectRowReadError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("selectRowReadError", language), err)
}
