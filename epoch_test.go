package epoch_test

import (
	"testing"
	"time"
)

func TestEpochConversion(t *testing.T) {
	tsExpected, err := time.Parse(time.RFC3339, "2016-01-01T01:00:00Z")

	if err != nil {
		t.Error(err)
	}

	tsEpoch := tsExpected.Unix()

	tsObserved := time.Unix(tsEpoch, 0)

	if !tsObserved.Equal(tsExpected) {
		t.Errorf("Expected %v, got %v", tsExpected, tsObserved)
	}
}
