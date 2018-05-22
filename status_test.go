package icinga

import (
	"testing"
)

func TestSomething(t *testing.T) {
	tests := []struct {
		value           Status
		shouldBeOrdinal int
	}{
		{ServiceStatusOk, 0},
		{ServiceStatusWarning, 1},
		{ServiceStatusCritical, 2},
		{ServiceStatusUnknown, 3},
		{HostStatusUp, 0},
		{HostStatusDown, 2},
	}
	for _, test := range tests {
		t.Logf("Status(%v) is %d", test.value, test.value)
		if int(test.value) != test.shouldBeOrdinal {
			t.Errorf("Status(%v) should be %v but is %d", test.value, test.shouldBeOrdinal, test.value)
		}
	}
}
