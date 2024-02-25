package gateways

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
	"github.com/deemount/gobpmnCamunda/pkg/camunda"
)

/*
 * @Base
 */

// GatewayBase ...
type GatewayBase interface {
	camunda.CamundaDefaultAttributes
	attributes.AttributesBaseElements
}

/*
 * @Repositories
 */

// ComplexGatewayRepository ...
type ComplexGatewayRepository interface {
	GatewayBase
}

// EventBasedGatewayRepository ...
type EventBasedGatewayRepository interface {
	GatewayBase
}

// ExclusiveGatewayRepository ...
type ExclusiveGatewayRepository interface {
	GatewayBase
}

// InclusiveGatewayRepository ...
type InclusiveGatewayRepository interface {
	GatewayBase
}

// ParallelGatewayRepository ...
type ParallelGatewayRepository interface {
	GatewayBase
}
