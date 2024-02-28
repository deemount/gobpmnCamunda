package tasks

import (
	"github.com/deemount/gobpmnCamunda/pkg/extension_elements"
	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

// NewReceiveTask ...
func NewReceiveTask() ReceiveTaskRepository {
	return &ReceiveTask{}
}

/*
 * @Setters
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
	receiveTask.ExtensionElements = make(extension_elements.EXTENSION_ELEMENTS_SLC, 1)
}

/*
 * @Getters
 */

/* Attributes */

/** Camunda **/

// GetCamundaAsyncBefore ...
func (receiveTask ReceiveTask) GetCamundaAsyncBefore() gobpmnTypes.BOOL_PTR {
	return &receiveTask.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (receiveTask ReceiveTask) GetCamundaAsyncAfter() gobpmnTypes.BOOL_PTR {
	return &receiveTask.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (receiveTask ReceiveTask) GetCamundaJobPriority() gobpmnTypes.INT_PTR {
	return &receiveTask.CamundaJobPriority
}

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (receiveTask ReceiveTask) GetExtensionElements() extension_elements.EXTENSION_ELEMENTS_PTR {
	return &receiveTask.ExtensionElements[0]
}
