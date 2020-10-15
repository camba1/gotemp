package statements

import (
	"fmt"
	"goTemp/globalUtils"
	"log"
	"time"
)

// language is the default language in which messages are returned
var language globalUtils.Languages = globalUtils.LangEN

//  SetLanguage Overrides the language in which the messages are returned
func SetLanguage(newLanguage globalUtils.Languages) {
	language = newLanguage
}

// PromoErr defines promotion specific error messages
type PromoErr string

var errTxtEn = map[string]PromoErr{
	"internalError":             "Internal error. Error: %v\n",
	"insertError":               "Unable to create promotion. Error: %v\n",
	"UpdateError":               "Unable to update promotion. Error: %v \n",
	"DeleteError":               "Unable to delete promotion %v. Error: %v\n",
	"DeleteRowNotFoundError":    "row with id %d not found. Unable to delete the row",
	"SelectReadError":           "Unable to get rows from the DB. Error: %v \n",
	"SelectScanError":           "Unable to read the promotion rows returned from the Db. Error: %v\n",
	"SelectRowReadError":        "Unable to get row from the DB. Error: %v \n",
	"MissingField":              "%s must not be empty\n",
	"DtInvalidValidityDates":    "The valid thru date (%v) must take place after the valid from date (%v)\n",
	"DelPromoNotInitialState":   "Promotion cannot be deleted because it is not in initial state \n",
	"cacheCustomerNameNotFound": "Customer name not found in cache, getting it from service. Error: %v\n",
	"customerNameNotFound":      "Unable to get customer %s from cache or customer service. Error: %v\n",
}

var errTxtES = map[string]PromoErr{
	"internalError":             "Error interno. Error: %v\n",
	"insertError":               "No se pudo crear la promocion. Error: %v\n",
	"UpdateError":               "No se pudo actualizar la promocion. Error: %v \n",
	"DeleteError":               "No se pudo borrar la promocion %v. Error: %v\n",
	"DeleteRowNotFoundError":    "Promocion %d no se pudo encontrar. No se pudo borrar la promocion",
	"SelectReadError":           "No su pudo leer datos de la base de datos. Error: %v \n",
	"SelectScanError":           "No se pudo leer los datos recibidos de la base de datos. Error: %v\n",
	"SelectRowReadError":        "No se pudo leer la promocion de la base de datos. Error: %v \n",
	"MissingField":              "%s no debe estar vacio\n",
	"DtInvalidValidityDates":    "La fecha final (%v) no puede ser menor a la fecha inicial (%v)\n",
	"DelPromoNotInitialState":   "Promocion no puede ser borrada porque no esta en estado inicial \n",
	"cacheCustomerNameNotFound": "El nombre del cliente no se encontro en el cache. Obteninedolo del servicio. Error: %v\n",
	"customerNameNotFound":      "El nombre del cliente %s no se encontro en el cache o en el servicio. Error: %v\n",
}

// getSqlTxt pull an error message in the correct language
func (ge *PromoErr) getSqlTxt(errKey string, myLanguage globalUtils.Languages) string {
	var returnstr string
	switch myLanguage {
	case globalUtils.LangEN:
		returnstr = string(errTxtEn[errKey])
	case globalUtils.LangES:
		returnstr = string(errTxtES[errKey])
	case globalUtils.LangFR:
		log.Fatalf("%s language not implemented for promotions", myLanguage)
	default:
		log.Fatalf("%s language not implemented for promotions", myLanguage)
	}
	return returnstr
}

// internalError returns relevant error in the selected language
func (ge *PromoErr) internalError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("internalError", language), err)
}

// InsertError returns relevant error in the selected language
func (ge *PromoErr) InsertError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("insertError", language), err)
}

// UpdateError returns relevant error in the selected language
func (ge *PromoErr) UpdateError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("UpdateError", language), err)
}

// DeleteError returns relevant error in the selected language
func (ge *PromoErr) DeleteError(Id int64, err error) string {
	return fmt.Sprintf(ge.getSqlTxt("DeleteError", language), Id, err)
}

// DeleteRowNotFoundError returns relevant error in the selected language
func (ge *PromoErr) DeleteRowNotFoundError(id int64) string {
	return fmt.Sprintf(ge.getSqlTxt("SelectRowReadError", language), id)
}

// SelectReadError returns relevant error in the selected language
func (ge *PromoErr) SelectReadError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("SelectReadError", language), err)
}

// SelectScanError returns relevant error in the selected language
func (ge *PromoErr) SelectScanError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("SelectScanError", language), err)
}

// SelectRowReadError returns relevant error in the selected language
func (ge *PromoErr) SelectRowReadError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("SelectRowReadError", language), err)
}

// MissingField returns relevant error in the selected language
func (ge *PromoErr) MissingField(fieldName string) string {
	return fmt.Sprintf(ge.getSqlTxt("MissingField", language), fieldName)
}

// DtInvalidValidityDates returns relevant error in the selected language
func (ge *PromoErr) DtInvalidValidityDates(validFrom, validThru time.Time) string {
	dateLayout := globalUtils.DateLayoutISO
	return fmt.Sprintf(ge.getSqlTxt("DtInvalidValidityDates", language), validThru.Format(dateLayout), validFrom.Format(dateLayout))
}

// DelPromoNotInitialState returns relevant error in the selected language
func (ge *PromoErr) DelPromoNotInitialState() string {
	return fmt.Sprintf(ge.getSqlTxt("DelPromoNotInitialState", language))
}

// CacheCustomerNameNotFound returns relevant error in the selected language
func (ge *PromoErr) CacheCustomerNameNotFound(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("cacheCustomerNameNotFound", language), err)
}

// CustomerNameNotFound returns relevant error in the selected language
func (ge *PromoErr) CustomerNameNotFound(customerId string, err error) string {
	return fmt.Sprintf(ge.getSqlTxt("customerNameNotFound", language), customerId, err)
}
