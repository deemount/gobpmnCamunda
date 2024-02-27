package tasks

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
	impl "github.com/deemount/gobpmnTypes"
)

// NewServiceTask ...
func NewServiceTask() ServiceTaskRepository {
	return &ServiceTask{}
}

/*
 * @Setters
 */

/* Attributes */

/** Camunda **/

// SetCamundaAsyncBefore ...
func (serviceTask *ServiceTask) SetCamundaAsyncBefore(asyncBefore bool) {
	serviceTask.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (serviceTask *ServiceTask) SetCamundaAsyncAfter(asyncAfter bool) {
	serviceTask.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (serviceTask *ServiceTask) SetCamundaJobPriority(priority int) {
	serviceTask.CamundaJobPriority = priority
}

/*** Make Elements ***/

// SetExtensionElements ...
func (serviceTask *ServiceTask) SetExtensionElements() {
	serviceTask.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*
 * @Getters
 */

/* Attributes */

/** Camunda **/

// GetCamundaAsyncBefore ...
func (serviceTask ServiceTask) GetCamundaAsyncBefore() impl.BOOL_PTR {
	return &serviceTask.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (serviceTask ServiceTask) GetCamundaAsyncAfter() impl.BOOL_PTR {
	return &serviceTask.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (serviceTask ServiceTask) GetCamundaJobPriority() impl.INT_PTR {
	return &serviceTask.CamundaJobPriority
}

/* Elements */

// GetExtensionElements ...
func (serviceTask ServiceTask) GetExtensionElements() attributes.EXTENSION_ELEMENTS_PTR {
	return &serviceTask.ExtensionElements[0]
}
