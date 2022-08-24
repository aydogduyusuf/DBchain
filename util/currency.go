package util

const (
	USD = "USD"
	EUR = "EUR"
	TL = "TL"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case EUR, USD, TL:
		return true
	}
	return false
}