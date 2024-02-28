package marker

import "github.com/deemount/gobpmnCamunda/pkg/extension_elements"

/*
 * Elementary
 */

// IncomingOutgoing ...
type IncomingOutgoing struct{}

// TIncomingOutgoing ...
type TIncomingOutgoing struct{}

// Category ...
type Category struct {
	extension_elements.Extension_Elements
}

// TCategory ...
type TCategory struct {
	extension_elements.TExtension_Elements
}

// CategoryValue ...
type CategoryValue struct{}

// Group ...
type Group struct {
	extension_elements.Extension_Elements
}

// TGroup ...
type TGroup struct {
	extension_elements.TExtension_Elements
}

// Incoming ...
type Incoming struct{}

// Outgoing ...
type Outgoing struct{}
