package icinga

import (
	"fmt"
)

// EscalationLevel defines a escalation mode
type EscalationLevel int

const (
	// None 0
	None EscalationLevel = iota
	// Warning 1
	Warning
	// Critical 2
	Critical
)

// Ordinal returns the int value
func (s EscalationLevel) Ordinal() int {
	return int(s)
}

func (s EscalationLevel) String() string {
	names := [...]string{
		"NONE",
		"WARNING",
		"CRITICAL",
	}
	if s < None || s > Critical {
		panic("invalid icinga.EscalationLevel")
	}

	return names[s]
}

type (
	// Escalation contains the thresholds for warning and critical escalation
	Escalation interface {
		Check(float64) EscalationLevel
		CheckInt(int) EscalationLevel
		CheckInt32(int32) EscalationLevel
	}

	escalationImpl struct {
		warning  Range
		critical Range
	}
)

// NewEscalation parse warning and critical thresholds into an Escalation object
func NewEscalation(warning string, critical string) (Escalation, error) {
	warningRange, err := NewRange(warning)
	if err != nil {
		return nil, fmt.Errorf("can't parse warning threshold string %v: %v", warning, err)
	}
	criticalRange, err := NewRange(critical)
	if err != nil {
		return nil, fmt.Errorf("can't parse warning threshold string %v: %v", critical, err)
	}
	return &escalationImpl{warningRange, criticalRange}, nil
}

// Check returns the escalation level if we have
func (e *escalationImpl) Check(value float64) EscalationLevel {
	isWarning := e.warning.Check(value)
	isCritical := e.critical.Check(value)
	if isWarning && isCritical {
		return Critical
	} else if isCritical {
		return Critical
	} else if isWarning {
		return Warning
	}
	return None
}

// CheckInt is a convenience method which does an unchecked type
// conversion from an int to a float64.
func (e *escalationImpl) CheckInt(value int) EscalationLevel {
	return e.Check(float64(value))
}

// CheckInt32 is a convenience method which does an unchecked type
// conversion from an int to a float64.
func (e *escalationImpl) CheckInt32(value int32) EscalationLevel {
	return e.Check(float64(value))
}
