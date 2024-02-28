package flow

import "github.com/deemount/gobpmnCamunda/pkg/extension_elements"

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
	sequenceFlow.ExtensionElements = make(extension_elements.EXTENSION_ELEMENTS_SLC, 1)
}

/*
 * @Getters
 */

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (sequenceFlow SequenceFlow) GetExtensionElements() extension_elements.EXTENSION_ELEMENTS_PTR {
	return &sequenceFlow.ExtensionElements[0]
}
