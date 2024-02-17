package main

import (
	"testing"

	stats "github.com/montanaflynn/stats"
)

// unit test was developed using jennifer library
// see jennifer_code.go for the original code
func TestRegress(t *testing.T) {
	tests := []struct {
		name              string
		data              []stats.Coordinate
		expectedSlope     float64
		expectedIntercept float64
		expectError       bool
	}{{"Test Case 1: y = 2x + 1", []stats.Coordinate{{1.0, 3.0}, {2.0, 5.0}, {3.0, 7.0}}, 2.0, 1.0, false}, {"Test Case 2: y = 0x + 2", []stats.Coordinate{{1.0, 2.0}, {2.0, 2.0}, {3.0, 2.0}}, 0.0, 2.0, false}, {"Test Case 3: y = 2x + 4", []stats.Coordinate{{1.0, 6.0}, {2.0, 8.0}, {3.0, 10.0}}, 2.0, 4.0, false}}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			slope, intercept, err := regress(tc.data)
			if err != nil && tc.expectError == false {
				t.Errorf("regress() returned an unexpected error: %v", err)
			}
			if err == nil && tc.expectError {
				t.Errorf("regress() expected an error, got none")
			}
			if slope != tc.expectedSlope {
				t.Errorf("%s: Expected slope %v, got %v", tc.name, tc.expectedSlope, slope)
			}
			if intercept != tc.expectedIntercept {
				t.Errorf("%s: Expected intercept %v, got %v", tc.name, tc.expectedIntercept, intercept)
			}
		})
	}
}
