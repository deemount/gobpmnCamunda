package tasks

import (
	"github.com/deemount/gobpmnCamunda/pkg/extension_elements"
	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

// NewSendTask ...
func NewSendTask() SendTaskRepository {
	return &SendTask{}
}

/*
 * @Setters
 */

/* Attributes */

/** Camunda **/

// SetCamundaAsyncBefore ...
func (sendTask *SendTask) SetCamundaAsyncBefore(asyncBefore bool) {
	sendTask.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (sendTask *SendTask) SetCamundaAsyncAfter(asyncAfter bool) {
	sendTask.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (sendTask *SendTask) SetCamundaJobPriority(priority int) {
	sendTask.CamundaJobPriority = priority
}

/* Elements */

/** BPMN **/

// SetExtensionElements ...
func (sendTask *SendTask) SetExtensionElements() {
	sendTask.ExtensionElements = make(extension_elements.EXTENSION_ELEMENTS_SLC, 1)
}

/*
 * @Getters
 */

/* Attributes */

/** Camunda **/

// GetCamundaAsyncBefore ...
func (sendTask SendTask) GetCamundaAsyncBefore() gobpmnTypes.BOOL_PTR {
	return &sendTask.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (sendTask SendTask) GetCamundaAsyncAfter() gobpmnTypes.BOOL_PTR {
	return &sendTask.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (sendTask SendTask) GetCamundaJobPriority() gobpmnTypes.INT_PTR {
	return &sendTask.CamundaJobPriority
}

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (sendTask SendTask) GetExtensionElements() extension_elements.EXTENSION_ELEMENTS_PTR {
	return &sendTask.ExtensionElements[0]
}
