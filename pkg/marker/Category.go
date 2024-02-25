package marker

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
)

// NewCategory ...
func NewCategory() CategoryRepository {
	return &Category{}
}

/*
 * Default Setters
 */

/* Elements */

/** BPMN **/

// SetExtensionElements ...
func (category *Category) SetExtensionElements() {
	category.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*
 * Default Getters
 */

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (category Category) GetExtensionElements() attributes.EXTENSION_ELEMENTS_PTR {
	return &category.ExtensionElements[0]
}
