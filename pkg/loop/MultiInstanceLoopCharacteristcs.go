package loop

import (
	"github.com/deemount/gobpmnCamunda/pkg/extension_elements"
	gobpmnTypes "github.com/deemount/gobpmnTypes"
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
	multiInstanceLoopCharacteristics.ExtensionElements = make(extension_elements.EXTENSION_ELEMENTS_SLC, 1)
}

/*
 * @Getters
 */

/* Attributes */

/** Camunda **/

// GetCamundaAsyncBefore ...
func (multiInstanceLoopCharacteristics MultiInstanceLoopCharacteristics) GetCamundaAsyncBefore() gobpmnTypes.BOOL_PTR {
	return &multiInstanceLoopCharacteristics.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (multiInstanceLoopCharacteristics MultiInstanceLoopCharacteristics) GetCamundaAsyncAfter() gobpmnTypes.BOOL_PTR {
	return &multiInstanceLoopCharacteristics.CamundaAsyncAfter
}

// GetCamundaCollection ...
func (multiInstanceLoopCharacteristics MultiInstanceLoopCharacteristics) GetCamundaCollection() gobpmnTypes.STR_PTR {
	return &multiInstanceLoopCharacteristics.CamundaCollection
}

// GetCamundaElementVariable ...
func (multiInstanceLoopCharacteristics MultiInstanceLoopCharacteristics) GetCamundaElementVariable() gobpmnTypes.STR_PTR {
	return &multiInstanceLoopCharacteristics.CamundaElementVariable
}

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (multiInstanceLoopCharacteristics MultiInstanceLoopCharacteristics) GetExtensionElements() extension_elements.EXTENSION_ELEMENTS_PTR {
	return &multiInstanceLoopCharacteristics.ExtensionElements[0]
}
