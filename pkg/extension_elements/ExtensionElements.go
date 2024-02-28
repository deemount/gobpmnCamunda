package extension_elements

import "github.com/deemount/gobpmnCamunda/pkg/camunda"

// NewExtensionElements ...
func NewExtensionElements() ExtensionElementsRepository {
	return &ExtensionElements{}
}

/*
 * @Setters
 */

/* Elements */

/** Camunda **/

// SetCamundaConnector ...
func (extensionElements *ExtensionElements) SetCamundaConnector() {
	extensionElements.CamundaConnector = make(camunda.CAMUNDA_CONNECTOR_SLC, 1)
}

// SetCamundaProperties ...
func (extensionElements *ExtensionElements) SetCamundaProperties() {
	extensionElements.CamundaProperties = make(camunda.CAMUNDA_PROPERTIES_SLC, 1)
}

// SetCamundaFailedJobRetryCycle ...
func (extensionElements *ExtensionElements) SetCamundaFailedJobRetryCycle() {
	extensionElements.CamundaFailedJobRetryCycle = make(camunda.CAMUNDA_FAILED_JOB_RETRY_CYCLE_SLC, 1)
}

// SetCamundaFormData ...
func (extensionElements *ExtensionElements) SetCamundaFormData() {
	extensionElements.CamundaFormData = make(camunda.CAMUNDA_FORM_DATA_SLC, 1)
}

// SetCamundaInputOutput ...
func (extensionElements *ExtensionElements) SetCamundaInputOutput() {
	extensionElements.CamundaInputOutput = make([]camunda.CamundaInputOutput, 1)
}

// SetCamundaTaskListener ...
func (extensionElements *ExtensionElements) SetCamundaTaskListener(num int) {
	extensionElements.CamundaTaskListener = make(camunda.CAMUNDA_TASK_LISTENER_SLC, num)
}

// SetCamundaExecutionListener ...
func (extensionElements *ExtensionElements) SetCamundaExecutionListener(num int) {
	extensionElements.CamundaExecutionListener = make(camunda.CAMUNDA_EXECUTION_LISTENER_SLC, num)
}

// SetCamundaIn ...
func (extensionElements *ExtensionElements) SetCamundaIn(num int) {
	extensionElements.CamundaIn = make(camunda.CAMUNDA_IN_SLC, num)
}

// SetCamundaOut ...
func (extensionElements *ExtensionElements) SetCamundaOut(num int) {
	extensionElements.CamundaOut = make(camunda.CAMUNDA_OUT_SLC, num)
}

/*
 * @Getters
 */

/* Elements */

/** Camunda **/

// GetCamundaConnector ...
func (extensionElements ExtensionElements) GetCamundaConnector() camunda.CAMUNDA_CONNECTOR_PTR {
	return &extensionElements.CamundaConnector[0]
}

// GetCamundaProperties ...
func (extensionElements ExtensionElements) GetCamundaProperties() camunda.CAMUNDA_PROPERTIES_PTR {
	return &extensionElements.CamundaProperties[0]
}

// GetCamundaFailedJobRetryCycle ...
func (extensionElements ExtensionElements) GetCamundaFailedJobRetryCycle() camunda.CAMUNDA_FAILED_JOB_RETRY_CYCLE_PTR {
	return &extensionElements.CamundaFailedJobRetryCycle[0]
}

// GetCamundaFormData ...
func (extensionElements ExtensionElements) GetCamundaFormData() camunda.CAMUNDA_FORM_DATA_PTR {
	return &extensionElements.CamundaFormData[0]
}

// GetCamundaInputOutput ...
func (extensionElements ExtensionElements) GetCamundaInputOutput() camunda.CAMUNDA_IO_PTR {
	return &extensionElements.CamundaInputOutput[0]
}

// GetCamundaTaskListener ...
func (extensionElements ExtensionElements) GetCamundaTaskListener(num int) camunda.CAMUNDA_TASK_LISTENER_PTR {
	return &extensionElements.CamundaTaskListener[num]
}

// GetCamundaExecutionListener ...
func (extensionElements ExtensionElements) GetCamundaExecutionListener(num int) camunda.CAMUNDA_EXECUTION_LISTENER_PTR {
	return &extensionElements.CamundaExecutionListener[num]
}

// GetCamundaIn ...
func (extensionElements ExtensionElements) GetCamundaIn(num int) camunda.CAMUNDA_IN_PTR {
	return &extensionElements.CamundaIn[num]
}

// GetCamundaOut ...
func (extensionElements ExtensionElements) GetCamundaOut(num int) camunda.CAMUNDA_OUT_PTR {
	return &extensionElements.CamundaOut[num]
}
