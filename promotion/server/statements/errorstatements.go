package statements

import (
	"fmt"
	"log"
)

type Languages int

const (
	en Languages = iota
	es
	fr
)

func (l Languages) String() string {
	return [...]string{"English", "Spanish", "French"}[l]
}

var language Languages = en

func SetLanguage(newLanguage Languages) {
	language = newLanguage
}

type PromoErr string

var errTxtEn = map[string]PromoErr{
	"internalError":          "Internal error. Error: %v\n",
	"insertError":            "Unable to create promotion. Error: %v\n",
	"UpdateError":            "Unable to update promotion. Error: %v \n",
	"DeleteError":            "Unable to delete promotion %v. Error: %v\n",
	"DeleteRowNotFoundError": "row with id %d not found. Unable to delete the row",
	"SelectReadError":        "Unable to get rows from the DB. Error: %v \n",
	"SelectScanError":        "Unable to read the promotion rows returned from the Db. Error: %v\n",
	"SelectRowReadError":     "Unable to get row from the DB. Error: %v \n",
}

var errTxtES = map[string]PromoErr{
	"internalError":          "Error interno. Error: %v\n",
	"insertError":            "No se pudo crear la promocion. Error: %v\n",
	"UpdateError":            "No se pudo actualizar la promocion. Error: %v \n",
	"DeleteError":            "No se pudo borrar la promocion %v. Error: %v\n",
	"DeleteRowNotFoundError": "Promocion %d no se pudo encontrar. No se pudo borrar la promocion",
	"SelectReadError":        "No su pudo leer datos de la base de datos. Error: %v \n",
	"SelectScanError":        "No se pudo leer los datos recibidos de la base de datos. Error: %v\n",
	"SelectRowReadError":     "No se pudo leer la promocion de la base de datos. Error: %v \n",
}

func (ge *PromoErr) getSqlTxt(errKey string, myLanguage Languages) string {
	var returnstr string
	switch myLanguage {
	case en:
		returnstr = string(errTxtEn[errKey])
	case es:
		returnstr = string(errTxtES[errKey])
	case fr:
		log.Fatalf("%s language not implemented for promotions", myLanguage)
	default:
		log.Fatalf("%s language not implemented for promotions", myLanguage)
	}
	return returnstr
}

func (ge *PromoErr) internalError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("internalError", language), err)
}

func (ge *PromoErr) InsertError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("insertError", language), err)
}

func (ge *PromoErr) UpdateError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("UpdateError", language), err)
}

func (ge *PromoErr) DeleteError(Id int64, err error) string {
	return fmt.Sprintf(ge.getSqlTxt("DeleteError", language), Id, err)
}

func (ge *PromoErr) DeleteRowNotFoundError(id int64) string {
	return fmt.Sprintf(ge.getSqlTxt("SelectRowReadError", language), id)
}

func (ge *PromoErr) SelectReadError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("SelectReadError", language), err)
}

func (ge *PromoErr) SelectScanError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("SelectScanError", language), err)
}

func (ge *PromoErr) SelectRowReadError(err error) string {
	return fmt.Sprintf(ge.getSqlTxt("SelectRowReadError", language), err)
}
