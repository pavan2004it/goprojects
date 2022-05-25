package weather_test

import (
	"github.com/google/go-cmp/cmp"
	"os"
	"testing"
	"weather"
)

func TestParseResponse(t *testing.T) {
	t.Parallel()
	data, err := os.ReadFile("testdata/weather.json")
	if err != nil {
		t.Fatal(err)
	}
	want := weather.Conditions{
		Summary: "Clouds",
	}
	got, err := weather.ParseResponse(data)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want.Summary, got.Summary) {
		t.Error(cmp.Diff(want.Summary, got.Summary))
	}
}
