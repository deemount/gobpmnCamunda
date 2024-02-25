package marker

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
)

/*
 * Elementary
 */

// IncomingOutgoing ...
type IncomingOutgoing struct{}

// TIncomingOutgoing ...
type TIncomingOutgoing struct{}

// Category ...
type Category struct {
	attributes.Attributes
}

// TCategory ...
type TCategory struct {
	attributes.TAttributes
}

// CategoryValue ...
type CategoryValue struct{}

// Group ...
type Group struct {
	attributes.Attributes
}

// TGroup ...
type TGroup struct {
	attributes.TAttributes
}

// Incoming ...
type Incoming struct{}

// Outgoing ...
type Outgoing struct{}
