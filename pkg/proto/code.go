package proto

const (
	codeUnknown = -100

	CodeUnknownError = -1
	CodeSuccess      = 0
	CodeFailed       = 1
	CodeSessionError = 400
	Code404          = 404
)

var codeTexts = map[int]string{
	codeUnknown: "Unknown",

	CodeUnknownError: "UnknownError",
	CodeSuccess:      "Success",
	CodeFailed:       "Failed",
	CodeSessionError: "SessionError",
	Code404:          "404",
}

func CodeText(code int) string {
	if codeText, ok := codeTexts[code]; ok {
		return codeText
	}

	return codeTexts[codeUnknown]
}
