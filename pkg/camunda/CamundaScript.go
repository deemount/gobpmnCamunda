package camunda

import (
	"fmt"

	impl "github.com/deemount/gobpmnTypes"
)

// NewCamundaScript ...
func NewCamundaScript() CamundaScriptRepository {
	return &CamundaScript{}
}

/*
 * Default Setters
 */

/* Attributes */

// SetScriptFormat ...
func (cscript *CamundaScript) SetScriptFormat(format string) {
	cscript.ScriptFormat = fmt.Sprintf("%s", format)
}

/*
 * Default Getters
 */

/* Attributes */

// GetScriptFormat ...
func (cscript CamundaScript) GetScriptFormat() impl.STR_PTR {
	return &cscript.ScriptFormat
}
