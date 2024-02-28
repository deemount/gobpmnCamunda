package camunda

import gobpmnTypes "github.com/deemount/gobpmnTypes"

// NewCamundaConstraint ...
func NewCamunadConstraint() CamundaConstraintRepository {
	return &CamundaConstraint{}
}

/*
 * @Setters
 */

/* Attributes */

// SetName ...
func (cconstraint *CamundaConstraint) SetName(name string) {
	cconstraint.Name = name
}

// SetConfig ...
func (cconstraint *CamundaConstraint) SetConfig(config string) {
	cconstraint.Config = config
}

/*
 * @Getters
 */

/* Attributes */

// GetName ...
func (cconstraint CamundaConstraint) GetName() gobpmnTypes.STR_PTR {
	return &cconstraint.Name
}

// GetConfig ...
func (cconstraint CamundaConstraint) GetConfig() gobpmnTypes.STR_PTR {
	return &cconstraint.Config
}
