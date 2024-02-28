package data

import "github.com/deemount/gobpmnCamunda/pkg/extension_elements"

/*
 * @Repositories
 */

// DataObjectRepository ...
type DataObjectRepository interface{}

// DataObjectReferenceRepository ...
type DataObjectReferenceRepository interface {
	extension_elements.ExtensionElementsBaseElements
}

// DataStoreRepository ...
type DataStoreReferenceRepository interface {
	extension_elements.ExtensionElementsBaseElements
}
