package flow

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
)

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
	dia.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*
 * @Getters
 */

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (dia DataInputAssociation) GetExtensionElements() attributes.EXTENSION_ELEMENTS_PTR {
	return &dia.ExtensionElements[0]
}
