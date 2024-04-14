package domain 

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer {
		{ Id: "1001", Name: "Carlos", City: "São Paulo", ZipCode: "00000-000", DateOfBirth: "08/14/1975", Status: "1" },
		{ Id: "1002", Name: "Neto", City: "São Paulo", ZipCode: "00000-000", DateOfBirth: "08/14/1975", Status: "1" },
	}
	return CustomerRepositoryStub{customers: customers}
}