package gateways

import (
	"github.com/deemount/gobpmnCamunda/pkg/extension_elements"
	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

// NewComplexGateway ...
func NewComplexGateway() ComplexGatewayRepository {
	return &ComplexGateway{}
}

/*
 * @Setters
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
	complexGateway.ExtensionElements = make(extension_elements.EXTENSION_ELEMENTS_SLC, 1)
}

/*
 * @Getters
 */

/* Attributes */

/** Camunda **/

// SetCamundaAsyncBefore ...
func (complexGateway ComplexGateway) GetCamundaAsyncBefore() gobpmnTypes.BOOL_PTR {
	return &complexGateway.CamundaAsyncBefore
}

// SetCamundaAsyncAfter ...
func (complexGateway ComplexGateway) GetCamundaAsyncAfter() gobpmnTypes.BOOL_PTR {
	return &complexGateway.CamundaAsyncAfter
}

// SetCamundaJobPriority ...
func (complexGateway ComplexGateway) GetCamundaJobPriority() gobpmnTypes.INT_PTR {
	return &complexGateway.CamundaJobPriority
}

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (complexGateway ComplexGateway) GetExtensionElements() extension_elements.EXTENSION_ELEMENTS_PTR {
	return &complexGateway.ExtensionElements[0]
}
