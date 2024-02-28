package gateways

import (
	"github.com/deemount/gobpmnCamunda/pkg/camunda"
	"github.com/deemount/gobpmnCamunda/pkg/extension_elements"
)

/*
 * @Elementary
 */

// ComplexGateway ...
type ComplexGateway struct {
	camunda.CoreAttributes
	extension_elements.Extension_Elements
}

// TComplexGateway ...
type TComplexGateway struct {
	camunda.TCoreAttributes
	extension_elements.Extension_Elements
}

// EventBasedGateway ...
type EventBasedGateway struct {
	camunda.CoreAttributes
	extension_elements.Extension_Elements
}

// TEventBasedGateway ...
type TEventBasedGateway struct {
	camunda.TCoreAttributes
	extension_elements.TExtension_Elements
}

// ExclusiveGateway ...
type ExclusiveGateway struct {
	camunda.CoreAttributes
	extension_elements.Extension_Elements
}

// TExclusiveGateway ...
type TExclusiveGateway struct {
	camunda.TCoreAttributes
	extension_elements.TExtension_Elements
}

// InclusiveGateway ...
type InclusiveGateway struct {
	camunda.CoreAttributes
	extension_elements.Extension_Elements
}

// TInclusiveGateway ...
type TInclusiveGateway struct {
	camunda.TCoreAttributes
	extension_elements.TExtension_Elements
}

// ParallelGateway ...
type ParallelGateway struct {
	camunda.CoreAttributes
	extension_elements.Extension_Elements
}

// TParallelGateway ...
type TParallelGateway struct {
	camunda.TCoreAttributes
	extension_elements.TExtension_Elements
}
