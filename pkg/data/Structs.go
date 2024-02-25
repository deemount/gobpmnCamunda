package data

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
)

// DataObject ...
type DataObject struct{}

// DataObjectReference ...
type DataObjectReference struct {
	attributes.Attributes
}

// TDataObjectReference ...
type TDataObjectReference struct {
	attributes.TAttributes
}

// DataStoreReference ...
type DataStoreReference struct {
	attributes.Attributes
}

// TDataStoreReference ...
type TDataStoreReference struct {
	attributes.TAttributes
}
