package errnum

const (
	// ErrCodeRequestJsonDecode - 410: Request JSON cannot be decoded.
	ErrCodeRequestJsonDecode = iota + 410
	// ErrCodeRequestBodyNil - 411: Request body is nil.
	ErrCodeRequestBodyNil
	// ErrCodeRequestBodyParseFail - 412: Request body cannot be parsed.
	ErrCodeRequestBodyParseFail
	// ErrCodeEntryNotFound - 412: Generic error for when the required entry cannot be found in the database.
	ErrCodeEntryNotFound
)

const (
	// ErrCodeQueryFail - 510: General error, SQL query failed.
	ErrCodeQueryFail = iota + 510
	// ErrCodePrepareFail - 511: Cannot prepare SQL statement.
	ErrCodePrepareFail
	// ErrCodeExecFail - 512: General error, SQL execution failed.
	ErrCodeExecFail
)

var (
	// Errs map ensures that there are no duplicate error codes in this service.
	// This map contains the default error message for each error code, you don't have to use it, but each error
	// code must be registered in this map.
	Errs = map[int]string{
		ErrCodeRequestJsonDecode:    "json decode error",
		ErrCodeRequestBodyNil:       "request body is nil",
		ErrCodeRequestBodyParseFail: "request body parse fail",
		ErrCodeEntryNotFound:        "entry not found",

		ErrCodeQueryFail:   "SQL query failed",
		ErrCodePrepareFail: "prepare fail",
		ErrCodeExecFail:    "SQL execution failed",
	}
)
