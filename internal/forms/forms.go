package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Custom Form struct. embeds a url.Values object
type Form struct {
	url.Values
	Errors errors
}

// Retur true if no errors
func (f *Form) Valid() bool {
	// return len(f.Errors) == 0
	if len(f.Errors) == 0 {
		return true
	} else {
		return false
	}
}

// New Initializes a Form Struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Check the required field
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This Field Cannot be blank")
		}
	}
}

// Has Checks ir form field is in post and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	input := r.Form.Get(field)
	// return input != ""
	if input == "" {
		return false
	}
	return true
}

// Checks the mininum length for the inputs on the forms
func (f *Form) MinimunLength(field string, length int, r *http.Request) bool {
	input := r.Form.Get(field)
	if len(input) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d character long", length))
		return false
	}
	return true
}
