package subprocesses

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
	impl "github.com/deemount/gobpmnTypes"
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
	ca.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*
 * @Getters
 */

/* Attributes */

/** Camunda **/

// GetCamundaAsyncBefore ...
func (ca CallActivity) GetCamundaAsyncBefore() impl.BOOL_PTR {
	return &ca.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (ca CallActivity) GetCamundaAsyncAfter() impl.BOOL_PTR {
	return &ca.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (ca CallActivity) GetCamundaJobPriority() impl.INT_PTR {
	return &ca.CamundaJobPriority
}

// GetCamundaCalledElementTenantID ...
func (ca CallActivity) GetCamundaCalledElementTenantID() impl.STR_PTR {
	return &ca.CamundaCalledElementTenantID
}

// GetCamundaVariableMappingClass ...
func (ca CallActivity) GetCamundaVariableMappingClass() impl.STR_PTR {
	return &ca.CamundaVariableMappingClass
}

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (ca CallActivity) GetExtensionElements() attributes.EXTENSION_ELEMENTS_PTR {
	return &ca.ExtensionElements[0]
}
