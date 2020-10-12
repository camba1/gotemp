package statements

import (
	"fmt"
	"goTemp/globalUtils"
	"log"
)

//language is the default language in which messages are returned
var language = globalUtils.LangEN

// SetLanguage Overrides the language in which the messages are returned
func SetLanguage(newLanguage globalUtils.Languages) {
	language = newLanguage
}

//UserErr defines user specific error messages
type UserErr string

var errTxtEn = map[string]UserErr{
	"internalError":          "Internal error. Error: %v\n",
	"insertError":            "Unable to create user. Error: %v\n",
	"insertDupEmail":         "Email address already exists in the system\v",
	"updateError":            "Unable to update user. Error: %v \n",
	"updateDupEmail":         "This email address is already associated with another user in the system\v",
	"deleteError":            "Unable to delete user %v. Error: %v\n",
	"deleteRowNotFoundError": "row with id %d not found. Unable to delete the row",
	"selectReadError":        "Unable to get rows from the DB. Error: %v \n",
	"selectScanError":        "Unable to read the user rows returned from the Db. Error: %v\n",
	"selectRowReadError":     "Unable to get row from the DB. Error: %v \n",
	"delUserActive":          "User cannot be deleted because it is active \n",
	"userNotFound":           "User not found\n",
}

var errTxtES = map[string]UserErr{
	"internalError":          "Error interno. Error: %v\n",
	"insertError":            "No se pudo crear el usuario. Error: %v\n",
	"insertDupEmail":         "Correo electornico ya existe en la base de datos\v",
	"updateError":            "No se pudo actualizar el usuario. Error: %v \n",
	"updateDupEmail":         "Este correo electornico ya esta associado con un usuario en el systema\v",
	"deleteError":            "No se pudo borrar el usuario %v. Error: %v\n",
	"deleteRowNotFoundError": "usuario %d no se pudo encontrar. No se pudo borrar el usuario",
	"selectReadError":        "No su pudo leer datos de la base de datos. Error: %v \n",
	"selectScanError":        "No se pudo leer los datos recibidos de la base de datos. Error: %v\n",
	"selectRowReadError":     "No se pudo leer el usuario de la base de datos. Error: %v \n",
	"delUserActive":          "Usario no puede ser borrado porque esta activo \n",
	"userNotFound":           "Usuario no se puedo encontrar\n",
}

//getSqlTxt pull an error message in the correct language
func (ge *UserErr) getSqlTxt(errKey string, myLanguage globalUtils.Languages) string {
	var returnstr string
	switch myLanguage {
	case globalUtils.LangEN:
		returnstr = string(errTxtEn[errKey])
	case globalUtils.LangES:
		returnstr = string(errTxtES[errKey])
	case globalUtils.LangFR:
		log.Fatalf("%s language not implemented for users", myLanguage)
	default:
		log.Fatalf("%s language not implemented for users", myLanguage)
	}
	return returnstr
}

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

//DeleteRowNotFoundError returns relevant error in the selected language
func (ge *UserErr) DeleteRowNotFoundError(id int64) string {
	return fmt.Sprintf(ge.getSqlTxt("selectRowReadError", language), id)
}

//SelectReadError returns relevant error in the selected language
func (ge *UserErr) SelectReadError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("selectReadError", language), err)
}

//SelectScanError returns relevant error in the selected language
func (ge *UserErr) SelectScanError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("selectScanError", language), err)
}

//SelectRowReadError returns relevant error in the selected language
func (ge *UserErr) SelectRowReadError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("selectRowReadError", language), err)
}

//DelUserActive returns relevant error in the selected language
func (ge *UserErr) DelUserActive() string {
	return fmt.Sprintf(ge.getSqlTxt("delUserActive", language))
}

//UpdateDupEmail returns relevant error in the selected language
func (ge *UserErr) UpdateDupEmail() string {
	return fmt.Sprintf(ge.getSqlTxt("updateDupEmail", language))
}

//InsertDupEmail returns relevant error in the selected language
func (ge *UserErr) InsertDupEmail() string {
	return fmt.Sprintf(ge.getSqlTxt("insertDupEmail", language))
}

//UserNotFound returns relevant error in the selected language
func (ge *UserErr) UserNotFound() string {
	return fmt.Sprintf(ge.getSqlTxt("userNotFound", language))
}
