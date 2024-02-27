package flow

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
)

// NewSequenceFlow ...
func NewSequenceFlow() SequenceFlowRepository {
	return &SequenceFlow{}
}

/*
 * @Setters
 */

/* Elements */

/** BPMN **/

// SetExtensionElements ...
func (sequenceFlow *SequenceFlow) SetExtensionElements() {
	sequenceFlow.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*
 * @Getters
 */

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (sequenceFlow SequenceFlow) GetExtensionElements() attributes.EXTENSION_ELEMENTS_PTR {
	return &sequenceFlow.ExtensionElements[0]
}
