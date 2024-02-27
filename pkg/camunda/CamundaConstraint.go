package camunda

import impl "github.com/deemount/gobpmnTypes"

// NewCamundaConstraint ...
func NewCamunadConstraint() CamundaConstraintRepository {
	return &CamundaConstraint{}
}

/*
 * Default Setters
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
 * Default Getters
 */

/* Attributes */

// GetName ...
func (cconstraint CamundaConstraint) GetName() impl.STR_PTR {
	return &cconstraint.Name
}

// GetConfig ...
func (cconstraint CamundaConstraint) GetConfig() impl.STR_PTR {
	return &cconstraint.Config
}
