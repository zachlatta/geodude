package main

import (
	"testing"

	"github.com/kellydunn/golang-geo"
)

func TestGeocode(t *testing.T) {
	query := "1600 amphitheatre parkway"
	expectedAddress := "1600 Amphitheatre Parkway, Mountain View, CA 94043, USA"
	expectedLatitude, expectedLongitude := 37.4219998, -122.0839596

	result, err := geocode(query, new(geo.GoogleGeocoder))
	if err != nil {
		t.Error(err)
	}

	if expectedAddress != result.Address {
		t.Errorf("Address doesn't match. Expected %s, got %s", expectedAddress,
			result.Address)
	}

	if expectedLatitude != result.Point.Lat() {
		t.Errorf("Latitude doesn't match. Expected %s, got %s", expectedLatitude,
			result.Point.Lat())
	}

	if expectedLongitude != result.Point.Lng() {
		t.Errorf("Longitude doesn't match. Expected %s, got %s", expectedLongitude,
			result.Point.Lng())
	}
}
