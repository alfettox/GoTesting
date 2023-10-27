package mytest2

import (
	"testing"
)


func TestCheckIfBornBeforeY2K(test *testing.T){
	birthYear := 1999
	isBefore2000 := checkIfBornBeforeY2K(1999)

	if isBefore2000 != false{
		t.Errorf("%d is less than 2000, so this must be %v", birthYear, isBefore2000)
	}
}