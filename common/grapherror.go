package common

import "fmt"

// GraphErrorResponse represents an error response from the MS Graph API.
//
// https://docs.microsoft.com/en-us/graph/errors
// https://docs.oasis-open.org/odata/odata-json-format/v4.0/os/odata-json-format-v4.0-os.html#_Toc372793091
type GraphErrorResponse struct {
	Error *GraphError `json:"error"`
}

// GraphError is the structure within an MS Graph API error.
//
// https://docs.microsoft.com/en-us/graph/errors
// https://docs.oasis-open.org/odata/odata-json-format/v4.0/os/odata-json-format-v4.0-os.html#_Toc372793091
type GraphError struct {
	Code       string                 `json:"code"`
	Message    string                 `json:"message"`
	Details    *GraphErrorDetails     `json:"details"`
	InnerError map[string]interface{} `json:"innerError"` // The spec says "innererror" but MS returns "innerError"
}

var _ error = GraphError{}

func (err GraphError) Error() string {
	return fmt.Sprintf("%s: %s", err.Code, err.Message)
}

// Is adds support for `error.Is(error, error)`. If `target` is a `GraphError`, the `Code` fields
// are compared.
func (err GraphError) Is(target error) bool {
	e, ok := target.(GraphError)
	if !ok {
		return false
	}

	return e.Code == err.Code
}

// GraphErrorDetails represents the `details` field of a graph API error response.
type GraphErrorDetails struct {
	Code    string  `json:"code"`
	Message string  `json:"message"`
	Target  *string `json:"target"`
}
