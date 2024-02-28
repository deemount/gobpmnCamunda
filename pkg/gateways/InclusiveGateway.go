package gateways

import (
	"github.com/deemount/gobpmnCamunda/pkg/extension_elements"
	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

func NewInclusiveGateway() InclusiveGatewayRepository {
	return &InclusiveGateway{}
}

/*
 * @Setters
 */

/* Attributes */

/** Camunda **/

// SetCamundaAsyncBefore ...
func (inclusiveGateway *InclusiveGateway) SetCamundaAsyncBefore(asyncBefore bool) {
	inclusiveGateway.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (inclusiveGateway *InclusiveGateway) SetCamundaAsyncAfter(asyncAfter bool) {
	inclusiveGateway.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (inclusiveGateway *InclusiveGateway) SetCamundaJobPriority(priority int) {
	inclusiveGateway.CamundaJobPriority = priority
}

/* Elements */

/** BPMN **/

// SetExtensionElements ...
func (inclusiveGateway *InclusiveGateway) SetExtensionElements() {
	inclusiveGateway.ExtensionElements = make(extension_elements.EXTENSION_ELEMENTS_SLC, 1)
}

/*
 * @Setters
 */

/* Attributes */

/** Camunda **/

// SetCamundaAsyncBefore ...
func (inclusiveGateway InclusiveGateway) GetCamundaAsyncBefore() gobpmnTypes.BOOL_PTR {
	return &inclusiveGateway.CamundaAsyncBefore
}

// SetCamundaAsyncAfter ...
func (inclusiveGateway InclusiveGateway) GetCamundaAsyncAfter() gobpmnTypes.BOOL_PTR {
	return &inclusiveGateway.CamundaAsyncAfter
}

// SetCamundaJobPriority ...
func (inclusiveGateway InclusiveGateway) GetCamundaJobPriority() gobpmnTypes.INT_PTR {
	return &inclusiveGateway.CamundaJobPriority
}

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (inclusiveGateway InclusiveGateway) GetExtensionElements() extension_elements.EXTENSION_ELEMENTS_PTR {
	return &inclusiveGateway.ExtensionElements[0]
}
