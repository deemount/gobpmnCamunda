package flow

import "github.com/deemount/gobpmnCamunda/pkg/extension_elements"

// NewDataInputAssociation ...
func NewDataInputAssociation() DataInputAssociationRepository {
	return &DataInputAssociation{}
}

/*
 * @Setters
 */

/* Elements */

/** BPMN **/

// SetExtensionElements ...
func (dia *DataInputAssociation) SetExtensionElements() {
	dia.ExtensionElements = make(extension_elements.EXTENSION_ELEMENTS_SLC, 1)
}

/*
 * @Getters
 */

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (dia DataInputAssociation) GetExtensionElements() extension_elements.EXTENSION_ELEMENTS_PTR {
	return &dia.ExtensionElements[0]
}
