# Icinga plugin library for go

Library to implement nagios checks in Go (golang).

## Usage example

The general usage pattern looks like this::

```go
func run() Result {
    // defined thresholds acording to https://nagios-plugins.org/doc/guidelines.html
    warningThresholdString := "5:"
    criticalThresholdString := "2:"

    // implemented somewhere else
    value := getActualValue()

    // check escalation
    e, err := NewEscalation(warningThresholdString, criticalThresholdString)
    if err != nil {
        t.Fatalf("failed to initialize escalation: %v", err)
    }
    level := e.Check(value)

    // If everything is ok and you have a service check
    if level == None {
        return NewResultOk("MyCheck")
    }

    // If the level is critical and you have a host check
    if level == Critical {
        return NewResult("MyCheck", HostStatusDown, "your message")
    }

    // the easiest way
    return NewResult("MyCheck", ServiceStatusForEscalationLevel(level), "your message")
}
```