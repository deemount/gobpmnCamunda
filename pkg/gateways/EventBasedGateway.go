package gateways

import (
	"github.com/deemount/gobpmnCamunda/pkg/extension_elements"
	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

// NewEventBasedGateway ...
func NewEventBasedGateway() EventBasedGatewayRepository {
	return &InclusiveGateway{}
}

/*
 * @Setters
 */

/* Attributes */

/** Camunda **/

// SetCamundaAsyncBefore ...
func (eventBasedGateway *EventBasedGateway) SetCamundaAsyncBefore(asyncBefore bool) {
	eventBasedGateway.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (eventBasedGateway *EventBasedGateway) SetCamundaAsyncAfter(asyncAfter bool) {
	eventBasedGateway.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (eventBasedGateway *EventBasedGateway) SetCamundaJobPriority(priority int) {
	eventBasedGateway.CamundaJobPriority = priority
}

/* Elements */

/** BPMN **/

// SetExtensionElements ...
func (eventBasedGateway *EventBasedGateway) SetExtensionElements() {
	eventBasedGateway.ExtensionElements = make(extension_elements.EXTENSION_ELEMENTS_SLC, 1)
}

/*
 * @Getters
 */

/* Attributes */

/** Camunda **/

// GetCamundaAsyncBefore ...
func (eventBasedGateway EventBasedGateway) GetCamundaAsyncBefore() gobpmnTypes.BOOL_PTR {
	return &eventBasedGateway.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (eventBasedGateway EventBasedGateway) GetCamundaAsyncAfter() gobpmnTypes.BOOL_PTR {
	return &eventBasedGateway.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (eventBasedGateway EventBasedGateway) GetCamundaJobPriority() gobpmnTypes.INT_PTR {
	return &eventBasedGateway.CamundaJobPriority
}

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (eventBasedGateway EventBasedGateway) GetExtensionElements() extension_elements.EXTENSION_ELEMENTS_PTR {
	return &eventBasedGateway.ExtensionElements[0]
}
