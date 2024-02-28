package camunda

import gobpmnTypes "github.com/deemount/gobpmnTypes"

// NewCamundaOutputParameter ...
func NewCamundaOutputParameter() CamundaOutputParameterRepository {
	return &CamundaOutputParameter{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetName ...
func (outputParameter *CamundaOutputParameter) SetLocalVariableName(variable string) {
	outputParameter.LocalVariableName = variable
}

// SetVariableAssignmentValue ...
func (outputParameter *CamundaOutputParameter) SetVariableAssignmentValue(value string) {
	outputParameter.VariableAssignmentValue = value
}

/* Elements */

/** Camunda **/

// SetCamundaScript ...
func (outputParameter *CamundaOutputParameter) SetCamundaScript() {
	outputParameter.CamundaScript = make(CAMUNDA_SCRIPT_SLC, 1)
}

// SetCamundaList ...
func (outputParameter *CamundaOutputParameter) SetCamundaList() {
	outputParameter.CamundaList = make(CAMUNDA_LIST_SLC, 1)
}

// SetCamundaMap ...
func (outputParameter *CamundaOutputParameter) SetCamundaMap() {
	outputParameter.CamundaMap = make(CAMUNDA_MAP_SLC, 1)
}

/*
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetName ...
func (outputParameter CamundaOutputParameter) GetLocalVariableName() gobpmnTypes.STR_PTR {
	return &outputParameter.LocalVariableName
}

// GetVariableAssignmentValue ...
func (outputParameter CamundaOutputParameter) GetVariableAssignmentValue() gobpmnTypes.STR_PTR {
	return &outputParameter.VariableAssignmentValue
}

/* Elements */

/** Camunda **/

// GetCamundaScript ...
func (outputParameter CamundaOutputParameter) GetCamundaScript() CAMUNDA_SCRIPT_PTR {
	return &outputParameter.CamundaScript[0]
}

// GetCamundaList ...
func (outputParameter CamundaOutputParameter) GetCamundaList() CAMUNDA_LIST_PTR {
	return &outputParameter.CamundaList[0]
}

// GetCamundaMap ...
func (outputParameter CamundaOutputParameter) GetCamundaMap() CAMUNDA_MAP_PTR {
	return &outputParameter.CamundaMap[0]
}
