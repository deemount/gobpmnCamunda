package camunda

import impl "github.com/deemount/gobpmnTypes"

// NewCamundaValue ...
func NewCamundaValue() CamundaValueRepository {
	return &CamundaValue{}
}

/*
 * Default Setters
 */

/* Content */

// SetValue ...
func (value *CamundaValue) SetValue(val string) {
	value.Value = val
}

/*
 * Default Getters
 */

/* Content */

// GetValue ...
func (value CamundaValue) GetValue() impl.STR_PTR {
	return &value.Value
}
