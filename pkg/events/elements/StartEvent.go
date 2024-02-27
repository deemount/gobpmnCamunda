package elements

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
	impl "github.com/deemount/gobpmnTypes"
)

// NewStartEvent ...
func NewStartEvent() StartEventRepository {
	return &StartEvent{}
}

/*
 * @Setters
 */

/* Attributes */

/** Camunda **/

// SetCamundaFormKey ...
func (startEvent *StartEvent) SetCamundaFormKey(key string) {
	startEvent.CamundaFormKey = key
}

// SetCamundaFormRef ...
func (startEvent *StartEvent) SetCamundaFormRef(ref string) {
	startEvent.CamundaFormRef = ref
}

// SetCamundaFormRefBinding ...
// possible arguments: latest, deployment, version
// version is dependent to attribute camunda:formRefVersion
func (startEvent *StartEvent) SetCamundaFormRefBinding(bind string) {
	startEvent.CamundaFormRefBind = bind
}

// SetCamundaFormRefVersion ...
func (startEvent *StartEvent) SetCamundaFormRefVersion(version string) {
	startEvent.CamundaFormRefVersion = version
}

// SetCamundaAsyncBefore ...
func (startEvent *StartEvent) SetCamundaAsyncBefore(asyncBefore bool) {
	startEvent.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (startEvent *StartEvent) SetCamundaAsyncAfter(asyncAfter bool) {
	startEvent.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (startEvent *StartEvent) SetCamundaJobPriority(priority int) {
	startEvent.CamundaJobPriority = priority
}

// SetCamundaInitiator ...
func (startEvent *StartEvent) SetCamundaInitiator(initiator string) {
	startEvent.CamundaInitiator = initiator
}

/* Elements */

// SetExtensionElements ...
func (startEvent *StartEvent) SetExtensionElements() {
	startEvent.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*
 * @Getters
 */

/* Attributes */

/** Camunda **/

// SetCamundaFormKey ...
func (startEvent StartEvent) GetCamundaFormKey() impl.STR_PTR {
	return &startEvent.CamundaFormKey
}

// GetCamundaFormRef ...
func (startEvent StartEvent) GetCamundaFormRef() impl.STR_PTR {
	return &startEvent.CamundaFormRef
}

// GetCamundaFormRefBinding ...
func (startEvent StartEvent) GetCamundaFormRefBinding() impl.STR_PTR {
	return &startEvent.CamundaFormRefBind
}

// GetCamundaFormRefVersion ...
func (startEvent StartEvent) GetCamundaFormRefVersion() impl.STR_PTR {
	return &startEvent.CamundaFormRefVersion
}

// GetCamundaAsyncBefore ...
func (startEvent StartEvent) GetCamundaAsyncBefore() impl.BOOL_PTR {
	return &startEvent.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (startEvent StartEvent) GetCamundaAsyncAfter() impl.BOOL_PTR {
	return &startEvent.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (startEvent StartEvent) GetCamundaJobPriority() impl.INT_PTR {
	return &startEvent.CamundaJobPriority
}

// GetCamundaInitiator ...
func (startEvent StartEvent) GetCamundaInitiator() impl.STR_PTR {
	return &startEvent.CamundaInitiator
}

/* Elements */

// GetExtensionElements ...
func (startEvent StartEvent) GetExtensionElements() attributes.EXTENSION_ELEMENTS_PTR {
	return &startEvent.ExtensionElements[0]
}
