package statements

import (
	"fmt"
	"goTemp/globalUtils"
	"log"
	"time"
)

// language is the default language in which messages are returned
var language globalUtils.Languages = globalUtils.LangEN

// SetLanguage Overrides the language in which the messages are returned
func SetLanguage(newLanguage globalUtils.Languages) {
	language = newLanguage
}

// ProdErr defines product specific error messages
type ProdErr string

// errTxtEn: Error messages in English
var errTxtEn = map[string]ProdErr{
	"internalError":          "Internal error. Error: %v\n",
	"insertError":            "Unable to create product. Error: %v\n",
	"UpdateError":            "Unable to update product. Error: %v \n",
	"DeleteError":            "Unable to delete product %v. Error: %v\n",
	"DeleteRowNotFoundError": "row with id %s not found. Unable to delete the row",
	"SelectReadError":        "Unable to get rows from the DB. Error: %v \n",
	"SelectScanError":        "Unable to read the product rows returned from the Db. Error: %v\n",
	"SelectRowReadError":     "Unable to get row from the DB. Error: %v \n",
	"MissingField":           "%s must not be empty\n",
	"DtInvalidValidityDates": "The valid thru date (%v) must take place after the valid from date (%v)\n",
	"unableToOpenCollection": "Unable to open collection %s\n",
}

// errTxtES: Error messages in Spanish
var errTxtES = map[string]ProdErr{
	"internalError":          "Error interno. Error: %v\n",
	"insertError":            "No se pudo crear el producto. Error: %v\n",
	"UpdateError":            "No se pudo actualizar el producto. Error: %v \n",
	"DeleteError":            "No se pudo borrar el producto %v. Error: %v\n",
	"DeleteRowNotFoundError": "Producto %s no se pudo encontrar. No se pudo borrar la producto",
	"SelectReadError":        "No su pudo leer datos de la base de datos. Error: %v \n",
	"SelectScanError":        "No se pudo leer los datos recibidos de la base de datos. Error: %v\n",
	"SelectRowReadError":     "No se pudo leer el producto de la base de datos. Error: %v \n",
	"MissingField":           "%s no debe estar vacio\n",
	"DtInvalidValidityDates": "La fecha final (%v) no puede ser menor a la fecha inicial (%v)\n",
	"unableToOpenCollection": "No se pudo abrir la colleccion %s\n",
}

// getSqlTxt: Returns an error message text in a given language
func (ge *ProdErr) getSqlTxt(errKey string, myLanguage globalUtils.Languages) string {
	var returnstr string
	switch myLanguage {
	case globalUtils.LangEN:
		returnstr = string(errTxtEn[errKey])
	case globalUtils.LangES:
		returnstr = string(errTxtES[errKey])
	case globalUtils.LangFR:
		log.Fatalf("%s language not implemented for products", myLanguage)
	default:
		log.Fatalf("%s language not implemented for products", myLanguage)
	}
	return returnstr
}

// internalError returns relevant error in the selected language
func (ge *ProdErr) internalError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("internalError", language), err)
}

// InsertError returns relevant error in the selected language
func (ge *ProdErr) InsertError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("insertError", language), err)
}

// UpdateError returns relevant error in the selected language
func (ge *ProdErr) UpdateError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("UpdateError", language), err)
}

// DeleteError returns relevant error in the selected language
func (ge *ProdErr) DeleteError(Id string, err error) string {
	return fmt.Sprintf(ge.getSqlTxt("DeleteError", language), Id, err)
}

// DeleteRowNotFoundError returns relevant error in the selected language
func (ge *ProdErr) DeleteRowNotFoundError(id int64) string {
	return fmt.Sprintf(ge.getSqlTxt("SelectRowReadError", language), id)
}

// SelectReadError returns relevant error in the selected language
func (ge *ProdErr) SelectReadError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("SelectReadError", language), err)
}

// SelectScanError returns relevant error in the selected language
func (ge *ProdErr) SelectScanError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("SelectScanError", language), err)
}

// SelectRowReadError returns relevant error in the selected language
func (ge *ProdErr) SelectRowReadError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("SelectRowReadError", language), err)
}

// MissingField returns relevant error in the selected language
func (ge *ProdErr) MissingField(fieldName string) string {
	return fmt.Sprintf(ge.getSqlTxt("MissingField", language), fieldName)
}

// DtInvalidValidityDates returns relevant error in the selected language
func (ge *ProdErr) DtInvalidValidityDates(validFrom, validThru time.Time) string {
	dateLayout := globalUtils.DateLayoutISO
	return fmt.Sprintf(ge.getSqlTxt("DtInvalidValidityDates", language), validThru.Format(dateLayout), validFrom.Format(dateLayout))
}

// DelPromoNotInitialState returns relevant error in the selected language
func (ge *ProdErr) DelPromoNotInitialState() string {
	return fmt.Sprintf(ge.getSqlTxt("DelPromoNotInitialState", language))
}

// UnableToOpenCollection returns relevant error in the selected language
func (ge *ProdErr) UnableToOpenCollection(collectionName string) string {
	return fmt.Sprintf(ge.getSqlTxt("unableToOpenCollection", language), collectionName)
}
