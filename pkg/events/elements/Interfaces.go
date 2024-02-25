package elements

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/impl"
)

/*
 * @Base
 */

// EventElementsCamundaBase ...
type EventElementsCamundaBase interface {
	SetCamundaAsyncBefore(asyncBefore bool)
	GetCamundaAsyncBefore() impl.BOOL_PTR
	SetCamundaAsyncAfter(asyncAfter bool)
	GetCamundaAsyncAfter() impl.BOOL_PTR
	SetCamundaJobPriority(priority int)
	GetCamundaJobPriority() impl.INT_PTR
}

/*
 * Repositories
 */

// EndEventRepository ...
type EndEventRepository interface {
	EventElementsCamundaBase
	attributes.AttributesBaseElements
}

// IntermediateCatchEventRepository ...
type IntermediateCatchEventRepository interface {
	EventElementsCamundaBase
	attributes.AttributesBaseElements
}

// IntermediateThrowEventRepository ...
type IntermediateThrowEventRepository interface {
	attributes.AttributesBaseElements
}

// StartEventRepository ...
type StartEventRepository interface {
	EventElementsCamundaBase
	attributes.AttributesBaseElements

	SetCamundaFormKey(key string)
	GetCamundaFormKey() impl.STR_PTR
	SetCamundaFormRef(ref string)
	GetCamundaFormRef() impl.STR_PTR
	SetCamundaFormRefBinding(bind string)
	GetCamundaFormRefBinding() impl.STR_PTR
	SetCamundaFormRefVersion(version string)
	GetCamundaFormRefVersion() impl.STR_PTR
	SetCamundaInitiator(initiator string)
	GetCamundaInitiator() impl.STR_PTR
}
