package camunda

// NewCamundaValidation ...
func NewCamundaValidation() CamundaValidationRepository {
	return &CamundaValidation{}
}

/*
 * @Setters
 */

/* Elements */

/** Camunda **/

// SetCamundaConstraint ...
func (validation *CamundaValidation) SetCamundaConstraint(num int) {
	validation.CamundaConstraint = make(CAMUNDA_CONSTRAINT_SLC, num)
}

/*
 * @Getters
 */

/* Elements */

/** Camunda **/

// GetCamundaConstraint ...
func (validation CamundaValidation) GetCamundaConstraint(num int) CAMUNDA_CONSTRAINT_PTR {
	return &validation.CamundaConstraint[num]
}
