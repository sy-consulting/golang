/*
	This is SY-HOLDING's error library. It provides the all the defined errors for an application. If you think of a way to improve it, present the idea.

RESTRICTIONS:
	Do not use logging in this package. It makes the output of the log confusing.
	All errors must be output to StdErr

LICENSE:
	This code is licensed under MIT Licence (https://mit-license.org/)
*/
package cSystemError

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

//goland:noinspection ALL
const (
	// Formatted Lines
	ERROR_LOG_PREFIX = "%v.%v.%v:"
	ERROR_DETAILS    = "ERROR DETAILS Sequence: %v Info: %v"
	// Error Codes
	ITEM_ALREADY_EXISTS         = 100000
	ITEM_ALREADY_EXISTS_MESSAGE = "%v already exists"
	ITEM_NOT_FOUND              = 100099
	ITEM_NOT_FOUND_MESSAGE      = "%v was/were not found"
	NOT_POPULATED               = 100100
	NOT_POPULATED_MESSAGE       = "%v must be populated"
	NOT_AUTHORIZED              = 100100
	NOT_AUTHORIZED_MESSAGE      = "Your roles %v are not authorized to %v"
	//100100: {100100, USERERROR, 2, "List of users roles, Requested action", ": Your roles %v are not authorized to %v", EmptyMap, "", nil},
	//100200: {100200, PROCESSERROR, 0, "None", ": Row has been updated since reading it, re-read the row", EmptyMap, "", nil},
	//100500: {100500, PROCESSERROR, 1, "Thing being changed", ": You are making changes to a canceled or completed %v", EmptyMap, "", nil},
	//100600: {100600, PROCESSERROR, 1, "Item is not active", ": You are making changes to an inactive %v", EmptyMap, "", nil},
	//101010: {101010, PROCESSERROR, 1, "Service Name", ": %v timed out", EmptyMap, "", nil},
	//109999: {109999, USERERROR, 1, "Item name", ": %v was/were not found", EmptyMap, "", nil},
	//199999: {199999, GENERALERROR, 1, "Error Details", ": An error has occurred that is not expected. See Log! %v", EmptyMap, "", nil},
	//// ======================================================================
	//// Errors where the Back End can take action or the system needs to panic
	//200100: {200100, PROCESSERROR, 0, "None", ": Table doesn't exist", EmptyMap, "", nil},
	//200200: {200200, PROCESSERROR, 2, "Parameter name, Data type of parameter", ": %v must be of type %v", EmptyMap, "", nil},
	//200250: {200250, PROCESSERROR, 3, "Parameter name, Parameter value, List of values allowed", ": %v (%v) must contain one of these values: %v",
	//EmptyMap, "", nil},
	//200260: {200260, PROCESSERROR, 3, "Other parameter name, Parameter name, Parameter value", ": %v must be provided when %v is set to (%v)",
	//EmptyMap, "", nil},
	//200510: {200510, PROCESSERROR, 3, "Parameter name, Field name, Field value", ": %v can't be updated because %v is set to %v", EmptyMap, "",
	//nil},
	//200511: {200511, PROCESSERROR, 2, "Parameter name, Another parameter name", ": %v and %v must both be populated or null", EmptyMap, "", nil},
	//200512: {200512, PROCESSERROR, 2, "Parameter name, Another parameter name", ": %v and %v must both be populated", EmptyMap, "", nil},
	//200514: {200514, PROCESSERROR, 3, "Parameter name, Another parameter name, Another parameter name", ": %v, %v and %v must all be populated",
	//EmptyMap, "", nil},
	//200515: {200515, PROCESSERROR, 2, "Parameter name, Another parameter name", ": %v must be empty when %v is populated", EmptyMap, "", nil},
	//200600: {200600, PROCESSERROR, 1, "Info returned from HTTP/HTTPS Request", ": Bad HTTP/HTTPS Request - %v", EmptyMap, "", nil},
	//200700: {200700, PROCESSERROR, 1, "Environment Name", ": The API you are calling is not available in this environment (%v)", EmptyMap, "",
	//nil},
	//200800: {200800, PROCESSERROR, 0, "None", ": QuickSight error - see Details", EmptyMap, "", nil},
	//200900: {200900, PROCESSERROR, 0, "None", ": Database constraint error - see Details", EmptyMap, "", nil},
	//200999: {200999, PROCESSERROR, 0, "None", ": SQL error - see Details", EmptyMap, "", nil},
	//201999: {201999, PROCESSERROR, 0, "None", ": Cognito error - see Details", EmptyMap, "", nil},
	//203000: {203000, PROCESSERROR, 0, "None", ": The number of parameters provided for the error message does not match the required number",
	//EmptyMap, "", nil},
	//203050: {203050, PROCESSERROR, 2, "Name, Application/Package name", ": Number of parameters defined in the %v is not support by %v", EmptyMap,
	//"", nil},
	//203060: {203060, PROCESSERROR, 2, "Provided parameter count, Expected parameter count",
	//": Number of parameters provided (%v) doesn't match the number expected (%v)", EmptyMap, "", nil},
	//205000: {205000, PROCESSERROR, 0, "None", ": AWS SES error - see details in retPack", EmptyMap, "", nil},
	//205005: {205005, PROCESSERROR, 0, "None", ": AWS STS error - see details in retPack", EmptyMap, "", nil},
	//206000: {206000, NATSERROR, 0, "None", ": Jetstream is not enabled", EmptyMap, "", nil},
	//206050: {206050, NATSERROR, 2, "Subscription Name, Subject", ": (%v) is an invalid subscription. Subject: %v", EmptyMap, "", nil},
	//206100: {206100, NATSERROR, 1, "Key name", ": Upper or lower case %v key is missing", EmptyMap, "", nil},
	//206105: {206105, NATSERROR, 1, "Key name", ": Upper or lower case %v keys value is missing", EmptyMap, "", nil},
	//206200: {206200, NATSERROR, 1, "List of required parameters",
	//": Message doesn't match signature. Sender must provide the following parameter names: %v", EmptyMap, "", nil},
	//206300: {206300, NATSERROR, 0, "None", ": Stream pointer is nil. Must be a validate pointer to a stream.", EmptyMap, "", nil},
	//206400: {206400, NATSERROR, 1, "Stream Name", ": Stream creation encountered an error that is not expected. Stream Name: %v", EmptyMap, "",
	//nil},
	//206600: {206600, NATSERROR, 2, "Stream Name, Consumer Name", ": Consumer creation encountered an error that is not expected. " +
	//"Stream Name: %v Consumer Name: %v", EmptyMap, "", nil},
	//206700: {206700, NATSERROR, 2, "Stream Name, Consumer Subject Filter",
	//": The consumer subject filter must be a subset of the stream subject. " +
	//"Stream Name: %v Consumer Subject Filter: %v", EmptyMap, "", nil},
	//207000: {207000, CONTENTERROR, 2, "Field name, Field value", ": %v (%v) is not numeric", EmptyMap, "", nil},
	//207005: {207005, CONTENTERROR, 2, "Field name, Minimal length", ": %v must have a value greater than %v", EmptyMap, "", nil},
	//207010: {207010, CONTENTERROR, 2, "Field name, Field value", ": %v (%v) is not a string", EmptyMap, "", nil},
	//207020: {207020, CONTENTERROR, 2, "Field name, Field value", ": %v (%v) is not a float", EmptyMap, "", nil},
	//207030: {207030, CONTENTERROR, 2, "Field name, Field value", ": %v (%v) is not a array", EmptyMap, "", nil},
	//207040: {207040, CONTENTERROR, 2, "Field name, Field value", ": %v (%v) is not a json string", EmptyMap, "", nil},
	//207050: {207050, CONTENTERROR, 2, "Field name, Field value", ": %v (%v) is not a valid email address", EmptyMap, "", nil},
	//207060: {207060, CONTENTERROR, 2, "Field name, Field value", ": %v (%v) contains special characters which are not allowed", EmptyMap, "",
	//nil},
	//207065: {207065, CONTENTERROR, 2, "Field name, Field value", ": %v (%v) contains special characters other than underscore", EmptyMap, "",
	//nil},
	//207070: {207070, CONTENTERROR, 2, "Field name, Field value", ": %v (%v) is not a valid date", EmptyMap, "", nil},
	//207080: {207080, CONTENTERROR, 2, "Field name, Field value", ": %v (%v) is not a valid timestamp. Format's are UTC, GMT or Zulu", EmptyMap,
	//"", nil},
	//207090: {207090, CONTENTERROR, 6, "Field name, Field value, 'small' or 'large', 'Min' or 'Max', expected size, actual size",
	//": %v (%v) is too %v. %v size: %v Actual size: %v", EmptyMap, "", nil},
	//207095: {207095, CONTENTERROR, 4, "Field name, Field value, greater than value, less than value",
	//": %v (%v) must be greater than %v and less than %v", EmptyMap, "", nil},
	//207100: {207100, CONTENTERROR, 2, "Parameter name, Data Structure Type", ": %v couldn't be converted to an %v - JSON conversion error",
	//EmptyMap, "", nil},
	//207105: {207105, CONTENTERROR, 2, "Data Structure Name, Data Structure Type",
	//": %v (%v) couldn't be converted to JSON - JSON conversion error", EmptyMap, "", nil},
	//207110: {207110, CONTENTERROR, 1, "Parameter name", ": %v couldn't be parsed - Invalid JSON error", EmptyMap, "", nil},
	//207111: {207111, CONTENTERROR, 2, "Parameter name, Application/Package name", ": %v couldn't be converted to a map/keyed array - %v",
	//EmptyMap, "", nil},
	//207200: {207200, CONTENTERROR, 2, "Parameter name, Data Structure Type", ": %v couldn't be converted to an %v", EmptyMap, "", nil},
	//208000: {208000, CONTENTERROR, 0, "None", ": Column must have a non-null value. Details: ", EmptyMap, "", nil},
	//208010: {208010, CONTENTERROR, 0, "None", ": Column data type is not support or invalid. Details: ", EmptyMap, "", nil},
	//208110: {208110, CONTENTERROR, 2, "Thing being changed, System Id for the thing",
	//": No update is needed. No fields where changed for %v with id %v", EmptyMap, "", nil},
	//208120: {208120, CONTENTERROR, 3, "JSON array name, Thing being changed, System Id for the thing", ": The %v was empty for %v with id %v",
	//EmptyMap, "", nil},
	//208200: {208200, CONTENTERROR, 1, "Error message number", ": %v error message is missing from sError package", EmptyMap, "", nil},
	//208300: {208300, PERMISSIONERROR, 0, "None", ": iss (Issuer) is not valid", EmptyMap, "", nil},
	//208310: {208310, PERMISSIONERROR, 1, "Subject", ": sub (Subject: %v) was not present", EmptyMap, "", nil},
	//208320: {208320, PERMISSIONERROR, 0, "None", ": token_use is not valid", EmptyMap, "", nil},
	//208330: {208330, PERMISSIONERROR, 0, "None", ": client id is not valid", EmptyMap, "", nil},
	//208340: {208340, PERMISSIONERROR, 0, "None", ": client id is not valid for this application", EmptyMap, "", nil},
	//208350: {208350, PERMISSIONERROR, 0, "None", ": Token is expired", EmptyMap, "", nil},
	//208355: {208355, PERMISSIONERROR, 0, "None", ": Token is invalid", EmptyMap, "", nil},
	//208356: {208356, PERMISSIONERROR, 0, "None", ": Token contains an invalid number of segments", EmptyMap, "", nil},
	//208360: {208360, PERMISSIONERROR, 1, "Claim names", ": These claims are invalid: %v", EmptyMap, "", nil},
	//208370: {208370, PERMISSIONERROR, 0, "None", ": Required claim(s) is/are missing", EmptyMap, "", nil},
	//209000: {209000, CONFIGURATIONISSUE, 0, "None", ": .env files are missing", EmptyMap, "", nil},
	//209010: {209010, CONFIGURATIONISSUE, 2, "File name, Message returned from Open", ": %v file was not found. Message return: %v", EmptyMap, "",
	//nil},
	//209100: {209100, CONFIGURATIONISSUE, 1, "Environment name", ": environment variable is missing (%v)", EmptyMap, "", nil},
	//209110: {209110, CONFIGURATIONISSUE, 1, "Environment name", ": environment value (%v) is invalid", EmptyMap, "", nil},
	//209200: {209200, CONFIGURATIONISSUE, 3, "Database name, Database driver name, Port value",
	//": Unable to connect to database %v using driver %v on port %v", EmptyMap, "", nil},
	//209210: {209210, CONFIGURATIONISSUE, 0, "None", ": Unable to pass database authentication", EmptyMap, "", nil},
	//209220: {209220, CONFIGURATIONISSUE, 1, "SSL Mode", ": Only disable, allow, prefer and required are supported.", EmptyMap, "", nil},
	//209230: {209230, CONFIGURATIONISSUE, 1, "Connection Type", ": Only single or pool are supported.", EmptyMap, "", nil},
	//209299: {209299, CONFIGURATIONISSUE, 0, "None", ": No database connection has been established", EmptyMap, "", nil},
	//209398: {209398, CONFIGURATIONISSUE, 0, "None", ": no nkey seed found", EmptyMap, "", nil},
	//209499: {209499, CONFIGURATIONISSUE, 0, "None", ": No nats connection has been established", EmptyMap, "", nil},
	//209500: {209500, CONFIGURATIONISSUE, 0, "None", ": Unexpected signing method", EmptyMap, "", nil},
	//209510: {209510, CONFIGURATIONISSUE, 0, "None", ": Kid header not found", EmptyMap, "", nil},
	//209520: {209520, CONFIGURATIONISSUE, 1, "Kid", ": key (%v) was not found in token", EmptyMap, "", nil},
	//209521: {209521, CONFIGURATIONISSUE, 1, "Kid", ": Kid (%v) was not found in public key set", EmptyMap, "", nil},
	//210030: {210030, CONFIGURATIONISSUE, 2, "Region, Environment", ": Failed to fetch remote JWK (status = 404) for %v region %v environment",
	//EmptyMap, "", nil},
	//210090: {210090, CONFIGURATIONISSUE, 1, "Parameter name", ": URL is missing (%v)", EmptyMap, "", nil},
	//210098: {210098, CONFIGURATIONISSUE, 1, "Parameter name", ": Start up parameter is out of value range (%v)", EmptyMap, "", nil},
	//210099: {210099, CONFIGURATIONISSUE, 1, "Parameter name", ": Start up parameter is missing (%v)", EmptyMap, "", nil},
	//210100: {210100, APICONTRACTERROR, 1, "List of required parameters",
	//": Call doesn't match API signature. Caller must provide the following parameter names: %v", EmptyMap, "", nil},
	//210200: {210200, GENERALERROR, 0, "None", ": Postgres error has occurred that is not expected.", EmptyMap, "", nil},
	//210299: {210299, GENERALERROR, 0, "None", ": Postgres is not responding over TCP. Container may not be running.", EmptyMap, "", nil},
	//210399: {210399, GENERALERROR, 0, "None", ": AWS session error has occurred that is not expected", EmptyMap, "", nil},
	//210499: {210499, GENERALERROR, 0, "None", ": Synadia error has occurred that is not expected.", EmptyMap, "", nil},
	//210599: {210599, GENERALERROR, 0, "None", ": Business Service error has occurred that is not expected.", EmptyMap, "", nil},
	//}

)

