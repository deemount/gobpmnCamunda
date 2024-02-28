package marker

import "github.com/deemount/gobpmnCamunda/pkg/extension_elements"

// NewCategory ...
func NewCategory() CategoryRepository {
	return &Category{}
}

/*
 * @Setters
 */

/* Elements */

/** BPMN **/

// SetExtensionElements ...
func (category *Category) SetExtensionElements() {
	category.ExtensionElements = make(extension_elements.EXTENSION_ELEMENTS_SLC, 1)
}

/*
 * @Getters
 */

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (category Category) GetExtensionElements() extension_elements.EXTENSION_ELEMENTS_PTR {
	return &category.ExtensionElements[0]
}
