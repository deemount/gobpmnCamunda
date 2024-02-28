package tasks

import (
	"github.com/deemount/gobpmnCamunda/pkg/extension_elements"
	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

// NewScriptTask ...
func NewScriptTask() ScriptTaskRepository {
	return &ScriptTask{}
}

/*
 * @Setters
 */

/* Attributes */

/** Camunda **/

// SetCamundaAsyncBefore ...
func (scriptTask *ScriptTask) SetCamundaAsyncBefore(asyncBefore bool) {
	scriptTask.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (scriptTask *ScriptTask) SetCamundaAsyncAfter(asyncAfter bool) {
	scriptTask.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (scriptTask *ScriptTask) SetCamundaJobPriority(priority int) {
	scriptTask.CamundaJobPriority = priority
}

/* Elements */

/** BPMN **/

// SetExtensionElements ...
func (scriptTask *ScriptTask) SetExtensionElements() {
	scriptTask.ExtensionElements = make(extension_elements.EXTENSION_ELEMENTS_SLC, 1)
}

/*
 * @Getters
 */

/* Attributes */

/** Camunda **/

// GetCamundaAsyncBefore ...
func (scriptTask ScriptTask) GetCamundaAsyncBefore() gobpmnTypes.BOOL_PTR {
	return &scriptTask.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (scriptTask ScriptTask) GetCamundaAsyncAfter() gobpmnTypes.BOOL_PTR {
	return &scriptTask.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (scriptTask ScriptTask) GetCamundaJobPriority() gobpmnTypes.INT_PTR {
	return &scriptTask.CamundaJobPriority
}

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (scriptTask ScriptTask) GetExtensionElements() extension_elements.EXTENSION_ELEMENTS_PTR {
	return &scriptTask.ExtensionElements[0]
}
