package marker

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
)

/*
 * @Repositories
 */

// CategoryRepository ...
type CategoryRepository interface {
	attributes.AttributesBaseElements
}

// CategoryValueRepository ...
type CategoryValueRepository interface{}

// GroupRepository ...
type GroupRepository interface {
	attributes.AttributesBaseElements
}

// IncomingRepository ...
type IncomingRepository interface{}

// OutgoingRepository ...
type OutgoingRepository interface{}
