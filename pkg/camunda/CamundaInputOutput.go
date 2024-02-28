package camunda

// NewCamundaInputOutput ...
func NewCamundaInputOutput() CamundaInputOutputRepository {
	return &CamundaInputOutput{}
}

/*
 * @Setters
 */

/* Elements */

/** Camunda **/

// SetCamundaInputParameter ...
func (cio *CamundaInputOutput) SetCamundaInputParameter(num int) {
	cio.CamundaInputParameter = make(CAMUNDA_INPUT_PARAMETER_SLC, num)
}

// SetCamundaOutputParameter ...
func (cio *CamundaInputOutput) SetCamundaOutputParameter(num int) {
	cio.CamundaOutputParameter = make(CAMUNDA_OUTPUT_PARAMETER_SLC, num)
}

/*
 * @Getters
 */

/* Elements */

/** Camunda **/

// GetCamundaInputParameter ...
func (cio CamundaInputOutput) GetCamundaInputParameter(num int) CAMUNDA_INPUT_PARAMETER_PTR {
	return &cio.CamundaInputParameter[num]
}

// GetCamundaOutputParameter ...
func (cio CamundaInputOutput) GetCamundaOutputParameter(num int) CAMUNDA_OUTPUT_PARAMETER_PTR {
	return &cio.CamundaOutputParameter[num]
}
