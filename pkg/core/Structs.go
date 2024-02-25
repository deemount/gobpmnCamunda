package core

// Definitions represents the root element
type Definitions struct {
	CamundaSchema                   string `xml:"xmlns:camunda,attr,omitempty" json:"-"`
	Zeebe                           string `xml:"xmlns:zeebe,omitempty" json:"-"`
	Modeler                         string `xml:"xmlns:modeler,omitempty" json:"-"`
	ModelerExecutionPlatform        string `xml:"modeler:executionPlatform,omitempty" json:"-"`
	ModelerExecutionPlatformVersion string `xml:"modeler:executionPlatformVersion,omitempty" json:"-"`
	ID                              string `xml:"id,attr" json:"id"`
	TargetNamespace                 string `xml:"targetNamespace,attr" json:"-"`
	Exporter                        string `xml:"exporter,attr,omitempty" json:"-"`
	ExporterVersion                 string `xml:"exporterVersion,attr,omitempty" json:"-"`
}
