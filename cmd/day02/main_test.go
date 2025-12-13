package main

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  bool
	}{
		{"different halves", 1234, true},
		{"same halves", 1212, false},
		{"odd length", 123, true},
		{"all same digit", 5555, false},
		{"all same digit odd length", 555, true},
		{"single digit", 5, true},
		{"two digit same", 22, false},
		{"two digit different", 21, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValid(tt.value)
			if got != tt.want {
				t.Errorf("IsValid(%d) = %v, want %v", tt.value, got, tt.want)
			}
		})
	}
}

func TestIsValid2(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  bool
	}{
		{"different halves", 1234, true},
		{"same halves", 1212, false},
		{"odd length", 123, true},
		{"all same digit", 5555, false},
		{"all same digit odd length", 555, false},
		{"single digit", 5, true},
		{"two digit same", 22, false},
		{"two digit different", 21, true},
		{"repeated long sequence", 243243243, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValid2(tt.value)
			if got != tt.want {
				t.Errorf("IsValid2(%d) = %v, want %v", tt.value, got, tt.want)
			}
		})
	}
}
