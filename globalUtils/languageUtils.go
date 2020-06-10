package globalUtils

// Integre based type that represents the language available in the app
type Languages int

const (
	//LangEN: English
	LangEN Languages = iota
	//LangES: Spanish
	LangES
	//LangFR: French
	LangFR
)

//String: Returns the full description of the languages available in the app
func (l Languages) String() string {
	return [...]string{"English", "Spanish", "French"}[l]
}
