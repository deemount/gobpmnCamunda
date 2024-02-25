package tasks

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/impl"
)

// NewSendTask ...
func NewSendTask() SendTaskRepository {
	return &SendTask{}
}

/*
 * Default Setters
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
	sendTask.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*
 * Default Getters
 */

/* Attributes */

/** Camunda **/

// GetCamundaAsyncBefore ...
func (sendTask SendTask) GetCamundaAsyncBefore() impl.BOOL_PTR {
	return &sendTask.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (sendTask SendTask) GetCamundaAsyncAfter() impl.BOOL_PTR {
	return &sendTask.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (sendTask SendTask) GetCamundaJobPriority() impl.INT_PTR {
	return &sendTask.CamundaJobPriority
}

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (sendTask SendTask) GetExtensionElements() attributes.EXTENSION_ELEMENTS_PTR {
	return &sendTask.ExtensionElements[0]
}
