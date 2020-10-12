package globalUtils

//Languages is an Integer based type that represents the language available in the app
type Languages int

const (
	//LangEN is English
	LangEN Languages = iota
	//LangES is Spanish
	LangES
	//LangFR is French
	LangFR
)

//String returns the full description of the languages available in the app
func (l Languages) String() string {
	return [...]string{"English", "Spanish", "French"}[l]
}
