package process

import (
	"github.com/deemount/gobpmnCamunda/pkg/camunda"
	"github.com/deemount/gobpmnCamunda/pkg/extension_elements"
)

/*
 * @Repositories
 */

// ProcessRepository ...
type ProcessRepository interface {
	extension_elements.ExtensionElementsBaseElements
	camunda.CamundaProcessAttributes
}
