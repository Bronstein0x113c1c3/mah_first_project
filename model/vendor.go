package model

import (
	"golang.org/x/exp/slices"
)

type Vendor struct {
	Name string `json:"vendor_name"`
	ID   string `json:"vendor_id"`
}

var listVendor []Vendor

func AllVendor() []Vendor {
	return listVendor
}
func AddVendor(ve Vendor) Vendor {
	if slices.ContainsFunc(listVendor, func(v Vendor) bool {
		return v.ID == ve.ID
	}) {
		return Vendor{}
	}
	listVendor = append(listVendor, ve)
	return ve
}
func FindVendor(id string) (bool, int) {
	res := slices.IndexFunc(listVendor, func(v Vendor) bool {
		return v.ID == id
	})
	return (res != -1), res
}
func EditVendor(id string, ve Vendor) bool {
	if found, index := FindVendor(id); found {
		listVendor[index] = ve
		return true
	} else {
		return false
	}
}

// func DeleteVendor(id string) {
// 	slices.DeleteFunc()
// }
