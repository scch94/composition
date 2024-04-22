package customer

type Customer struct {
	name    string
	address string
	phone   string
}

func New(name string, address string, phone string) Customer {
	return Customer{name: name, address: address, phone: phone}
}
