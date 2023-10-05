package model

type Data struct {
	// user   *User
	vendor *Vendor
}

var data *Data

func init() {
	data = &Data{
		// user:   &User{},
		vendor: &Vendor{},
	}
}
