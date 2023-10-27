package main

import (
	// "fmt"
	"testing"
)

func TestSquare(t *testing.T) {
	ans := Square(2)

	if ans != 4 {
		t.Errorf("Square(2) = %d; Should be 4", ans)
	}
}

// func TestSquareTableDriven(t *testing.T) {
// 	tests := []struct {
// 		a      int
// 		expect int
// 	}{
// 		{2, 4}, // input 2 expect 4
// 		{6, 36},
// 	}

// 	for _, tt := range tests {
// 		testname := fmt.Sprintf("%d", tt.a)
// 		t.Run(testname, func(t *testing.T) {
// 			ans := Square(tt.a)
// 			if ans != tt.expect {
// 				t.Errorf("got %d, want %d", ans, tt.expect)
// 			}
// 		})
// 	}
// }
