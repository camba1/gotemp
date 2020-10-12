package statements

import (
	"fmt"
	"goTemp/globalUtils"
	"log"
	"time"
)

//language is the default language in which messages are returned
var language globalUtils.Languages = globalUtils.LangEN

// SetLanguage Overrides the language in which the messages are returned
func SetLanguage(newLanguage globalUtils.Languages) {
	language = newLanguage
}

//CustErr defines customer specific error messages
type CustErr string

var errTxtEn = map[string]CustErr{
	"internalError":          "Internal error. Error: %v\n",
	"insertError":            "Unable to create customer. Error: %v\n",
	"UpdateError":            "Unable to update customer. Error: %v \n",
	"DeleteError":            "Unable to delete customer %s. Error: %v\n",
	"DeleteRowNotFoundError": "Row with id %s not found. Unable to delete the row",
	"SelectReadError":        "Unable to get rows from the DB. Error: %v \n",
	"SelectScanError":        "Unable to read the customer rows returned from the Db. Error: %v\n",
	"SelectRowReadError":     "Unable to get row from the DB. Error: %v \n",
	"MissingField":           "%s must not be empty\n",
	"DtInvalidValidityDates": "The valid thru date (%v) must take place after the valid from date (%v)\n",
	"unableToOpenCollection": "Unable to open collection %s\n",
}

var errTxtES = map[string]CustErr{
	"internalError":          "Error interno. Error: %v\n",
	"insertError":            "No se pudo crear el cliente. Error: %v\n",
	"UpdateError":            "No se pudo actualizar el cliente. Error: %v \n",
	"DeleteError":            "No se pudo borrar el cliente %s. Error: %v\n",
	"DeleteRowNotFoundError": "Cliente %s no se pudo encontrar. No se pudo borrar el cliente",
	"SelectReadError":        "No su pudo leer datos de la base de datos. Error: %v \n",
	"SelectScanError":        "No se pudo leer los datos recibidos de la base de datos. Error: %v\n",
	"SelectRowReadError":     "No se pudo leer el cliente de la base de datos. Error: %v \n",
	"MissingField":           "%s no debe estar vacio\n",
	"DtInvalidValidityDates": "La fecha final (%v) no puede ser menor a la fecha inicial (%v)\n",
	"unableToOpenCollection": "No se pudo abrir la colleccion %s\n",
}

//getSqlTxt pull an error message in the correct language
func (ge *CustErr) getSqlTxt(errKey string, myLanguage globalUtils.Languages) string {
	var returnstr string
	switch myLanguage {
	case globalUtils.LangEN:
		returnstr = string(errTxtEn[errKey])
	case globalUtils.LangES:
		returnstr = string(errTxtES[errKey])
	case globalUtils.LangFR:
		log.Fatalf("%s language not implemented for customers", myLanguage)
	default:
		log.Fatalf("%s language not implemented for customers", myLanguage)
	}
	return returnstr
}

//internalError returns relevant error in the selected language
func (ge *CustErr) internalError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("internalError", language), err)
}

//InsertError returns relevant error in the selected language
func (ge *CustErr) InsertError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("insertError", language), err)
}

//UpdateError returns relevant error in the selected language
func (ge *CustErr) UpdateError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("UpdateError", language), err)
}

//DeleteError returns relevant error in the selected language
func (ge *CustErr) DeleteError(Id string, err error) string {
	return fmt.Sprintf(ge.getSqlTxt("DeleteError", language), Id, err)
}

//DeleteRowNotFoundError returns relevant error in the selected language
func (ge *CustErr) DeleteRowNotFoundError(id string) string {
	return fmt.Sprintf(ge.getSqlTxt("SelectRowReadError", language), id)
}

//SelectReadError returns relevant error in the selected language
func (ge *CustErr) SelectReadError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("SelectReadError", language), err)
}

//SelectScanError returns relevant error in the selected language
func (ge *CustErr) SelectScanError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("SelectScanError", language), err)
}

//SelectRowReadError returns relevant error in the selected language
func (ge *CustErr) SelectRowReadError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("SelectRowReadError", language), err)
}

//MissingField returns relevant error in the selected language
func (ge *CustErr) MissingField(fieldName string) string {
	return fmt.Sprintf(ge.getSqlTxt("MissingField", language), fieldName)
}

//DtInvalidValidityDates returns relevant error in the selected language
func (ge *CustErr) DtInvalidValidityDates(validFrom, validThru time.Time) string {
	dateLayout := globalUtils.DateLayoutISO
	return fmt.Sprintf(ge.getSqlTxt("DtInvalidValidityDates", language), validThru.Format(dateLayout), validFrom.Format(dateLayout))
}

//DelPromoNotInitialState returns relevant error in the selected language
func (ge *CustErr) DelPromoNotInitialState() string {
	return fmt.Sprintf(ge.getSqlTxt("DelPromoNotInitialState", language))
}

//UnableToOpenCollection returns relevant error in the selected language
func (ge *CustErr) UnableToOpenCollection(collectionName string) string {
	return fmt.Sprintf(ge.getSqlTxt("unableToOpenCollection", language), collectionName)
}
