package domain

type Contact struct {
	ID       int
	FullName struct {
		FirstName  string
		LastName   string
		MiddleName string
	}
	phoneNumber string
}

func NewContact(id int, lastName string, firstName string, middleName string, phoneNumber string) *Contact {
	contact := &Contact{
		ID: id,
	}
	contact.SetPhoneNumber(phoneNumber)
	contact.FullName = struct {
		FirstName  string
		LastName   string
		MiddleName string
	}{FirstName: firstName, LastName: lastName, MiddleName: middleName}
	return contact
}

func (c *Contact) PhoneNumber() string {
	return c.phoneNumber
}

func (c *Contact) SetPhoneNumber(phoneNumber string) {
	var cleanedNumber string
	for _, char := range phoneNumber {
		if char >= '0' && char <= '9' {
			cleanedNumber += string(char)
		}
	}
	c.phoneNumber = cleanedNumber
}
