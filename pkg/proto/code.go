package proto

const (
	codeUnknown = -100

	CodeUnknownError = -1
	CodeSuccessReply = 0
	CodeFailedReply  = 1
	CodeSessionError = 400
	Code404          = 404
)

var codeTexts = map[int]string{
	codeUnknown: "Unknown",

	CodeUnknownError: "UnknownError",
	CodeSuccessReply: "SuccessReply",
	CodeFailedReply:  "FailedReply",
	CodeSessionError: "SessionError",
	Code404:          "404Error",
}

func CodeText(code int) string {
	if codeText, ok := codeTexts[code]; ok {
		return codeText
	}

	return codeTexts[codeUnknown]
}
