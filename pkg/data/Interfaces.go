package data

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
)

/*
 * @Repositories
 */

// DataObjectRepository ...
type DataObjectRepository interface{}

// DataObjectReferenceRepository ...
type DataObjectReferenceRepository interface {
	attributes.AttributesBaseElements
}

// DataStoreRepository ...
type DataStoreReferenceRepository interface {
	attributes.AttributesBaseElements
}
