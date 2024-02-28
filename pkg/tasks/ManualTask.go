package tasks

import (
	"github.com/deemount/gobpmnCamunda/pkg/extension_elements"
	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

// NewManualTask ...
func NewManualTask() ManualTaskRepository {
	return &ManualTask{}
}

/*
 * @Setters
 */

/* Attributes */

/** Camunda **/

// SetCamundaAsyncBefore ...
func (manualTask *ManualTask) SetCamundaAsyncBefore(asyncBefore bool) {
	manualTask.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (manualTask *ManualTask) SetCamundaAsyncAfter(asyncAfter bool) {
	manualTask.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (manualTask *ManualTask) SetCamundaJobPriority(priority int) {
	manualTask.CamundaJobPriority = priority
}

/* Elements */

/** BPMN **/

// SetExtensionElements ...
func (manualTask *ManualTask) SetExtensionElements() {
	manualTask.ExtensionElements = make(extension_elements.EXTENSION_ELEMENTS_SLC, 1)
}

/*
 * @Getters
 */

/* Attributes */

/** Camunda **/

// GetCamundaAsyncBefore ...
func (manualTask ManualTask) GetCamundaAsyncBefore() gobpmnTypes.BOOL_PTR {
	return &manualTask.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (manualTask ManualTask) GetCamundaAsyncAfter() gobpmnTypes.BOOL_PTR {
	return &manualTask.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (manualTask ManualTask) GetCamundaJobPriority() gobpmnTypes.INT_PTR {
	return &manualTask.CamundaJobPriority
}

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (manualTask ManualTask) GetExtensionElements() extension_elements.EXTENSION_ELEMENTS_PTR {
	return &manualTask.ExtensionElements[0]
}
