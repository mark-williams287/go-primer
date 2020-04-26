package organization

import (
	"errors"
	"fmt"
	"strings"
)

// Type declaration
type TwitterHandler string

// Type alias
//type TwitterHandler = string

type Identifiable interface {
	ID() string
}

type Name struct {
	first string
	last  string
}

func (n *Name) FullName() string {
	return fmt.Sprintf("%s %s", n.first, n.last)
}

type Employee struct {
	Name
}

type Person struct {
	Name
	twitterHandler TwitterHandler
}

func NewPerson(firstName, lastName string) Person {
	return Person{
		Name: Name{
			first: firstName,
			last:  lastName,
		},
	}
}

func (p *Person) ID() string {
	return "12345"
}

func (p *Person) SetTwitterHandler(twitterHandler TwitterHandler) error {
	if len(twitterHandler) == 0 {
		p.twitterHandler = twitterHandler
	} else if !strings.HasPrefix(string(twitterHandler), "@") {
		return errors.New("Twitter Handler must start with an @ symbol")
	}

	p.twitterHandler = twitterHandler
	return nil
}

func (p *Person) TwitterHandler() TwitterHandler {
	return p.twitterHandler
}

func (twitterHandler TwitterHandler) RedirectUrl() string {
	cleanHandler := strings.TrimPrefix(string(twitterHandler), "@")
	return fmt.Sprintf("https://www.twitter.com/%s", cleanHandler)
}
