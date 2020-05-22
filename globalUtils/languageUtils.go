package globalUtils

type Languages int

const (
	LangEN Languages = iota
	LangES
	LangFR
)

func (l Languages) String() string {
	return [...]string{"English", "Spanish", "French"}[l]
}
