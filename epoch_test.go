package epoch_test

import (
	"testing"
	"time"
)

func TestEpochConversion(t *testing.T) {
	tsExpected := "2016-01-01T01:00:00Z"

	tsObject, err := time.Parse(time.RFC3339, tsExpected)

	if err != nil {
		t.Error(err)
	}

	tsEpoch := tsObject.Unix()

	tsLocal := time.Unix(tsEpoch, 0)
	tsUTC := tsLocal.In(time.UTC)

	tsObserved := tsUTC.Format(time.RFC3339)

	if tsObserved != tsExpected {
		t.Errorf("Expected %v, got %v", tsExpected, tsObserved)
	}
}
