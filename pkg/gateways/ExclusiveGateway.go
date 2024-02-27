package gateways

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
	impl "github.com/deemount/gobpmnTypes"
)

// NewExclusiveGateway ...
func NewExclusiveGateway() ExclusiveGatewayRepository {
	return &InclusiveGateway{}
}

/*
 * @Setters
 */

/* Attributes */

/** Camunda **/

// SetCamundaAsyncBefore ...
func (exclusiveGateway *ExclusiveGateway) SetCamundaAsyncBefore(asyncBefore bool) {
	exclusiveGateway.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (exclusiveGateway *ExclusiveGateway) SetCamundaAsyncAfter(asyncAfter bool) {
	exclusiveGateway.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (exclusiveGateway *ExclusiveGateway) SetCamundaJobPriority(priority int) {
	exclusiveGateway.CamundaJobPriority = priority
}

/* Elements */

/** BPMN **/

// SetExtensionElements ...
func (exclusiveGateway *ExclusiveGateway) SetExtensionElements() {
	exclusiveGateway.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*
 * @Getters
 */

/* Attributes */

/** Camunda **/

// GetCamundaAsyncBefore ...
func (exclusiveGateway ExclusiveGateway) GetCamundaAsyncBefore() impl.BOOL_PTR {
	return &exclusiveGateway.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (exclusiveGateway ExclusiveGateway) GetCamundaAsyncAfter() impl.BOOL_PTR {
	return &exclusiveGateway.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (exclusiveGateway ExclusiveGateway) GetCamundaJobPriority() impl.INT_PTR {
	return &exclusiveGateway.CamundaJobPriority
}

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (exclusiveGateway ExclusiveGateway) GetExtensionElements() attributes.EXTENSION_ELEMENTS_PTR {
	return &exclusiveGateway.ExtensionElements[0]
}
