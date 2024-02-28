package subprocesses

import (
	"github.com/deemount/gobpmnCamunda/pkg/extension_elements"
	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

// NewCallActivity ...
func NewCallActivity() CallActivityRepository {
	return &CallActivity{}
}

/*
 * @Setters
 */

/* Attributes */

/** Camunda **/

// SetCamundaAsyncBefore ...
func (ca *CallActivity) SetCamundaAsyncBefore(asyncBefore bool) {
	ca.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (ca *CallActivity) SetCamundaAsyncAfter(asyncAfter bool) {
	ca.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (ca *CallActivity) SetCamundaJobPriority(priority int) {
	ca.CamundaJobPriority = priority
}

// SetCamundaCalledElementTenantID ...
func (ca *CallActivity) SetCamundaCalledElementTenantID(tenantID string) {
	ca.CamundaCalledElementTenantID = tenantID
}

// SetCamundaVariableMappingClass ...
func (ca *CallActivity) SetCamundaVariableMappingClass(class string) {
	ca.CamundaVariableMappingClass = class
}

/* Elements */

/** BPMN **/

// SetExtensionElements ...
func (ca *CallActivity) SetExtensionElements() {
	ca.ExtensionElements = make(extension_elements.EXTENSION_ELEMENTS_SLC, 1)
}

/*
 * @Getters
 */

/* Attributes */

/** Camunda **/

// GetCamundaAsyncBefore ...
func (ca CallActivity) GetCamundaAsyncBefore() gobpmnTypes.BOOL_PTR {
	return &ca.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (ca CallActivity) GetCamundaAsyncAfter() gobpmnTypes.BOOL_PTR {
	return &ca.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (ca CallActivity) GetCamundaJobPriority() gobpmnTypes.INT_PTR {
	return &ca.CamundaJobPriority
}

// GetCamundaCalledElementTenantID ...
func (ca CallActivity) GetCamundaCalledElementTenantID() gobpmnTypes.STR_PTR {
	return &ca.CamundaCalledElementTenantID
}

// GetCamundaVariableMappingClass ...
func (ca CallActivity) GetCamundaVariableMappingClass() gobpmnTypes.STR_PTR {
	return &ca.CamundaVariableMappingClass
}

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (ca CallActivity) GetExtensionElements() extension_elements.EXTENSION_ELEMENTS_PTR {
	return &ca.ExtensionElements[0]
}
