package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRespositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Ashish", City: "New Delhi", ZipCode: "110011", DateOfBirth: "27/09/1996", Status: "1"},
		{Id: "1002", Name: "Aman", City: "New Delhi", ZipCode: "110011", DateOfBirth: "27/09/1996", Status: "1"},
	}
	return CustomerRepositoryStub{customers}
}
