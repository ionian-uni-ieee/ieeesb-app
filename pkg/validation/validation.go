package validation

type Validation map[string]error

func (v *Validation) HasError() bool {
	return len(*v) > 0
}

func (v *Validation) AddError(variableName string, err error) {
	(*v)[variableName] = err
}
