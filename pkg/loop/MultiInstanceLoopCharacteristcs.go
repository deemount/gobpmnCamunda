package loop

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
	impl "github.com/deemount/gobpmnTypes"
)

// NewMultiInstanceLoopCharacteristics ...
func NewMultiInstanceLoopCharacteristics() MultiInstanceLoopCharacteristicsRepository {
	return &MultiInstanceLoopCharacteristics{}
}

/*
 * @Setters
 */

/* Attributes */

/** Camunda **/

// SetCamundaAsyncBefore ...
func (multiInstanceLoopCharacteristics *MultiInstanceLoopCharacteristics) SetCamundaAsyncBefore(asyncBefore bool) {
	multiInstanceLoopCharacteristics.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (multiInstanceLoopCharacteristics *MultiInstanceLoopCharacteristics) SetCamundaAsyncAfter(asyncAfter bool) {
	multiInstanceLoopCharacteristics.CamundaAsyncAfter = asyncAfter
}

// SetCamundaCollection ...
func (multiInstanceLoopCharacteristics *MultiInstanceLoopCharacteristics) SetCamundaCollection(collection string) {
	multiInstanceLoopCharacteristics.CamundaCollection = collection
}

// SetCamundaElementVariable ...
func (multiInstanceLoopCharacteristics *MultiInstanceLoopCharacteristics) SetCamundaElementVariable(element string) {
	multiInstanceLoopCharacteristics.CamundaElementVariable = element
}

/* Elements */

/** BPMN **/

// SetExtensionElements ...
func (multiInstanceLoopCharacteristics *MultiInstanceLoopCharacteristics) SetExtensionElements() {
	multiInstanceLoopCharacteristics.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*
 * @Getters
 */

/* Attributes */

/** Camunda **/

// GetCamundaAsyncBefore ...
func (multiInstanceLoopCharacteristics MultiInstanceLoopCharacteristics) GetCamundaAsyncBefore() impl.BOOL_PTR {
	return &multiInstanceLoopCharacteristics.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (multiInstanceLoopCharacteristics MultiInstanceLoopCharacteristics) GetCamundaAsyncAfter() impl.BOOL_PTR {
	return &multiInstanceLoopCharacteristics.CamundaAsyncAfter
}

// GetCamundaCollection ...
func (multiInstanceLoopCharacteristics MultiInstanceLoopCharacteristics) GetCamundaCollection() impl.STR_PTR {
	return &multiInstanceLoopCharacteristics.CamundaCollection
}

// GetCamundaElementVariable ...
func (multiInstanceLoopCharacteristics MultiInstanceLoopCharacteristics) GetCamundaElementVariable() impl.STR_PTR {
	return &multiInstanceLoopCharacteristics.CamundaElementVariable
}

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (multiInstanceLoopCharacteristics MultiInstanceLoopCharacteristics) GetExtensionElements() attributes.EXTENSION_ELEMENTS_PTR {
	return &multiInstanceLoopCharacteristics.ExtensionElements[0]
}
