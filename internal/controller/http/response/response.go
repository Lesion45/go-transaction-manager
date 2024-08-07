package response

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

const (
	StatusOK    = "Ok"
	StatusError = "Error"
)

func OK() Response {
	return Response{Status: StatusOK}
}

func Error(msg string) Response {
	return Response{Status: StatusError,
		Error: msg}
}

func ValidataionError(errs validator.ValidationErrors) Response {
	var errsMsgs []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errsMsgs = append(errsMsgs, fmt.Sprintf("field %s is a required field", err.Field()))

		default:
			errsMsgs = append(errsMsgs, fmt.Sprintf("field %s is not valid"))
		}
	}
	return Response{
		Status: StatusError,
		Error:  strings.Join(errsMsgs, ", "),
	}
}

type Balance struct {
	Balance float64 `json:"balance"`
}

// TODO: DELETE IF WON'T BE USED
type ReservedBalance struct {
	ReservedBalance float64 `json:"reserved_balance"`
}

type MonthlyReport struct {
	ServiceName string  `json:"service_name"`
	Margin      float64 `json:"margin"`
}
