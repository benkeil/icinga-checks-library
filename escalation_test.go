package icinga

import "testing"

func TestEscalationLevels(t *testing.T) {
	warning := "5:"
	critical := "2:"
	e, err := NewEscalation(warning, critical)
	if err != nil {
		t.Fatalf("failed to initialize escalation: %v", err)
	}

	tests := []struct {
		value    float64
		shouldBe EscalationLevel
	}{
		{-5.0, Critical},
		{0.0, Critical},
		{1.0, Critical},
		{2.0, Warning},
		{4.0, Warning},
		{5.0, None},
		{6.0, None},
	}
	for _, test := range tests {
		level := e.Check(test.value)
		t.Logf("Check(%v) level: %v", test.value, level)
		if level != test.shouldBe {
			t.Errorf("Check(%v) should be: %v", test.value, test.shouldBe)
		}
	}
}
