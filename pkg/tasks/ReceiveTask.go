package tasks

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/impl"
)

// NewReceiveTask ...
func NewReceiveTask() ReceiveTaskRepository {
	return &ReceiveTask{}
}

/*
 * Default Setters
 */

/* Attributes */

/** Camunda **/

// SetCamundaAsyncBefore ...
func (receiveTask *ReceiveTask) SetCamundaAsyncBefore(asyncBefore bool) {
	receiveTask.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (receiveTask *ReceiveTask) SetCamundaAsyncAfter(asyncAfter bool) {
	receiveTask.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (receiveTask *ReceiveTask) SetCamundaJobPriority(priority int) {
	receiveTask.CamundaJobPriority = priority
}

/* Elements */

/** BPMN **/

// SetExtensionElements ...
func (receiveTask *ReceiveTask) SetExtensionElements() {
	receiveTask.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*
 * Default Getters
 */

/* Attributes */

/** Camunda **/

// GetCamundaAsyncBefore ...
func (receiveTask ReceiveTask) GetCamundaAsyncBefore() impl.BOOL_PTR {
	return &receiveTask.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (receiveTask ReceiveTask) GetCamundaAsyncAfter() impl.BOOL_PTR {
	return &receiveTask.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (receiveTask ReceiveTask) GetCamundaJobPriority() impl.INT_PTR {
	return &receiveTask.CamundaJobPriority
}

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (receiveTask ReceiveTask) GetExtensionElements() attributes.EXTENSION_ELEMENTS_PTR {
	return &receiveTask.ExtensionElements[0]
}
