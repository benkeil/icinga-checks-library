package icinga

import (
	"testing"
)

func TestDefaultStatusMessagePolicy(t *testing.T) {
	results := NewResults()
	results.Add(NewResult("check 1", ServiceStatusWarning, "some warning"))
	results.Add(NewResult("check 2", ServiceStatusWarning, "some warning"))
	results.Add(NewResult("check 3", ServiceStatusCritical, "some critical"))
	results.Add(NewResult("check 4", ServiceStatusOk, "some ok"))

	shouldBeOneOf := make(map[string]bool)
	shouldBeOneOf["CRITICAL: critical: [check 3] warning: [check 2 check 1] ok: [check 4]"] = true
	shouldBeOneOf["CRITICAL: critical: [check 3] warning: [check 1 check 2] ok: [check 4]"] = true
	message := results.GenerateMessage()
	t.Logf("GenerateMessage() is: %v", message)
	if _, found := shouldBeOneOf[message]; !found {
		t.Errorf("GenerateMessage() should be one of: %v", shouldBeOneOf)
	}
}
