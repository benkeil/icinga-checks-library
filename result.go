package icinga

import (
	"fmt"
	"os"
)

type (
	// Result interface for service check results
	Result interface {
		Name() string
		Status() Status
		Message() string
		Exit()
	}

	resultImpl struct {
		name    string
		status  Status
		message string
	}
)

const (
	// DefaultSuccessMessage the default message if the check was successful
	DefaultSuccessMessage string = "everything ok"
)

// NewResult creates a new instance of Result
func NewResult(name string, status Status, message string) Result {
	return &resultImpl{name, status, message}
}

// NewResultOk creates a new instance of Result and set result to ServiceStateOk
func NewResultOk(name string) Result {
	return &resultImpl{name, ServiceStatusOk, DefaultSuccessMessage}
}

// NewResultOkMessage creates a new instance of Result and set result to ServiceStateOk
func NewResultOkMessage(name string, message string) Result {
	return &resultImpl{name, ServiceStatusOk, message}
}

// NewResultUnknownMessage creates a new instance of Result and set result to ServiceStateOk
func NewResultUnknownMessage(name string, message string) Result {
	return &resultImpl{name, ServiceStatusUnknown, message}
}

func (r *resultImpl) Name() string {
	return r.name
}

func (r *resultImpl) Status() Status {
	return r.status
}

func (r *resultImpl) Message() string {
	return r.message
}

func (r *resultImpl) String() string {
	return fmt.Sprintf("{name: %s, status: %s, message: %s}", r.name, r.status, r.message)
}

// Exit prints the check result and exits the program
func (r *resultImpl) Exit() {
	fmt.Printf("%s: %s\n", r.Status(), r.Message())
	os.Exit(r.Status().Ordinal())
}
