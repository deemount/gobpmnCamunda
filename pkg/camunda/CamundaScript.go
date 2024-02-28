package camunda

import (
	"fmt"

	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

// NewCamundaScript ...
func NewCamundaScript() CamundaScriptRepository {
	return &CamundaScript{}
}

/*
 * @Setters
 */

/* Attributes */

// SetScriptFormat ...
func (cscript *CamundaScript) SetScriptFormat(format string) {
	cscript.ScriptFormat = fmt.Sprintf("%s", format)
}

/*
 * @Getters
 */

/* Attributes */

// GetScriptFormat ...
func (cscript CamundaScript) GetScriptFormat() gobpmnTypes.STR_PTR {
	return &cscript.ScriptFormat
}
