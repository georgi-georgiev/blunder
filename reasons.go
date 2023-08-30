package blunder

import "net/http"

type ReasonCode int

const (
	CORS_REQUEST_WITH_XORIGIN     ReasonCode = 110
	ENDPOINT_CONSTRAINT_MISMATCH  ReasonCode = 120
	INVALID                       ReasonCode = 130
	BAD_CONTENT                   ReasonCode = 100
	INVALID_HEADER                ReasonCode = 140
	INVALID_PARAMETER             ReasonCode = 150
	INVALID_QUERY                 ReasonCode = 160
	KEY_EXPIRED                   ReasonCode = 170
	KEY_INVALID                   ReasonCode = 180
	NOT_DOWNLOAD                  ReasonCode = 190
	NOT_UPLOAD                    ReasonCode = 200
	PARSE_ERROR                   ReasonCode = 210
	REQUIRED                      ReasonCode = 220
	TOO_MANY_PARTS                ReasonCode = 230
	UNKNOWN_API                   ReasonCode = 240
	UNSUPPORTED_MEDIA_PROTOCOL    ReasonCode = 250
	UNSUPPORTED_OUTPUT_FORMAT     ReasonCode = 260
	AUTHORIZATION_ERROR           ReasonCode = 300
	EXPIRED                       ReasonCode = 310
	AUTH_REQUIRED                 ReasonCode = 320
	ACCOUNT_DELETED               ReasonCode = 410
	ACCOUNT_DISABLED              ReasonCode = 420
	ACCOUNT_UNVERIFIED            ReasonCode = 430
	ACCESS_NOT_CONFIGURED         ReasonCode = 400
	CONCURRENT_LIMIT_EXCEEDED     ReasonCode = 440
	DAILY_LIMIT_EXCEEDED          ReasonCode = 450
	INSUFFICIENT_AUTHORIZED_PARTY ReasonCode = 460
	LIMIT_EXCEEDED                ReasonCode = 470
	QUOTA_EXCEEDED                ReasonCode = 480
	RATE_LIMIT_EXCEEDED           ReasonCode = 490
	RESPONSE_TOO_LARGE            ReasonCode = 500
	SERVING_LIMIT_EXCEEDED        ReasonCode = 510
	SLL_REQUIRED                  ReasonCode = 520
	USER_RATE_LIMIT_EXCEEDED      ReasonCode = 530
)

func (reasonCode ReasonCode) String() string {
	switch reasonCode {
	case CORS_REQUEST_WITH_XORIGIN:
		return "CORS_REQUEST_WITH_XORIGIN"
	case ENDPOINT_CONSTRAINT_MISMATCH:
		return "ENDPOINT_CONSTRAINT_MISMATCH"
	case INVALID:
		return "INVALID"
	case BAD_CONTENT:
		return "BAD_CONTENT"
	case INVALID_HEADER:
		return "INVALID_HEADER"
	case INVALID_PARAMETER:
		return "INVALID_PARAMETER"
	case INVALID_QUERY:
		return "INVALID_QUERY"
	case KEY_EXPIRED:
		return "KEY_EXPIRED"
	case KEY_INVALID:
		return "KEY_INVALID"
	case NOT_DOWNLOAD:
		return "NOT_DOWNLOAD"
	case NOT_UPLOAD:
		return "NOT_UPLOAD"
	case PARSE_ERROR:
		return "PARSE_ERROR"
	case REQUIRED:
		return "REQUIRED"
	case TOO_MANY_PARTS:
		return "TOO_MANY_PARTS"
	case UNKNOWN_API:
		return "UNKNOWN_API"
	case UNSUPPORTED_MEDIA_PROTOCOL:
		return "UNSUPPORTED_MEDIA_PROTOCOL"
	case UNSUPPORTED_OUTPUT_FORMAT:
		return "UNSUPPORTED_OUTPUT_FORMAT"
	case AUTHORIZATION_ERROR:
		return "AUTHORIZATION_ERROR"
	case EXPIRED:
		return "EXPIRED"
	case AUTH_REQUIRED:
		return "AUTH_REQUIRED"
	case ACCOUNT_DELETED:
		return "ACCOUNT_DELETED"
	case ACCOUNT_DISABLED:
		return "ACCOUNT_DISABLED"
	case ACCOUNT_UNVERIFIED:
		return "ACCOUNT_UNVERIFIED"
	case ACCESS_NOT_CONFIGURED:
		return "ACCESS_NOT_CONFIGURED"
	case CONCURRENT_LIMIT_EXCEEDED:
		return "CONCURRENT_LIMIT_EXCEEDED"
	case DAILY_LIMIT_EXCEEDED:
		return "INVADAILY_LIMIT_EXCEEDEDLID_PARAMETER"
	case INSUFFICIENT_AUTHORIZED_PARTY:
		return "INSUFFICIENT_AUTHORIZED_PARTY"
	case LIMIT_EXCEEDED:
		return "LIMIT_EXCEEDED"
	case QUOTA_EXCEEDED:
		return "QUOTA_EXCEEDED"
	case RATE_LIMIT_EXCEEDED:
		return "RATE_LIMIT_EXCEEDED"
	case RESPONSE_TOO_LARGE:
		return "RESPONSE_TOO_LARGE"
	case SERVING_LIMIT_EXCEEDED:
		return "SERVING_LIMIT_EXCEEDED"
	case SLL_REQUIRED:
		return "SLL_REQUIRED"
	case USER_RATE_LIMIT_EXCEEDED:
		return "USER_RATE_LIMIT_EXCEEDED"
	default:
		return "Unknown reason code"
	}
}

