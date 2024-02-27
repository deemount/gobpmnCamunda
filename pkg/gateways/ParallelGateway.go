package gateways

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
	impl "github.com/deemount/gobpmnTypes"
)

// NewParallelGateway ...
func NewParallelGateway() ParallelGatewayRepository {
	return &ParallelGateway{}
}

/*
 * @Setters
 */

/* Attributes */

/** Camunda **/

// SetCamundaAsyncBefore ...
func (parallelGateway *ParallelGateway) SetCamundaAsyncBefore(asyncBefore bool) {
	parallelGateway.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (parallelGateway *ParallelGateway) SetCamundaAsyncAfter(asyncAfter bool) {
	parallelGateway.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (parallelGateway *ParallelGateway) SetCamundaJobPriority(priority int) {
	parallelGateway.CamundaJobPriority = priority
}

/* Elements */

/** BPMN **/

// SetExtensionElements ...
func (parallelGateway *ParallelGateway) SetExtensionElements() {
	parallelGateway.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*
 * @Getters
 */

/* Attributes */

/** Camunda **/

// GetCamundaAsyncBefore ...
func (parallelGateway ParallelGateway) GetCamundaAsyncBefore() impl.BOOL_PTR {
	return &parallelGateway.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (parallelGateway ParallelGateway) GetCamundaAsyncAfter() impl.BOOL_PTR {
	return &parallelGateway.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (parallelGateway ParallelGateway) GetCamundaJobPriority() impl.INT_PTR {
	return &parallelGateway.CamundaJobPriority
}

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (parallelGateway ParallelGateway) GetExtensionElements() attributes.EXTENSION_ELEMENTS_PTR {
	return &parallelGateway.ExtensionElements[0]
}
