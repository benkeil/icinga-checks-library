package icinga

import "testing"

func TestStatusCheckThreshold(t *testing.T) {
	warning := "5:"
	critical := "2:"
	e, err := NewStatusCheck(warning, critical)
	if err != nil {
		t.Fatalf("failed to initialize escalation: %v", err)
	}

	tests := []struct {
		value    float64
		shouldBe Status
	}{
		{-5.0, ServiceStatusCritical},
		{0.0, ServiceStatusCritical},
		{1.0, ServiceStatusCritical},
		{2.0, ServiceStatusWarning},
		{4.0, ServiceStatusWarning},
		{5.0, ServiceStatusOk},
		{6.0, ServiceStatusOk},
	}
	for _, test := range tests {
		level := e.Check(test.value)
		t.Logf("Check(%v) level: %v", test.value, level)
		if level != test.shouldBe {
			t.Errorf("Check(%v) should be: %v", test.value, test.shouldBe)
		}
	}
}

func TestStatusCheckCompare(t *testing.T) {
	result := "WARNING"
	e, err := NewStatusCheckCompare(result)
	if err != nil {
		t.Fatalf("failed to initialize escalation: %v", err)
	}

	evaluatesToTrue := func() bool {
		return true
	}

	evaluatesToFlase := func() bool {
		return false
	}

	tests := []struct {
		value    func() bool
		shouldBe Status
	}{
		{evaluatesToTrue, ServiceStatusWarning},
		{evaluatesToFlase, ServiceStatusOk},
	}
	for _, test := range tests {
		level := e.Compare(test.value)
		t.Logf("Compare(%v) level: %v", test.value(), level)
		if level != test.shouldBe {
			t.Errorf("Compare(%v) should be: %v", test.value(), test.shouldBe)
		}
	}
}