type Reason struct {
	ReasonGroup ReasonGroup
	Message     string
	Tip         string
}

type ReasonGroup struct {
	Status      int
	Title       string
	Description string
	Resolution  string
}

var ReasonGroups map[int]ReasonGroup = map[int]ReasonGroup{
	http.StatusBadRequest: {
		Status:      http.StatusBadRequest,
		Title:       "REQUEST_VALIDATION_FAILURE",
		Description: "The request failed because it contained an invalid value or missing required value. The value could be a parameter value, a header value, or a property value.",
		Resolution:  "Please correct the request as per the error description/details provided in the error response.",
	},
	http.StatusUnauthorized: {
		Status:      http.StatusUnauthorized,
		Title:       "UNAUTHORIZED",
		Description: "The user does not have valid authentication credentials for the target resource.",
		Resolution:  "Check the value of the Authorization HTTP request header.",
	},
	http.StatusForbidden: {
		Status:      http.StatusForbidden,
		Title:       "FORBIDDEN OR ACCESS_DENIED",
		Description: "The requested operation is forbidden and cannot be completed. This may be due to the user not having the necessary permissions for a resource or needing an account of some sort, or attempting a prohibited action (for example, deleting the client or device).",
		Resolution:  "Please ensure that access has been granted to the current user.",
	},
}

