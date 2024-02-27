package flow

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
)

// NewAssociation ...
func NewAssociation() AssociationRepository {
	return &Association{}
}

/*
 * @Setters
 */

/* Elements */

/** BPMN **/

// SetExtensionElements ...
func (association *Association) SetExtensionElements() {
	association.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*
 * @Getters
 */

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (association Association) GetExtensionElements() attributes.EXTENSION_ELEMENTS_PTR {
	return &association.ExtensionElements[0]
}
