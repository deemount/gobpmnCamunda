package elements

import (
	"github.com/deemount/gobpmnCamunda/pkg/extension_elements"
	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

/*
 * @Base
 */

// EventElementsCamundaBase ...
type EventElementsCamundaBase interface {
	SetCamundaAsyncBefore(asyncBefore bool)
	GetCamundaAsyncBefore() gobpmnTypes.BOOL_PTR
	SetCamundaAsyncAfter(asyncAfter bool)
	GetCamundaAsyncAfter() gobpmnTypes.BOOL_PTR
	SetCamundaJobPriority(priority int)
	GetCamundaJobPriority() gobpmnTypes.INT_PTR
}

/*
 * Repositories
 */

// EndEventRepository ...
type EndEventRepository interface {
	EventElementsCamundaBase
	extension_elements.ExtensionElementsBaseElements
}

// IntermediateCatchEventRepository ...
type IntermediateCatchEventRepository interface {
	EventElementsCamundaBase
	extension_elements.ExtensionElementsBaseElements
}

// IntermediateThrowEventRepository ...
type IntermediateThrowEventRepository interface {
	extension_elements.ExtensionElementsBaseElements
}

// StartEventRepository ...
type StartEventRepository interface {
	EventElementsCamundaBase
	extension_elements.ExtensionElementsBaseElements

	SetCamundaFormKey(key string)
	GetCamundaFormKey() gobpmnTypes.STR_PTR
	SetCamundaFormRef(ref string)
	GetCamundaFormRef() gobpmnTypes.STR_PTR
	SetCamundaFormRefBinding(bind string)
	GetCamundaFormRefBinding() gobpmnTypes.STR_PTR
	SetCamundaFormRefVersion(version string)
	GetCamundaFormRefVersion() gobpmnTypes.STR_PTR
	SetCamundaInitiator(initiator string)
	GetCamundaInitiator() gobpmnTypes.STR_PTR
}
