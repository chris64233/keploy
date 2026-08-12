package models

import (
	"encoding/json"
	"testing"
	"time"
)

func TestPostgresV3CellJSONTimestampRoundTripWideYears(t *testing.T) {
	tests := []struct {
		name string
		in   time.Time
	}{
		{
			name: "postgres wide year utc",
			in:   time.Date(149206, time.December, 15, 16, 39, 16, 394721000, time.UTC),
		},
		{
			name: "postgres wide year with offset",
			in:   time.Date(12345, time.January, 2, 3, 4, 5, 678900000, time.FixedZone("", 5*3600+30*60)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wire, err := json.Marshal(PostgresV3Cell{Value: tt.in})
			if err != nil {
				t.Fatalf("marshal cell: %v", err)
			}

			var got PostgresV3Cell
			if err := json.Unmarshal(wire, &got); err != nil {
				t.Fatalf("unmarshal cell from %s: %v", wire, err)
			}

			gotTime, ok := got.Value.(time.Time)
			if !ok {
				t.Fatalf("decoded value type = %T, want time.Time", got.Value)
			}
			if gotTime.Format(time.RFC3339Nano) != tt.in.Format(time.RFC3339Nano) {
				t.Fatalf("decoded timestamp = %s, want %s",
					gotTime.Format(time.RFC3339Nano),
					tt.in.Format(time.RFC3339Nano))
			}
		})
	}
}
