package gateways

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/impl"
)

// NewEventBasedGateway ...
func NewEventBasedGateway() EventBasedGatewayRepository {
	return &InclusiveGateway{}
}

/*
 * Default Setters
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
	eventBasedGateway.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*
 * Default Getters
 */

/* Attributes */

/** Camunda **/

// GetCamundaAsyncBefore ...
func (eventBasedGateway EventBasedGateway) GetCamundaAsyncBefore() impl.BOOL_PTR {
	return &eventBasedGateway.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (eventBasedGateway EventBasedGateway) GetCamundaAsyncAfter() impl.BOOL_PTR {
	return &eventBasedGateway.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (eventBasedGateway EventBasedGateway) GetCamundaJobPriority() impl.INT_PTR {
	return &eventBasedGateway.CamundaJobPriority
}

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (eventBasedGateway EventBasedGateway) GetExtensionElements() attributes.EXTENSION_ELEMENTS_PTR {
	return &eventBasedGateway.ExtensionElements[0]
}
