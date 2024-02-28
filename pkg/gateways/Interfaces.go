package gateways

import (
	"github.com/deemount/gobpmnCamunda/pkg/camunda"
	"github.com/deemount/gobpmnCamunda/pkg/extension_elements"
)

/*
 * @Base
 */

// GatewayBase ...
type GatewayBase interface {
	camunda.CamundaDefaultAttributes
	extension_elements.ExtensionElementsBaseElements
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
