package core

/*
 * @Repositories
 */

// DefinitionsRepository ...
type DefinitionsRepository interface {
	SetCamundaSchema()
	SetZeebeSchema()
	SetModelerSchema()
	SetModelerExecutionPlatform()
	SetModelerExecutionPlatformVersion(version string)
	SetExporter()
	SetExporterVersion(version string)
}
