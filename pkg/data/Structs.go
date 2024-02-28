package data

import "github.com/deemount/gobpmnCamunda/pkg/extension_elements"

// DataObject ...
type DataObject struct{}

// DataObjectReference ...
type DataObjectReference struct {
	extension_elements.Extension_Elements
}

// TDataObjectReference ...
type TDataObjectReference struct {
	extension_elements.TExtension_Elements
}

// DataStoreReference ...
type DataStoreReference struct {
	extension_elements.Extension_Elements
}

// TDataStoreReference ...
type TDataStoreReference struct {
	extension_elements.TExtension_Elements
}
