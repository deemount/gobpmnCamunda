package subprocesses

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/impl"
)

// NewSubProcess ...
func NewSubProcess() SubProcessRepository {
	return &SubProcess{}
}

/*
 * Default Setters
 */

/* Attributes */

/** Camunda **/

// SetCamundaAsyncBefore ...
func (subprocess *SubProcess) SetCamundaAsyncBefore(asyncBefore bool) {
	subprocess.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (subprocess *SubProcess) SetCamundaAsyncAfter(asyncAfter bool) {
	subprocess.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (subprocess *SubProcess) SetCamundaJobPriority(priority int) {
	subprocess.CamundaJobPriority = priority
}

/* Elements */

/** BPMN **/

// SetExtensionElements ...
func (subprocess *SubProcess) SetExtensionElements() {
	subprocess.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*
 * Default Getters
 */

/* Attributes */

/** Camunda **/

// GetCamundaAsyncBefore ...
func (subprocess SubProcess) GetCamundaAsyncBefore() impl.BOOL_PTR {
	return &subprocess.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (subprocess SubProcess) GetCamundaAsyncAfter() impl.BOOL_PTR {
	return &subprocess.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (subprocess SubProcess) GetCamundaJobPriority() impl.INT_PTR {
	return &subprocess.CamundaJobPriority
}

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (subprocess SubProcess) GetExtensionElements() attributes.EXTENSION_ELEMENTS_PTR {
	return &subprocess.ExtensionElements[0]
}
