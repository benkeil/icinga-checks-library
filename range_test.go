package icinga

import (
	"testing"
)

func TestOutsideRangeOfZeroAndTen(t *testing.T) {
	threshold := "10"
	t.Logf("%v = < 0 or > 10, (outside the range of {0 .. 10})", threshold)
	r, err := NewRange(threshold)
	if err != nil {
		t.Fatalf("failed to parse %v: %v", threshold, err)
	}

	tests := []struct {
		value       float64
		shouldAlert bool
	}{
		{-1.0, true},
		{0.0, false},
		{1.0, false},
		{9.0, false},
		{10.0, false},
		{11.0, true},
	}
	for _, test := range tests {
		didAlert := r.Check(test.value)
		t.Logf("Check(%v) alert: %v", test.value, test.shouldAlert)
		if didAlert != test.shouldAlert {
			t.Errorf("Check(%v) should be: %v", test.value, test.shouldAlert)
		}
	}
}

func TestOutsideRangeOfTenAndInfinity(t *testing.T) {
	threshold := "10:"
	t.Logf("%v = < 10, (outside {10 .. ∞})", threshold)
	r, err := NewRange(threshold)
	if err != nil {
		t.Fatalf("failed to parse %v: %v", threshold, err)
	}

	tests := []struct {
		value       float64
		shouldAlert bool
	}{
		{-1.0, true},
		{9.0, true},
		{10.0, false},
		{11.0, false},
	}
	for _, test := range tests {
		didAlert := r.Check(test.value)
		t.Logf("Check(%v) alert: %v", test.value, test.shouldAlert)
		if didAlert != test.shouldAlert {
			t.Errorf("Check(%v) should be: %v", test.value, test.shouldAlert)
		}
	}
}

func TestOutsideRangeOfMinusInfinityAndTen(t *testing.T) {
	threshold := "~:10"
	t.Logf("%v = > 10, (outside the range of {-∞ .. 10})", threshold)
	r, err := NewRange(threshold)
	if err != nil {
		t.Fatalf("failed to parse %v: %v", threshold, err)
	}

	tests := []struct {
		value       float64
		shouldAlert bool
	}{
		{-100.0, false},
		{0.0, false},
		{10.0, false},
		{11.0, true},
	}
	for _, test := range tests {
		didAlert := r.Check(test.value)
		t.Logf("Check(%v) alert: %v", test.value, test.shouldAlert)
		if didAlert != test.shouldAlert {
			t.Errorf("Check(%v) should be: %v", test.value, test.shouldAlert)
		}
	}
}

func TestOutsideRangeOfTenAndTwenty(t *testing.T) {
	threshold := "10:20"
	t.Logf("%v = < 10 or > 20, (outside the range of {10 .. 20})", threshold)
	r, err := NewRange(threshold)
	if err != nil {
		t.Fatalf("failed to parse %v: %v", threshold, err)
	}

	tests := []struct {
		value       float64
		shouldAlert bool
	}{
		{0.0, true},
		{9.0, true},
		{10.0, false},
		{11.0, false},
		{19.0, false},
		{20.0, false},
		{21.0, true},
	}
	for _, test := range tests {
		didAlert := r.Check(test.value)
		t.Logf("Check(%v) alert: %v", test.value, test.shouldAlert)
		if didAlert != test.shouldAlert {
			t.Errorf("Check(%v) should be: %v", test.value, test.shouldAlert)
		}
	}
}

func TestInsideRangeOfTenAndTwenty(t *testing.T) {
	threshold := "@10:20"
	t.Logf("%v = ≥ 10 and ≤ 20, (inside the range of {10 .. 20})", threshold)
	r, err := NewRange(threshold)
	if err != nil {
		t.Fatalf("failed to parse %v: %v", threshold, err)
	}

	tests := []struct {
		value       float64
		shouldAlert bool
	}{
		{9.0, false},
		{10.0, true},
		{11.0, true},
		{19.0, true},
		{20.0, true},
		{21.0, false},
	}
	for _, test := range tests {
		didAlert := r.Check(test.value)
		t.Logf("Check(%v) alert: %v", test.value, test.shouldAlert)
		if didAlert != test.shouldAlert {
			t.Errorf("Check(%v) should be: %v", test.value, test.shouldAlert)
		}
	}
}
