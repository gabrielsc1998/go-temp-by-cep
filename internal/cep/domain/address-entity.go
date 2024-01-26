package address

type Address struct {
	Code     string
	State    string
	City     string
	District string
	Address  string
}

func New(code string, state string, city string, district string, address string) *Address {
	return &Address{
		Code:     code,
		State:    state,
		City:     city,
		District: district,
		Address:  address,
	}
}
