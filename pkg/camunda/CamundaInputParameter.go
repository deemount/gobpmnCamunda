package camunda

import gobpmnTypes "github.com/deemount/gobpmnTypes"

// NewCamundaInputParameter
func NewCamundaInputParameter() CamundaInputParameterRepository {
	return &CamundaInputParameter{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetLocalVariableName ...
func (inputParameter *CamundaInputParameter) SetLocalVariableName(variable string) {
	inputParameter.LocalVariableName = variable
}

// SetVariableAssignmentValue ...
func (inputParameter *CamundaInputParameter) SetVariableAssignmentValue(value string) {
	inputParameter.VariableAssignmentValue = value
}

/* Elements */

/** Camunda **/

// SetCamundaScript ...
func (inputParameter *CamundaInputParameter) SetCamundaScript() {
	inputParameter.CamundaScript = make(CAMUNDA_SCRIPT_SLC, 1)
}

// SetCamundaList ...
func (inputParameter *CamundaInputParameter) SetCamundaList() {
	inputParameter.CamundaList = make(CAMUNDA_LIST_SLC, 1)
}

// SetCamundaMap ...
func (inputParameter *CamundaInputParameter) SetCamundaMap() {
	inputParameter.CamundaMap = make(CAMUNDA_MAP_SLC, 1)
}

/*
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetLocalVariableName ...
func (inputParameter CamundaInputParameter) GetLocalVariableName() gobpmnTypes.STR_PTR {
	return &inputParameter.LocalVariableName
}

// GetVariableAssignmentValue ...
func (inputParameter CamundaInputParameter) GetVariableAssignmentValue() gobpmnTypes.STR_PTR {
	return &inputParameter.VariableAssignmentValue
}

/* Elements */

/** Camunda **/

// GetCamundaScript ...
func (inputParameter CamundaInputParameter) GetCamundaScript() CAMUNDA_SCRIPT_PTR {
	return &inputParameter.CamundaScript[0]
}

// GetCamundaList ...
func (inputParameter CamundaInputParameter) GetCamundaList() CAMUNDA_LIST_PTR {
	return &inputParameter.CamundaList[0]
}

// GetCamundaMap ...
func (inputParameter CamundaInputParameter) GetCamundaMap() CAMUNDA_MAP_PTR {
	return &inputParameter.CamundaMap[0]
}
