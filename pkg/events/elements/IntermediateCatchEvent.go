package elements

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
	impl "github.com/deemount/gobpmnTypes"
)

// NewIntermediateCatchEvent ...
func NewIntermediateCatchEvent() IntermediateCatchEventRepository {
	return &IntermediateCatchEvent{}
}

/*
 * @Setters
 */

/* Attributes */

/** Camunda **/

// SetCamundaAsyncBefore ...
func (ice *IntermediateCatchEvent) SetCamundaAsyncBefore(asyncBefore bool) {
	ice.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncBefore ...
func (ice *IntermediateCatchEvent) SetCamundaAsyncAfter(asyncAfter bool) {
	ice.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (ice *IntermediateCatchEvent) SetCamundaJobPriority(priority int) {
	ice.CamundaJobPriority = priority
}

/* Elements */

/** BPMN **/

// SetExtensionElements ...
func (ice *IntermediateCatchEvent) SetExtensionElements() {
	ice.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*
 * @Getters
 */

/* Attributes */

/** Camunda **/

// GetCamundaAsyncBefore ...
func (ice IntermediateCatchEvent) GetCamundaAsyncBefore() impl.BOOL_PTR {
	return &ice.CamundaAsyncBefore
}

// GetCamundaAsyncBefore ...
func (ice IntermediateCatchEvent) GetCamundaAsyncAfter() impl.BOOL_PTR {
	return &ice.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (ice IntermediateCatchEvent) GetCamundaJobPriority() impl.INT_PTR {
	return &ice.CamundaJobPriority
}

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (ice IntermediateCatchEvent) GetExtensionElements() attributes.EXTENSION_ELEMENTS_PTR {
	return &ice.ExtensionElements[0]
}
