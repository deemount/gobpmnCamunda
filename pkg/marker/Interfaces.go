package marker

import "github.com/deemount/gobpmnCamunda/pkg/extension_elements"

/*
 * @Repositories
 */

// CategoryRepository ...
type CategoryRepository interface {
	extension_elements.ExtensionElementsBaseElements
}

// CategoryValueRepository ...
type CategoryValueRepository interface{}

// GroupRepository ...
type GroupRepository interface {
	extension_elements.ExtensionElementsBaseElements
}

// IncomingRepository ...
type IncomingRepository interface{}

// OutgoingRepository ...
type OutgoingRepository interface{}
