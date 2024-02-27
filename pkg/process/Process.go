package process

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
	impl "github.com/deemount/gobpmnTypes"
)

// NewProcess ...
func NewProcess() ProcessRepository {
	return &Process{}
}

/*
 * @Setters
 */

/* Attributes */

/** Camunda **/

// SetCamundaVersionTag ...
func (process *Process) SetCamundaVersionTag(tag string) {
	process.CamundaVersionTag = tag
}

// SetCamundaJobpriority ...
func (process *Process) SetCamundaJobPriority(priority int) {
	process.CamundaJobPriority = priority
}

// SetCamundaTaskPriority ...
func (process *Process) SetCamundaTaskPriority(priority int) {
	process.CamundaTaskPriority = priority
}

// SetCamundaCandidateStarterGroups ...
func (process *Process) SetCamundaCandidateStarterGroups(groups string) {
	process.CamundaCandidateStarterGroups = groups
}

// SetCamundaCandidateStarterUsers
func (process *Process) SetCamundaCandiddateStarterUsers(users string) {
	process.CamundaCandidateStarterUsers = users
}

// SetCamundaHistoryTimeToLive ...
func (process *Process) SetCamundaHistoryTimeToLive(tolive string) {
	process.CamundaHistoryTimeToLive = tolive
}

/* Elements */

/** BPMN **/

// SetExtensionElements ...
func (process *Process) SetExtensionElements() {
	process.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*
 * @Getters
 */

/* Attributes */

/** Camunda **/

// GetCamundaVersionTag ...
func (process *Process) GetCamundaVersionTag() impl.STR_PTR {
	return &process.CamundaVersionTag
}

// GetCamundaJobpriority ...
func (process *Process) GetCamundaJobPriority() impl.INT_PTR {
	return &process.CamundaJobPriority
}

// GetCamundaTaskPriority ...
func (process *Process) GetCamundaTaskPriority() impl.INT_PTR {
	return &process.CamundaTaskPriority
}

// GetCamundaCandidateStarterGroups ...
func (process *Process) GetCamundaCandidateStarterGroups() impl.STR_PTR {
	return &process.CamundaCandidateStarterGroups
}

// GetCamundaCandidateStarterUsers
func (process *Process) GetCamundaCandiddateStarterUsers() impl.STR_PTR {
	return &process.CamundaCandidateStarterUsers
}

// GetCamundaHistoryTimeToLive ...
func (process *Process) GetCamundaHistoryTimeToLive() impl.STR_PTR {
	return &process.CamundaHistoryTimeToLive
}

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (process Process) GetExtensionElements() attributes.EXTENSION_ELEMENTS_PTR {
	return &process.ExtensionElements[0]
}
