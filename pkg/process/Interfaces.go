package process

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
	"github.com/deemount/gobpmnCamunda/pkg/camunda"
)

/*
 * @Repositories
 */

// ProcessRepository ...
type ProcessRepository interface {
	attributes.AttributesBaseElements
	camunda.CamundaProcessAttributes
}
