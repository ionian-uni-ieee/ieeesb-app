package validation

type Validation map[string]error

func (v *Validation) HasError() bool {
	if len(*v) > 0 {
		return true
	}
	for _, err := range *v {
		if err != nil {
			return true
		}
	}
	return false
}

func (v *Validation) AddError(variableName string, err error) {
	(*v)[variableName] = err
}
