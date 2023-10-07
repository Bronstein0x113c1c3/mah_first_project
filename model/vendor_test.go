package model

import "testing"

func TestAllVendor(t *testing.T) {
	if len(AllVendor()) == 0 {
		t.Log("Initial, without adding :)")
	} else {
		t.Fatal("Nah, get prolem....")
	}
}
func TestAddVendor(t *testing.T) {
	if vendor := AddVendor(Vendor{
		ID:   "1",
		Name: "Bronstein",
	}); vendor.ID != "" {
		t.Log("Adding successful")
		t.Log(vendor)
	} else {
		t.Fatal("Nah, we got problem....")
	}
}