var Reasons map[ReasonCode]Reason = map[ReasonCode]Reason{
	BAD_CONTENT: {
		Message:     "The content type of the request data or the content type of a part of a multipart request is not supported.",
		ReasonGroup: ReasonGroups[http.StatusBadRequest],
	},
	CORS_REQUEST_WITH_XORIGIN: {
		Message:     "The CORS request contains an XD3 X-Origin header, which is indicative of a bad CORS request.",
		ReasonGroup: ReasonGroups[http.StatusBadRequest],
	},
	ENDPOINT_CONSTRAINT_MISMATCH: {
		Message:     "The request failed because it did not match the specified API.",
		ReasonGroup: ReasonGroups[http.StatusBadRequest],
		Tip:         "Check the value of the URL path to make sure it is correct.",
	},
	INVALID: {
		Message:     "The request failed because it contained an invalid value.",
		ReasonGroup: ReasonGroups[http.StatusBadRequest],
		Tip:         "The value could be a parameter value, a header value, or a property value.",
	},
	INVALID_HEADER: {
		Message:     "The request failed because it contained an invalid header.",
		ReasonGroup: ReasonGroups[http.StatusBadRequest],
	},
	INVALID_PARAMETER: {
		Message:     "The request failed because it contained an invalid parameter or parameter value.",
		ReasonGroup: ReasonGroups[http.StatusBadRequest],
	},
	INVALID_QUERY: {
		Message:     "The request failed because it contained an invalid value.",
		ReasonGroup: ReasonGroups[http.StatusBadRequest],
	},
	KEY_EXPIRED: {
		Message:     "The API key provided in the request expired, which means the API server is unable to check the quota limit for the application making the request.",
		ReasonGroup: ReasonGroups[http.StatusBadRequest],
		Tip:         "Check the for more information or to obtain a new key.",
	},
	KEY_INVALID: {
		Message:     "The API key provided in the request expired, which means the API server is unable to check the quota limit for the application making the request.",
		ReasonGroup: ReasonGroups[http.StatusBadRequest],
		Tip:         "Check the for more information or to obtain a new key.",
	},
	NOT_DOWNLOAD: {
		Message:     "The request failed because it is not an download request.",
		ReasonGroup: ReasonGroups[http.StatusBadRequest],
		Tip:         "Resend the request to the same path.",
	},
	NOT_UPLOAD: {
		Message:     "The request failed because it is not an upload request.",
		ReasonGroup: ReasonGroups[http.StatusBadRequest],
		Tip:         "Resend the request to the same path.",
	},
	PARSE_ERROR: {
		Message:     "The API server cannot parse the request body.",
		ReasonGroup: ReasonGroups[http.StatusBadRequest],
	},
	REQUIRED: {
		Message:     "The API request is missing required information.",
		ReasonGroup: ReasonGroups[http.StatusBadRequest],
		Tip:         "The required information could be a parameter or resource property.",
	},
	TOO_MANY_PARTS: {
		Message:     "The multipart request failed because it contains too many parts.",
		ReasonGroup: ReasonGroups[http.StatusBadRequest],
	},
	UNKNOWN_API: {
		Message:     "The client is using an unsupported media protocol.",
		ReasonGroup: ReasonGroups[http.StatusBadRequest],
	},
	UNSUPPORTED_MEDIA_PROTOCOL: {
		Message:     "The client is using an unsupported media protocol.",
		ReasonGroup: ReasonGroups[http.StatusBadRequest],
	},
	UNSUPPORTED_OUTPUT_FORMAT: {
		Message:     "The alt parameter value specifies an output format that is not supported for this service.",
		ReasonGroup: ReasonGroups[http.StatusBadRequest],
		Tip:         "Check the value of the alt request parameter.",
	},
	AUTHORIZATION_ERROR: {
		Message:     "The user is not authorized to make the request.",
		ReasonGroup: ReasonGroups[http.StatusUnauthorized],
	},
	EXPIRED: {
		Message:     "Session Expired.",
		ReasonGroup: ReasonGroups[http.StatusUnauthorized],
	},
	AUTH_REQUIRED: {
		Message:     "The user must be logged in to make this API request",
		ReasonGroup: ReasonGroups[http.StatusUnauthorized],
	},
	ACCESS_NOT_CONFIGURED: {
		Message:     "The project has been blocked due to abuse.",
		ReasonGroup: ReasonGroups[http.StatusForbidden],
	},
	ACCOUNT_DELETED: {
		Message:     "The user account associated with the request's authorization credentials has been deleted.",
		ReasonGroup: ReasonGroups[http.StatusForbidden],
	},
	ACCOUNT_DISABLED: {
		Message:     "The user account associated with the request's authorization credentials has been disabled.",
		ReasonGroup: ReasonGroups[http.StatusForbidden],
	},
	ACCOUNT_UNVERIFIED: {
		Message:     "The email address for the user making the request has not been verified.",
		ReasonGroup: ReasonGroups[http.StatusForbidden],
	},
	CONCURRENT_LIMIT_EXCEEDED: {
		Message:     "The request failed because a concurrent usage limit has been reached.",
		ReasonGroup: ReasonGroups[http.StatusForbidden],
	},
	DAILY_LIMIT_EXCEEDED: {
		Message:     "A daily quota limit for the API has been reached.",
		ReasonGroup: ReasonGroups[http.StatusForbidden],
	},
	INSUFFICIENT_AUTHORIZED_PARTY: {
		Message:     "The request cannot be completed for this application.",
		ReasonGroup: ReasonGroups[http.StatusForbidden],
	},
	LIMIT_EXCEEDED: {
		Message:     "The request cannot be completed due to access or rate limitations.",
		ReasonGroup: ReasonGroups[http.StatusForbidden],
	},
	QUOTA_EXCEEDED: {
		Message:     "The requested operation requires more resources than the quota allows.",
		ReasonGroup: ReasonGroups[http.StatusForbidden],
	},
	RATE_LIMIT_EXCEEDED: {
		Message:     "Too many requests have been sent within a given time span.",
		ReasonGroup: ReasonGroups[http.StatusForbidden],
	},
	RESPONSE_TOO_LARGE: {
		Message:     "The requested resource is too large to return.",
		ReasonGroup: ReasonGroups[http.StatusForbidden],
	},
	SERVING_LIMIT_EXCEEDED: {
		Message:     "The overall rate limit specified for the API has already been reached.",
		ReasonGroup: ReasonGroups[http.StatusForbidden],
	},
	SLL_REQUIRED: {
		Message:     "SSL is required to perform this operation.",
		ReasonGroup: ReasonGroups[http.StatusForbidden],
	},
	USER_RATE_LIMIT_EXCEEDED: {
		Message:     "The request failed because a per-user rate limit has been reached.",
		ReasonGroup: ReasonGroups[http.StatusForbidden],
	},
}
