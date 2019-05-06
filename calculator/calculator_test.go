package calculator

import "testing"

func TestNewWithDefaults(t *testing.T) {
	calc := New(Config{})

	if calc.config.IterationCount != defaultIterationCount {
		t.Errorf("Expected default iteration count %d, got %d", defaultIterationCount, calc.config.IterationCount)
	}

	if calc.config.TransparencyTreshold != defaultTransparencyTreshold {
		t.Errorf("Expected default transparency treshold %d, got %d", defaultTransparencyTreshold, calc.config.TransparencyTreshold)
	}
}

func TestNew(t *testing.T) {
	config := Config{TransparencyTreshold: 123, IterationCount: 1, MinLuminance: 15, MaxLuminance: 200}
	calc := New(config)

	if calc.config != config {
		t.Errorf("Calculator config did not match")
	}
}
