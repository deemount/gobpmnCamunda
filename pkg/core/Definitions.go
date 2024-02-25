package core

var (
	schemaCamunda        = "http://camunda.org/schema/1.0/bpmn"
	schemaCamundaZebee   = "http://camunda.org/schema/zeebe/1.0"
	schemaCamundaModeler = "http://camunda.org/schema/modeler/1.0"
)

/*
 * Default Setters
 */

/* Attributes */

/** Camunda **/

// SetCamundaSchema ...
func (definitions *Definitions) SetCamundaSchema() {
	definitions.CamundaSchema = schemaCamunda
}

// SetZeebe ...
func (definitions *Definitions) SetZeebeSchema() {
	definitions.Zeebe = schemaCamundaZebee
}

// SetModeler ...
func (definitions *Definitions) SetModelerSchema() {
	definitions.Modeler = schemaCamundaModeler
}

// SetModelerExecutionPlatform ...
func (definitions *Definitions) SetModelerExecutionPlatform() {
	definitions.Modeler = "Camunda Cloud"
}

// SetModelerExecutionPlatform ...
// can be ^1.1.0
func (definitions *Definitions) SetModelerExecutionPlatformVersion(version string) {
	definitions.Modeler = version
}

// SetExporter ...
func (definitions *Definitions) SetExporter() {
	definitions.Exporter = "Camunda Modeler"
}

// SetExporterVersion ...
// can be ^4.5.0 (start with this library)
func (definitions *Definitions) SetExporterVersion(version string) {
	definitions.ExporterVersion = version
}
