package camunda

import gobpmnTypes "github.com/deemount/gobpmnTypes"

// NewCamundaProperty ...
func NewCamundaProperty() CamundaPropertyRepository {
	return &CamundaProperty{}
}

/*
 * @Setters
 */

/* Attributes */

// SetName ...
func (property *CamundaProperty) SetName(name string) {
	property.Name = name
}

// SetValue ...
func (property *CamundaProperty) SetValue(value string) {
	property.Value = value
}

/*
 * @Getters
 */

/* Attributes */

// GetName ...
func (property CamundaProperty) GetName() gobpmnTypes.STR_PTR {
	return &property.Name
}

// GetValue ...
func (property CamundaProperty) GetValue() gobpmnTypes.STR_PTR {
	return &property.Value
}
