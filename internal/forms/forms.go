package forms

import (
	"net/http"
	"net/url"
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

// Has Checks ir form field is in post and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		f.Errors.Add(field, "This Field Cannot be blank")
		return false
	}
	return true
}
