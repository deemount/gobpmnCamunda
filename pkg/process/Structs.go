package process

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
)

// Process ...
type Process struct {
	attributes.Attributes
	CamundaVersionTag             string `xml:"camunda:versionTag,attr,omitempty" json:"versionTag,omitempty"`
	CamundaJobPriority            int    `xml:"camunda:jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	CamundaTaskPriority           int    `xml:"camunda:taskPriority,attr,omitempty" json:"taskPriority,omitempty"`
	CamundaCandidateStarterGroups string `xml:"camunda:candidateStarterGroups,attr,omitempty" json:"candidateStarterGroups,omitempty"`
	CamundaCandidateStarterUsers  string `xml:"camunda:candidateStarterUsers,attr,omitempty" json:"candidateStarterUsers,omitempty"`
	CamundaHistoryTimeToLive      string `xml:"camunda:historyTimeToLive,attr,omitempty" json:"historyTimeToLive,omitempty"`
}

// TProcess ...
type TProcess struct {
	attributes.TAttributes
	VersionTag             string `xml:"versionTag,attr,omitempty" json:"versionTag,omitempty"`
	JobPriority            int    `xml:"jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	TaskPriority           int    `xml:"taskPriority,attr,omitempty" json:"taskPriority,omitempty"`
	CandidateStarterGroups string `xml:"candidateStarterGroups,attr,omitempty" json:"candidateStarterGroups,omitempty"`
	CandidateStarterUsers  string `xml:"candidateStarterUsers,attr,omitempty" json:"candidateStarterUsers,omitempty"`
	HistoryTimeToLive      string `xml:"historyTimeToLive,attr,omitempty" json:"historyTimeToLive,omitempty"`
}
