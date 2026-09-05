package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// Ex 1
var ErrInvalidID = errors.New("invalid ID")

type ErrInvalidField struct {
	MissingField string
}

func (e ErrInvalidField) Error() string {
	return fmt.Sprintf("%v", e.MissingField)
}

func main() {
	d := json.NewDecoder(strings.NewReader(data))
	count := 0
	for d.More() {
		count++
		var emp Employee
		err := d.Decode(&emp)
		if err != nil {
			fmt.Printf("record %d: %v\n", count, err)
			continue
		}
		err = ValidateEmployee(emp)
		message := fmt.Sprintf("record %d: %+v", count, emp)
		if err != nil {
			switch err := err.(type) {
			case interface{ Unwrap() []error }:
				allErrors := err.Unwrap()
				var messages []string
				for _, e := range allErrors {
					messages = append(messages, processError(e, emp))
				}
				message = fmt.Sprintf("%s allErrors: %s", message, strings.Join(messages, ", "))
			default:
				message = fmt.Sprintf("%s allErrors: %s", message, processError(err, emp))
			}
		}
		fmt.Println(message)
	}
}

func processError(err error, emp Employee) string {
	var badField ErrInvalidField
	if err != nil {
		if errors.Is(err, ErrInvalidID) { // ex 1
			return fmt.Sprintf("'%s' is not a valid ID", emp.ID)
		} else if errors.As(err, &badField) { // ex 2
			return fmt.Sprintf("missing %s", badField)
		}
		return fmt.Sprintf("%v", err)
	}
	return ""
}

const data = `
{
	"id": "ABCD-123",
	"first_name": "Bob",
	"last_name": "Bobson",
	"title": "Senior Manager"
}
{
	"id": "XYZ-123",
	"first_name": "Mary",
	"last_name": "Maryson",
	"title": "Vice President"
}
{
	"id": "BOTX-263",
	"first_name": "",
	"last_name": "Garciason",
	"title": "Manager"
}
{
	"id": "HLXO-829",
	"first_name": "Pierre",
	"last_name": "",
	"title": "Intern"
}
{
	"id": "MOXW-821",
	"first_name": "Franklin",
	"last_name": "Watanabe",
	"title": ""
}
{
	"id": "",
	"first_name": "Shelly",
	"last_name": "Shellson",
	"title": "CEO"
}
{
	"id": "YDOD-324",
	"first_name": "",
	"last_name": "",
	"title": ""
}
`

type Employee struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Title     string `json:"title"`
}

var validID = regexp.MustCompile(`\w{4}-\d{3}`)

func ValidateEmployee(e Employee) error {
	var errs []error
	if len(e.ID) == 0 {
		errs = append(errs, ErrInvalidField{MissingField: "ID"})
	}
	if !validID.MatchString(e.ID) {
		errs = append(errs, ErrInvalidID)
	}
	if len(e.FirstName) == 0 {
		errs = append(errs, ErrInvalidField{MissingField: "FirstName"})
	}
	if len(e.LastName) == 0 {
		errs = append(errs, ErrInvalidField{MissingField: "LastName"})
	}
	if len(e.Title) == 0 {
		errs = append(errs, ErrInvalidField{MissingField: "Title"})
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}
