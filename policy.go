package icinga

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
