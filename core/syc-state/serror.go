/*
core oversees all the shared services, such as logging, errors, messaging, etc.
	The input can be from a file or on the command line. The init file is in JSON format and must be in the same
	directory as the binary using the core services.
		Example: Business Service X (/BusinessService/X)is using core logging, so the init file must be in the

	All errors are in JSON format.
	JSON Fields:
		application             An assigned number or nil.  Nil means that there was no error
		Type             The category of the error message
		ParamCount       The number of parameters have are needed for the message
		ParamDescription Description of the parameters that need to be supplied
		FormattedMessage This is the raw formatted message before the parameters are applied
		Details          This can be used for anything including look up value errors.

Notes:
	Don't argue that a code of nil, is not an error therefore, it should not be in syc-state. This is not college, an
	error is whatever we say it is.
*/
package core

import (
	"fmt"
	_ "github.com/sy-consulting/golang/core"
)

type SYCState struct {
	// Machine-readable state information
	Code             interface{}
	Type             string
	ParamCount       int
	ParamDescription string
	// Human-readable state information
	FormattedStateMessage string
	StateDetails          map[string]string
	SYCError              error
}

//// Error categories
//const (
//	USERERROR          = "User_Error"
//	PROCESSERROR       = "Process_Error"
//	NATSERROR          = "NATS_Error"
//	CONTENTERROR       = "Content_Error"
//	PERMISSIONERROR    = "Permission_Error"
//	CONFIGURATIONISSUE = "Configuration_Issue"
//	APICONTRACTERROR   = "API_Contract_Error"
//	GENERALERROR       = "General_Error"
//	MARKDOWNTITLEBAR   = "| Error Code | Category | Parameter Description | Formatted Error Text |\n|--------|--------|--------|--------|\n"
//	FUNCCOMMENTSHEADER = "\tError Code with requiring parameters:\n"
//	SQLSTATE           = "SQLSTATE"
//	PREFIX             = ""
//	INDENT             = "  "
//)
//
//var (
//	EmptyMap = make(map[string]string)
//)

// Error returns the string representation of the error message.
func (m core.Manager) state() string {
	return fmt.Sprintf("The state is good for %v in %v!", m.Application, m.Environment)
}
