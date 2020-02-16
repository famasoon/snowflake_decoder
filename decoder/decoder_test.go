package decoder

import (
	"testing"
)

const (
	ID           = "1228330164804837377"
	Timestamp    = 292856732560
	Unixtime     = 1581691707217
	Date         = "2020-02-14 23:48:27.00000001 +0900 JST"
	WorkerID     = 14
	DatacenterID = 11
	Sequence     = 1
)

func TestDecoder(t *testing.T) {
	input := int64(1228330164804837377)
	snowFlake := DecodeSnowFlake(input)

	if snowFlake.ID != ID {
		t.Fatalf("expected = %q , got=%q", ID, snowFlake.ID)
	}

	if snowFlake.Timestamp != Timestamp {
		t.Fatalf("expected = %q , got=%q", Timestamp, snowFlake.Timestamp)
	}

	if snowFlake.Unixtime != Unixtime {
		t.Fatalf("expected = %q , got=%q", Unixtime, snowFlake.Unixtime)
	}

	if snowFlake.Date != Date {
		t.Fatalf("expected = %q , got=%q", Date, snowFlake.Date)
	}

	if snowFlake.WorkerID != WorkerID {
		t.Fatalf("expected = %q , got=%q", WorkerID, snowFlake.WorkerID)
	}

	if snowFlake.DatacenterID != DatacenterID {
		t.Fatalf("expected = %q , got=%q", DatacenterID, snowFlake.DatacenterID)
	}

	if snowFlake.Sequence != Sequence {
		t.Fatalf("expected = %q , got=%q", Sequence, snowFlake.Sequence)
	}
}
