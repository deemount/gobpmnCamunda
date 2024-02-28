package loop

import (
	"github.com/deemount/gobpmnCamunda/pkg/extension_elements"
	gobpmnTypes "github.com/deemount/gobpmnTypes"
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
	GetCamundaAsyncBefore() gobpmnTypes.BOOL_PTR
	SetCamundaAsyncAfter(asyncAfter bool)
	GetCamundaAsyncAfter() gobpmnTypes.BOOL_PTR
	SetCamundaCollection(collection string)
	GetCamundaCollection() gobpmnTypes.STR_PTR
	SetCamundaElementVariable(element string)
	GetCamundaElementVariable() gobpmnTypes.STR_PTR
	extension_elements.ExtensionElementsBaseElements
}

// ParticipantMultiplicityRepository ...
type ParticipantMultiplicityRepository interface{}

// StandardLoopCharacteristicsRepository ...
type StandardLoopCharacteristicsRepository interface{}
