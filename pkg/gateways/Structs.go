package gateways

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
	"github.com/deemount/gobpmnCamunda/pkg/camunda"
)

/*
 * Elementary
 */

// ComplexGateway ...
type ComplexGateway struct {
	camunda.CoreAttributes
	attributes.Attributes
}

// TComplexGateway ...
type TComplexGateway struct {
	camunda.TCoreAttributes
	attributes.Attributes
}

// EventBasedGateway ...
type EventBasedGateway struct {
	camunda.CoreAttributes
	attributes.Attributes
}

// TEventBasedGateway ...
type TEventBasedGateway struct {
	camunda.TCoreAttributes
	attributes.TAttributes
}

// ExclusiveGateway ...
type ExclusiveGateway struct {
	camunda.CoreAttributes
	attributes.Attributes
}

// TExclusiveGateway ...
type TExclusiveGateway struct {
	camunda.TCoreAttributes
	attributes.TAttributes
}

// InclusiveGateway ...
type InclusiveGateway struct {
	camunda.CoreAttributes
	attributes.Attributes
}

// TInclusiveGateway ...
type TInclusiveGateway struct {
	camunda.TCoreAttributes
	attributes.TAttributes
}

// ParallelGateway ...
type ParallelGateway struct {
	camunda.CoreAttributes
	attributes.Attributes
}

// TParallelGateway ...
type TParallelGateway struct {
	camunda.TCoreAttributes
	attributes.TAttributes
}
