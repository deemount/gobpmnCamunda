package elements

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
	impl "github.com/deemount/gobpmnTypes"
)

// NewEndEvent ...
func NewEndEvent() EndEventRepository {
	return &EndEvent{}
}

/*
 * @Setters
 */

/* Attributes */

/** Camunda **/

// SetCamundaAsyncBefore ...
func (endEvent *EndEvent) SetCamundaAsyncBefore(asyncBefore bool) {
	endEvent.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncBefore ...
func (endEvent *EndEvent) SetCamundaAsyncAfter(asyncAfter bool) {
	endEvent.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (endEvent *EndEvent) SetCamundaJobPriority(priority int) {
	endEvent.CamundaJobPriority = priority
}

/* Elements */

/** BPMN **/

// SetExtensionElements ...
func (endEvent *EndEvent) SetExtensionElements() {
	endEvent.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*
 * @Getters
 */

/* Attributes */

/** Camunda **/

// GetCamundaAsyncBefore ...
func (endEvent EndEvent) GetCamundaAsyncBefore() impl.BOOL_PTR {
	return &endEvent.CamundaAsyncBefore
}

// GetCamundaAsyncBefore ...
func (endEvent EndEvent) GetCamundaAsyncAfter() impl.BOOL_PTR {
	return &endEvent.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (endEvent EndEvent) GetCamundaJobPriority() impl.INT_PTR {
	return &endEvent.CamundaJobPriority
}

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (endEvent EndEvent) GetExtensionElements() attributes.EXTENSION_ELEMENTS_PTR {
	return &endEvent.ExtensionElements[0]
}
