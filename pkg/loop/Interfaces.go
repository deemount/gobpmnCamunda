package loop

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/impl"
)

type LoopBaseID interface{}
type LoopBaseName interface{}
type LoopCamundaBase interface{}
type LoopBaseCoreElements interface{}
type LoopBase interface{}

/*
 * @Repositories
 */

// LoopCardinalityRepository ...
type LoopCardinalityRepository interface{}

// MultiInstanceLoopCharacteristicsRepository ...
type MultiInstanceLoopCharacteristicsRepository interface {
	SetCamundaAsyncBefore(asyncBefore bool)
	GetCamundaAsyncBefore() impl.BOOL_PTR
	SetCamundaAsyncAfter(asyncAfter bool)
	GetCamundaAsyncAfter() impl.BOOL_PTR
	SetCamundaCollection(collection string)
	GetCamundaCollection() impl.STR_PTR
	SetCamundaElementVariable(element string)
	GetCamundaElementVariable() impl.STR_PTR
	attributes.AttributesBaseElements
}

// ParticipantMultiplicityRepository ...
type ParticipantMultiplicityRepository interface{}

// StandardLoopCharacteristicsRepository ...
type StandardLoopCharacteristicsRepository interface{}
