package icinga

import (
	"bytes"
	"fmt"
	"strings"
)

type (
	// StatusPolicy interface for all status policies
	StatusPolicy interface {
		Calculate(Results) Status
	}

	defaultStatusPolicy struct{}
)

// NewDefaultStatusPolicy returns a status policy that assigns relative
// severity in accordance with conventional Nagios plugin return codes.
// Statuses associated with higher return codes are more severe.
func NewDefaultStatusPolicy() StatusPolicy {
	return &defaultStatusPolicy{}
}

func (p *defaultStatusPolicy) Calculate(results Results) Status {
	calculatedStatus := ServiceStatusOk
	for _, result := range results.All() {
		if result.Status() > calculatedStatus {
			calculatedStatus = result.Status()
		}
	}
	return calculatedStatus
}

type (
	// StatusMessagePolicy interface for all status policies
	StatusMessagePolicy interface {
		Generate(Results) string
	}

	defaultStatusMessagePolicy struct {
		order []Status
	}
)

// NewDefaultStatusMessagePolicy returns a status policy that assigns relative
// severity in accordance with conventional Nagios plugin return codes.
// Statuses associated with higher return codes are more severe.
func NewDefaultStatusMessagePolicy() StatusMessagePolicy {
	return &defaultStatusMessagePolicy{[]Status{
		ServiceStatusUnknown,
		ServiceStatusCritical,
		ServiceStatusWarning,
		ServiceStatusOk,
	}}
}

func (p *defaultStatusMessagePolicy) Generate(results Results) string {
	// group all checks by status
	statusMap := make(map[Status][]string)
	for _, result := range results.All() {
		statusMap[result.Status()] = append(statusMap[result.Status()], result.Name())
	}

	// concatente checks per status
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%s: ", results.CalculateStatus()))
	for _, status := range p.order {
		if checks, found := statusMap[status]; found {
			buffer.WriteString(fmt.Sprintf("%s: %s ", strings.ToLower(status.String()), checks))
		}
	}

	// remove white space at the end
	return strings.Trim(buffer.String(), " ")
}
