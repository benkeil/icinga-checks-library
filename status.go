package icinga

// Status defines the service status
type Status int

const (
	// ServiceStatusOk 0
	ServiceStatusOk Status = iota
	// ServiceStatusWarning 1
	ServiceStatusWarning
	// ServiceStatusCritical 2
	ServiceStatusCritical
	// ServiceStatusUnknown 3
	ServiceStatusUnknown
)

var statusMap = map[string]Status{
	"OK":       ServiceStatusOk,
	"WARNING":  ServiceStatusWarning,
	"CRITICAL": ServiceStatusCritical,
	"UNKNOWN":  ServiceStatusUnknown,
}

// Ordinal returns the int value
func (s Status) Ordinal() int {
	return int(s)
}

func (s Status) String() string {
	names := [...]string{
		"OK",
		"WARNING",
		"CRITICAL",
		"UNKNOWN",
	}
	if s < ServiceStatusOk || s > ServiceStatusUnknown {
		panic("invalid icinga.Status")
	}

	return names[s]
}

// NewStatus return a Status for a given string
func NewStatus(statusString string) Status {
	status, found := statusMap[statusString]
	if found {
		return status
	}
	panic("invalid icinga.Status")
}
