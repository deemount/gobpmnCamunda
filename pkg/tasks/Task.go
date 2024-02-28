package tasks

import (
	"github.com/deemount/gobpmnCamunda/pkg/extension_elements"
	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

// NewTask ...
func NewTask() TaskRepository {
	return &Task{}
}

/*
 * @Setters
 */

/* Attributes */

/** Camunda **/

// SetCamundaAsyncBefore ...
func (task *Task) SetCamundaAsyncBefore(asyncBefore bool) {
	task.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (task *Task) SetCamundaAsyncAfter(asyncAfter bool) {
	task.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (task *Task) SetCamundaJobPriority(priority int) {
	task.CamundaJobPriority = priority
}

/* Elements */

/** BPMN **/

// SetExtensionElements ...
func (task *Task) SetExtensionElements() {
	task.ExtensionElements = make(extension_elements.EXTENSION_ELEMENTS_SLC, 1)
}

/*
 * @Getters
 */

/* Attributes */

/** Camunda **/

// GetCamundaAsyncBefore ...
func (task Task) GetCamundaAsyncBefore() gobpmnTypes.BOOL_PTR {
	return &task.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (task Task) GetCamundaAsyncAfter() gobpmnTypes.BOOL_PTR {
	return &task.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (task Task) GetCamundaJobPriority() gobpmnTypes.INT_PTR {
	return &task.CamundaJobPriority
}

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (task Task) GetExtensionElements() extension_elements.EXTENSION_ELEMENTS_PTR {
	return &task.ExtensionElements[0]
}
