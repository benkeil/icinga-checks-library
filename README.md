# Icinga plugin library for go

Library to implement nagios checks in Go (golang).

## Usage example

The general usage pattern looks like this::

```go
func run() Result {
    // defined thresholds according to https://nagios-plugins.org/doc/guidelines.html
    warningThresholdString := "5:"
    criticalThresholdString := "2:"

    // implemented somewhere else
    value := getActualValue()

    // check escalation
    e, err := icinga.NewEscalation(warningThresholdString, criticalThresholdString)
    if err != nil {
        fmt.Errorf("failed to initialize escalation: %v", err)
    }
    // get the escalation level
    level := e.Check(value)

    // If everything is ok and you have a service check
    if level == icinga.None {
        return icinga.NewResultOk("MyCheck")
    }

    // If the level is critical and you have a host check
    if level == icinga.Critical {
        return icinga.NewResult("MyCheck", icinga.HostStatusDown, "your message")
    }

    // the easiest way
    return icinga.NewResult("MyCheck", icinga.ServiceStatusForEscalationLevel(level), "your message")
}
```