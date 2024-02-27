package subprocesses

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
	impl "github.com/deemount/gobpmnTypes"
)

// NewAdHocSubProcess ...
func NewAdHocSubProcess() AdHocSubProcessRepository {
	return &AdHocSubProcess{}
}

/*
 * @Setters
 */

/* Attributes */

/** Camunda **/

// SetCamundaAsyncBefore ...
func (adhoc *AdHocSubProcess) SetCamundaAsyncBefore(asyncBefore bool) {
	adhoc.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (adhoc *AdHocSubProcess) SetCamundaAsyncAfter(asyncAfter bool) {
	adhoc.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (adhoc *AdHocSubProcess) SetCamundaJobPriority(priority int) {
	adhoc.CamundaJobPriority = priority
}

/* Elements */

/** BPMN **/

// SetExtensionElements ...
func (adhoc *AdHocSubProcess) SetExtensionElements() {
	adhoc.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*
 * @Getters
 */

/* Attributes */

/** Camunda **/

// GetCamundaAsyncBefore ...
func (adhoc AdHocSubProcess) GetCamundaAsyncBefore() impl.BOOL_PTR {
	return &adhoc.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (adhoc AdHocSubProcess) GetCamundaAsyncAfter() impl.BOOL_PTR {
	return &adhoc.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (adhoc AdHocSubProcess) GetCamundaJobPriority() impl.INT_PTR {
	return &adhoc.CamundaJobPriority
}

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (adhoc AdHocSubProcess) GetExtensionElements() attributes.EXTENSION_ELEMENTS_PTR {
	return &adhoc.ExtensionElements[0]
}
