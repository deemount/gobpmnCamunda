package tasks

import (
	"github.com/deemount/gobpmnCamunda/pkg/extension_elements"
	gobpmnTypes "github.com/deemount/gobpmnTypes"
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
	serviceTask.ExtensionElements = make(extension_elements.EXTENSION_ELEMENTS_SLC, 1)
}

/*
 * @Getters
 */

/* Attributes */

/** Camunda **/

// GetCamundaAsyncBefore ...
func (serviceTask ServiceTask) GetCamundaAsyncBefore() gobpmnTypes.BOOL_PTR {
	return &serviceTask.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (serviceTask ServiceTask) GetCamundaAsyncAfter() gobpmnTypes.BOOL_PTR {
	return &serviceTask.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (serviceTask ServiceTask) GetCamundaJobPriority() gobpmnTypes.INT_PTR {
	return &serviceTask.CamundaJobPriority
}

/* Elements */

// GetExtensionElements ...
func (serviceTask ServiceTask) GetExtensionElements() extension_elements.EXTENSION_ELEMENTS_PTR {
	return &serviceTask.ExtensionElements[0]
}
