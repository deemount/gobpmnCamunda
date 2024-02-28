package process

import "github.com/deemount/gobpmnCamunda/pkg/extension_elements"

/*
 * @Elementary
 */

// Process ...
type Process struct {
	extension_elements.Extension_Elements
	CamundaVersionTag             string `xml:"camunda:versionTag,attr,omitempty" json:"versionTag,omitempty"`
	CamundaJobPriority            int    `xml:"camunda:jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	CamundaTaskPriority           int    `xml:"camunda:taskPriority,attr,omitempty" json:"taskPriority,omitempty"`
	CamundaCandidateStarterGroups string `xml:"camunda:candidateStarterGroups,attr,omitempty" json:"candidateStarterGroups,omitempty"`
	CamundaCandidateStarterUsers  string `xml:"camunda:candidateStarterUsers,attr,omitempty" json:"candidateStarterUsers,omitempty"`
	CamundaHistoryTimeToLive      string `xml:"camunda:historyTimeToLive,attr,omitempty" json:"historyTimeToLive,omitempty"`
}

// TProcess ...
type TProcess struct {
	extension_elements.TExtension_Elements
	VersionTag             string `xml:"versionTag,attr,omitempty" json:"versionTag,omitempty"`
	JobPriority            int    `xml:"jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	TaskPriority           int    `xml:"taskPriority,attr,omitempty" json:"taskPriority,omitempty"`
	CandidateStarterGroups string `xml:"candidateStarterGroups,attr,omitempty" json:"candidateStarterGroups,omitempty"`
	CandidateStarterUsers  string `xml:"candidateStarterUsers,attr,omitempty" json:"candidateStarterUsers,omitempty"`
	HistoryTimeToLive      string `xml:"historyTimeToLive,attr,omitempty" json:"historyTimeToLive,omitempty"`
}
