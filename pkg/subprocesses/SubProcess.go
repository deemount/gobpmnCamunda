package subprocesses

import (
	"github.com/deemount/gobpmnCamunda/pkg/extension_elements"
	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

// NewSubProcess ...
func NewSubProcess() SubProcessRepository {
	return &SubProcess{}
}

/*
 * @Setters
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
	subprocess.ExtensionElements = make(extension_elements.EXTENSION_ELEMENTS_SLC, 1)
}

/*
 * @Getters
 */

/* Attributes */

/** Camunda **/

// GetCamundaAsyncBefore ...
func (subprocess SubProcess) GetCamundaAsyncBefore() gobpmnTypes.BOOL_PTR {
	return &subprocess.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (subprocess SubProcess) GetCamundaAsyncAfter() gobpmnTypes.BOOL_PTR {
	return &subprocess.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (subprocess SubProcess) GetCamundaJobPriority() gobpmnTypes.INT_PTR {
	return &subprocess.CamundaJobPriority
}

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (subprocess SubProcess) GetExtensionElements() extension_elements.EXTENSION_ELEMENTS_PTR {
	return &subprocess.ExtensionElements[0]
}