type Error struct {
	ErrorCode    uint32         `json:"ErrorCode"`
	ErrorMsg     string         `json:"ErrorMessage"`
	ErrorDetails map[int]string `json:"ErrorDetails"`
	FileName     string         `json:"FileName"`
	LineNumber   uint           `json:"LineNumber"`
}

type SystemError struct {
	application string
	environment string
	internalIP  string
	errorLogPtr *log.Logger
	errorInfo   Error
}

func New(application, environment, internalIP string) (SystemErrorPtr *SystemError) {
	SystemErrorPtr = &SystemError{
		application: application,
		environment: environment,
		internalIP:  internalIP,
		errorLogPtr: initialize(application, environment, internalIP),
	}

	return
}

func initialize(application, environment, internalIP string) *log.Logger {

	return log.New(os.Stderr, fmt.Sprintf(ERROR_LOG_PREFIX, application, environment, internalIP), log.Lmsgprefix|log.LstdFlags|log.Lmicroseconds|log.LUTC)
}

//func (se SystemError) ErrItemAlreadyExists_100000(message string, details map[int]string, err error) (myError *Error) {
//	_, filename, line, _ := runtime.Caller(1)
//
//	myError = &Error{
//		ErrorCode:    ITEM_ALREADY_EXISTS,
//		ErrorMsg:     message,
//		ErrorDetails: details,
//		FileName:     filename,
//		LineNumber:   uint(line),
//	}
//
//	return
//}

func (se SystemError) ErrItemNotPopulated_100100(itemName string) (myError *Error) {
	_, filename, line, _ := runtime.Caller(1)

	myError = &Error{
		ErrorCode:  NOT_POPULATED,
		ErrorMsg:   fmt.Sprintf(NOT_POPULATED_MESSAGE, itemName),
		FileName:   filename,
		LineNumber: uint(line),
	}

	return
}

//func formatErrorMessage(message, filename string, details map[int]string, line, errorCode int) (formattedErrorMessage string) {
//
//	formattedErrorMessage = fmt.Sprintf(FORMATTED_ERROR_MESSAGE, errorCode, message, formatErrorDetails(details), filename, line)
//
//	return
//}

//func formatErrorDetails(details map[int]string) (errDetails string) {
//	for seq, info := range details {
//		errDetails = errDetails + fmt.Sprintf(ERROR_DETAILS, strconv.Itoa(seq), info)
//	}
//
//	return
//}
