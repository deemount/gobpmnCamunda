package tasks

import (
	"github.com/deemount/gobpmnCamunda/pkg/extension_elements"
	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

// NewBusinessRuleTask ...
func NewBusinessRuleTask() BusinessRuleTaskRepository {
	return &BusinessRuleTask{}
}

/*
 * @Setters
 */

/* Attributes */

/** Camunda **/

// SetCamundaAsyncBefore ...
func (businessRuleTask *BusinessRuleTask) SetCamundaAsyncBefore(asyncBefore bool) {
	businessRuleTask.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (businessRuleTask *BusinessRuleTask) SetCamundaAsyncAfter(asyncAfter bool) {
	businessRuleTask.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (businessRuleTask *BusinessRuleTask) SetCamundaJobPriority(priority int) {
	businessRuleTask.CamundaJobPriority = priority
}

// SetCamundaClass ...
func (businessRuleTask *BusinessRuleTask) SetCamundaClass(class string) {
	businessRuleTask.CamundaClass = class
}

/* Elements */

/** BPMN **/

// SetExtensionElements ...
func (businessRuleTask *BusinessRuleTask) SetExtensionElements() {
	businessRuleTask.ExtensionElements = make(extension_elements.EXTENSION_ELEMENTS_SLC, 1)
}

/*
 * @Getters
 */

/* Attributes */

/** Camunda **/

// GetCamundaAsyncBefore ...
func (businessRuleTask BusinessRuleTask) GetCamundaAsyncBefore() gobpmnTypes.BOOL_PTR {
	return &businessRuleTask.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (businessRuleTask BusinessRuleTask) GetCamundaAsyncAfter() gobpmnTypes.BOOL_PTR {
	return &businessRuleTask.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (businessRuleTask BusinessRuleTask) GetCamundaJobPriority() gobpmnTypes.INT_PTR {
	return &businessRuleTask.CamundaJobPriority
}

// GetCamundaClass ...
func (businessRuleTask BusinessRuleTask) GetCamundaClass() gobpmnTypes.STR_PTR {
	return &businessRuleTask.CamundaClass
}

/* Elements */

/** BPMN **/

// SetExtensionElements ...
func (businessRuleTask BusinessRuleTask) GetExtensionElements() extension_elements.EXTENSION_ELEMENTS_PTR {
	return &businessRuleTask.ExtensionElements[0]
}
