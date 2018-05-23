package icinga

import "fmt"

type (
	// StatusCheck contains the thresholds for warning and critical escalation
	StatusCheck interface {
		Check(float64) Status
		CheckInt(int) Status
		CheckInt32(int32) Status
		Compare(func() bool) Status
	}

	statusCheckImpl struct {
		warning  Range
		critical Range
		result   string
	}
)

// NewStatusCheck parse warning and critical thresholds into an StatusCheck object
func NewStatusCheck(warning string, critical string) (StatusCheck, error) {
	warningRange, err := NewRange(warning)
	if err != nil {
		return nil, fmt.Errorf("can't parse warning threshold string %v: %v", warning, err)
	}
	criticalRange, err := NewRange(critical)
	if err != nil {
		return nil, fmt.Errorf("can't parse warning threshold string %v: %v", critical, err)
	}
	return &statusCheckImpl{warningRange, criticalRange, ""}, nil
}

// NewStatusCheckCompare returns a new StatusCheck to evaluate a status based on a closure
// and return a Status based on result
func NewStatusCheckCompare(result string) (StatusCheck, error) {
	return &statusCheckImpl{nil, nil, result}, nil
}

// Check returns the escalation level if we have
func (e *statusCheckImpl) Check(value float64) Status {
	isWarning := e.warning.Check(value)
	isCritical := e.critical.Check(value)
	if isCritical {
		return ServiceStatusCritical
	} else if isWarning {
		return ServiceStatusWarning
	}
	return ServiceStatusOk
}

// CheckInt is a convenience method which does an unchecked type
// conversion from an int to a float64.
func (e *statusCheckImpl) CheckInt(value int) Status {
	return e.Check(float64(value))
}

// CheckInt32 is a convenience method which does an unchecked type
// conversion from an int to a float64.
func (e *statusCheckImpl) CheckInt32(value int32) Status {
	return e.Check(float64(value))
}

// Compare evaluates a clousre and return the Status parsed from the
// result string
func (e *statusCheckImpl) Compare(value func() bool) Status {
	if value() {
		return NewStatus(e.result)
	}
	return ServiceStatusOk
}
