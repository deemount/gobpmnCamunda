package tasks

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/impl"
)

// NewManualTask ...
func NewManualTask() ManualTaskRepository {
	return &ManualTask{}
}

/*
 * Default Setters
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
	manualTask.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*
 * Default Getters
 */

/* Attributes */

/** Camunda **/

// GetCamundaAsyncBefore ...
func (manualTask ManualTask) GetCamundaAsyncBefore() impl.BOOL_PTR {
	return &manualTask.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (manualTask ManualTask) GetCamundaAsyncAfter() impl.BOOL_PTR {
	return &manualTask.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (manualTask ManualTask) GetCamundaJobPriority() impl.INT_PTR {
	return &manualTask.CamundaJobPriority
}

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (manualTask ManualTask) GetExtensionElements() attributes.EXTENSION_ELEMENTS_PTR {
	return &manualTask.ExtensionElements[0]
}
