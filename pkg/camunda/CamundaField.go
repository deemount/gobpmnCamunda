package camunda

import gobpmnTypes "github.com/deemount/gobpmnTypes"

// NewCamundaField ...
func NewCamundaField() CamundaFieldRepository {
	return &CamundaField{}
}

/*
 * @Setters
 */

/* Attributes */

/** Camunda **/

/* Elements */

/** Camunda **/

// SetName ...
func (cf *CamundaField) SetName(name string) {
	cf.Name = name
}

/* Elements */

/** Camunda **/

// SetCamundaExpression ...
func (cf *CamundaField) SetCamundaExpression() {
	cf.CamundaExpression = make(CAMUNDA_EXPRESSION_SLC, 1)
}

// SetCamundaString ...
func (cf *CamundaField) SetCamundaString() {
	cf.CamundaString = make(CAMUNDA_STRING_SLC, 1)
}

/*
 * @Getters
 */

/* Attributes */

// GetName ...
func (cf CamundaField) GetName() gobpmnTypes.STR_PTR {
	return &cf.Name
}

/* Elements */

/** Camunda **/

// GetCamundaExpression ...
func (cf CamundaField) GetCamundaExpression() CAMUNDA_EXPRESSION_PTR {
	return &cf.CamundaExpression[0]
}

// GetCamundaString ...
func (cf CamundaField) GetCamundaString() CAMUNDA_STRING_PTR {
	return &cf.CamundaString[0]
}
