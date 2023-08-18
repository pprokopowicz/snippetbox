package main

import (
	"testing"
	"time"

	"github.com/pprokopowicz/snippetbox/internal/assert"
)

func TestHumanDate(t *testing.T) {
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2023, 8, 18, 16, 54, 0, 0, time.UTC),
			want: "18 Aug 2023 at 16:54",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2023, 8, 18, 16, 54, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "18 Aug 2023 at 15:54",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)

			assert.Equal(t, hd, tt.want)
		})
	}
}
