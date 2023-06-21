package base

// type ApiError int64

const (
	ErrUnsupportedCode  string = "unsupported-code"  // if we don't support the supplied currency code (see supported currencies...) https://www.exchangerate-api.com/docs/supported-currencies
	ErrMalformedRequest string = "malformed-request" // when some part of your request doesn't follow the structure shown above.
	ErrInvalidKey       string = "invalid-key"       // when your API key is not valid.
	ErrInactiveAccount  string = "inactive-account"  // if your email address wasn't confirmed.
	ErrApiLimitReached  string = "quota-reached"     // when your account has reached the the number of requests allowed by your plan.
)

var ApiErrorArray = [5]string{ErrUnsupportedCode, ErrMalformedRequest, ErrInvalidKey, ErrInactiveAccount, ErrApiLimitReached}

// const (
// 	NoError ApiError = 0
// 	ErrorMissingApiKey ApiError = 100
// 	ErrorInvalidApiKey ApiError = 110
// 	ErrorMissingFromParam ApiError = 200
// 	ErrorInvalidFromCurrency ApiError = 210
// 	ErrorMissingToParam ApiError = 250
// 	ErrorInvalidToCurrency ApiError = 260
// 	ErrorAmountNonNumeric ApiError = 300
// 	ErrorInvalidAmount ApiError = 310
// 	ErrorZeroAmount ApiError = 320
// 	ErrApiLimitReached ApiError = 400 // API call limit reached for the month
// )

// func (ae ApiError) String() string {
// 	switch ae {
// 	case NoError:
// 		return "No error."
// 	case ErrorMissingApiKey:
// 		return "API key not provided."
// 	case ErrorInvalidApiKey:
// 		return "Invalid API key."
// 	case ErrorMissingFromParam:
// 		return "From currency not provided."
// 	case ErrorInvalidFromCurrency:
// 		return "Invalid from currency."
// 	case ErrorMissingToParam:
// 		return "To currency not provided."
// 	case ErrorInvalidToCurrency:
// 		return "Invalid to currency."
// 	case ErrorAmountNonNumeric:
// 		return "Amount must be numeric."
// 	case ErrorInvalidAmount:
// 		return "Invalid amount."
// 	case ErrorZeroAmount:
// 		return "Amount cannot be zero."
// 	case ErrApiLimitReached:
// 		return "API limit reached for the month."
// 	default:
// 		return "invalid error code"
// 	}
// }
