package blunder

import (
	"encoding/json"
)

// A problem details object can have the following members:
type HTTPError struct {

	//(string) A URI reference [RFC3986] that identifies the
	//problem type.  This specification encourages that, when
	//dereferenced, it provide human-readable documentation for the
	//problem type (e.g., using HTML [W3C.REC-html5-20141028]).  When
	//this member is not present, its value is assumed to be
	//"about:blank".
	Type string `json:"type,omitempty" message:"https://example.com/problems/request-parameters-missing"`

	//(strinrg) A short, human-readable summary of the problem
	//type.  It SHOULD NOT change from occurrence to occurrence of the
	//problem, except for purposes of localization (e.g., using
	//proactive content negotiation; see [RFC7231], Section 3.4).
	Title string `json:"title,omitempty" example:"required parameters are missing"`

	//(string) A human-readable explanation specific to this occurrence of the problem.
	Detail string `json:"detail,omitempty" example:"The parameters: limit, date were not provided"`

	ReasonCode int    `json:"reason_code,omitempty" example:"150"`
	Reason     string `json:"reason,omitempty" example:"invalidParameter"`

	Placement  string `json:"placement,omitempty" example:"field"`
	Field      string `json:"field,omitempty" example:"email"`
	Expression string `json:"expression,omitempty" example:"greater"` //greater, not equals, equals, min, max, required, etc.
	Argument   string `json:"argument,omitempty" example:"number"`    //number, string

	Action string `json:"action,omitempty" example:"Resubmit the request with a valid queue manager name or no queue manager name, to retrieve a list of queue managers."`

	StackTrace string `json:"stack_trace,omitempty" example:"asda"`
}

func (e HTTPError) Error() string {
	bytes, err := json.Marshal(e)
	if err != nil {
		return ""
	}

	return string(bytes)
}
