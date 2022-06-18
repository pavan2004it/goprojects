package battery_test

import (
	"battery"
	"github.com/google/go-cmp/cmp"
	"os"
	"testing"
)

func TestParsePmsetOutput(t *testing.T) {
	t.Parallel()
	data, err := os.ReadFile("testdata/pmset.txt")
	if err != nil {
		t.Fatal(err)
	}
	want := battery.Status{ChargePercent: 80}
	got, err := battery.ParsePmsetOutput(string(data))
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestToJson(t *testing.T) {
	t.Parallel()
	batt := battery.Battery{
		Name:             "MacBattery-01",
		ID:               1089765,
		ChargePercent:    70,
		TimeToFullCharge: "0:00",
		Present:          true,
	}
	wantBytes, err := os.ReadFile("testdata/battery.json")
	if err != nil {
		t.Fatal(err)
	}
	want := string(wantBytes)
	got := batt.ToJSON()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
