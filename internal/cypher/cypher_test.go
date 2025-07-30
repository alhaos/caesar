package cypher

import "testing"

func TestCryptographer_Encrypt(t *testing.T) {

	data := []struct {
		inputString string
		shift       int
		expected    string
	}{
		{
			inputString: "привет мир",
			shift:       0,
			expected:    "привет мир",
		},
		{
			inputString: "привет мир",
			shift:       1,
			expected:    "рсйгёу нйс",
		},
	}

	crp := New()

	for _, datum := range data {
		result := crp.Encrypt(datum.inputString, datum.shift)
		if result != datum.expected {
			t.Error("Expected", datum.expected, "got", result)
		}
	}
}

func TestCryptographer_Decrypt(t *testing.T) {

	data := []struct {
		inputString string
		shift       int
		expected    string
	}{
		{
			inputString: "привет мир",
			shift:       0,
			expected:    "привет мир",
		},
		{
			inputString: "рсйгёу нйс",
			shift:       1,
			expected:    "привет мир",
		},
	}

	crp := New()

	for _, datum := range data {
		result := crp.Decrypt(datum.inputString, datum.shift)
		if result != datum.expected {
			t.Error("Expected", datum.expected, "got", result)
		}
	}
}
