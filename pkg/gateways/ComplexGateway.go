package gateways

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/impl"
)

// NewComplexGateway ...
func NewComplexGateway() ComplexGatewayRepository {
	return &ComplexGateway{}
}

/*
 * Default Setters
 */

/* Attributes */

/** Camunda **/

// SetCamundaAsyncBefore ...
func (complexGateway *ComplexGateway) SetCamundaAsyncBefore(asyncBefore bool) {
	complexGateway.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (complexGateway *ComplexGateway) SetCamundaAsyncAfter(asyncAfter bool) {
	complexGateway.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (complexGateway *ComplexGateway) SetCamundaJobPriority(priority int) {
	complexGateway.CamundaJobPriority = priority
}

/* Elements */

/** BPMN **/

// SetExtensionElements ...
func (complexGateway *ComplexGateway) SetExtensionElements() {
	complexGateway.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*
 * Default Getters
 */

/* Attributes */

/** Camunda **/

// SetCamundaAsyncBefore ...
func (complexGateway ComplexGateway) GetCamundaAsyncBefore() impl.BOOL_PTR {
	return &complexGateway.CamundaAsyncBefore
}

// SetCamundaAsyncAfter ...
func (complexGateway ComplexGateway) GetCamundaAsyncAfter() impl.BOOL_PTR {
	return &complexGateway.CamundaAsyncAfter
}

// SetCamundaJobPriority ...
func (complexGateway ComplexGateway) GetCamundaJobPriority() impl.INT_PTR {
	return &complexGateway.CamundaJobPriority
}

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (complexGateway ComplexGateway) GetExtensionElements() attributes.EXTENSION_ELEMENTS_PTR {
	return &complexGateway.ExtensionElements[0]
}
