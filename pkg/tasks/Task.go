package tasks

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/impl"
)

// NewTask ...
func NewTask() TaskRepository {
	return &Task{}
}

/*
 * Default Setters
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
	task.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*
 * Default Getters
 */

/* Attributes */

/** Camunda **/

// GetCamundaAsyncBefore ...
func (task Task) GetCamundaAsyncBefore() impl.BOOL_PTR {
	return &task.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (task Task) GetCamundaAsyncAfter() impl.BOOL_PTR {
	return &task.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (task Task) GetCamundaJobPriority() impl.INT_PTR {
	return &task.CamundaJobPriority
}

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (task Task) GetExtensionElements() attributes.EXTENSION_ELEMENTS_PTR {
	return &task.ExtensionElements[0]
}
