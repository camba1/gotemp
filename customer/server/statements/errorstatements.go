package statements

import (
	"fmt"
	"goTemp/globalUtils"
	"log"
	"time"
)

var language globalUtils.Languages = globalUtils.LangEN

func SetLanguage(newLanguage globalUtils.Languages) {
	language = newLanguage
}

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

func (ge *CustErr) internalError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("internalError", language), err)
}

func (ge *CustErr) InsertError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("insertError", language), err)
}

func (ge *CustErr) UpdateError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("UpdateError", language), err)
}

func (ge *CustErr) DeleteError(Id string, err error) string {
	return fmt.Sprintf(ge.getSqlTxt("DeleteError", language), Id, err)
}

func (ge *CustErr) DeleteRowNotFoundError(id string) string {
	return fmt.Sprintf(ge.getSqlTxt("SelectRowReadError", language), id)
}

func (ge *CustErr) SelectReadError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("SelectReadError", language), err)
}

func (ge *CustErr) SelectScanError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("SelectScanError", language), err)
}

func (ge *CustErr) SelectRowReadError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("SelectRowReadError", language), err)
}

func (ge *CustErr) MissingField(fieldName string) string {
	return fmt.Sprintf(ge.getSqlTxt("MissingField", language), fieldName)
}

func (ge *CustErr) DtInvalidValidityDates(validFrom, validThru time.Time) string {
	dateLayout := globalUtils.DateLayoutISO
	return fmt.Sprintf(ge.getSqlTxt("DtInvalidValidityDates", language), validThru.Format(dateLayout), validFrom.Format(dateLayout))
}

func (ge *CustErr) DelPromoNotInitialState() string {
	return fmt.Sprintf(ge.getSqlTxt("DelPromoNotInitialState", language))
}

//
func (ge *CustErr) UnableToOpenCollection(collectionName string) string {
	return fmt.Sprintf(ge.getSqlTxt("unableToOpenCollection", language), collectionName)
}
