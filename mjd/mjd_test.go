package mjd

import (
	"testing"
)

func TestMJD(t *testing.T) {
	result := Now()
	if result.Day() < 40587 {
		t.Fatal("Day seems wrong")
	}

}
