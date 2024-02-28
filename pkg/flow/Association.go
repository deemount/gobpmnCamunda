package flow

import "github.com/deemount/gobpmnCamunda/pkg/extension_elements"

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
	association.ExtensionElements = make(extension_elements.EXTENSION_ELEMENTS_SLC, 1)
}

/*
 * @Getters
 */

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (association Association) GetExtensionElements() extension_elements.EXTENSION_ELEMENTS_PTR {
	return &association.ExtensionElements[0]
}
