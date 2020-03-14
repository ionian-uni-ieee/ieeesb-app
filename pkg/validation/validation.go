package validation

// Validation contains a map with variable name as key and variable's error as value
type Validation map[string]error

// HasError returns if any map's variable has an error
func (v *Validation) HasError() bool {
	for _, err := range *v {
		if err != nil {
			return true
		}
	}
	return false
}

// AddError adds a new variable with its according error to the validation map
func (v *Validation) AddError(variableName string, err error) {
	(*v)[variableName] = err
}
