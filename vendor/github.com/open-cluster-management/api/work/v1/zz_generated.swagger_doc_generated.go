package v1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_Manifest = map[string]string{
	"": "Manifest represents a resource to be deployed on spoke cluster",
}

func (Manifest) SwaggerDoc() map[string]string {
	return map_Manifest
}

var map_ManifestCondition = map[string]string{
	"":             "ManifestCondition represents the conditions of the resources deployed on spoke cluster",
	"resourceMeta": "ResourceMeta represents the gvk, name and namespace of a resoure",
	"conditions":   "Conditions represents the conditions of this resource on spoke cluster",
}

func (ManifestCondition) SwaggerDoc() map[string]string {
	return map_ManifestCondition
}

var map_ManifestResourceMeta = map[string]string{
	"":          "ManifestResourceMeta represents the gvk, gvr, name and namespace of a resoure",
	"ordinal":   "Ordinal represents the index of the manifest on spec",
	"group":     "Group is the API Group of the kubernetes resource",
	"version":   "Version is the version of the kubernetes resource",
	"kind":      "Kind is the kind of the kubernetes resource",
	"resource":  "Resource is the resource name of the kubernetes resource",
	"name":      "Name is the name of the kubernetes resource",
	"namespace": "Name is the namespace of the kubernetes resource",
}

func (ManifestResourceMeta) SwaggerDoc() map[string]string {
	return map_ManifestResourceMeta
}

var map_ManifestResourceStatus = map[string]string{
	"":          "ManifestResourceStatus represents the status of each resource in manifest work deployed on spoke cluster",
	"manifests": "Manifests represents the condition of manifests deployed on spoke cluster. Valid condition types are: 1. Progressing represents the resource is being applied on spoke cluster. 2. Applied represents the resource is applied successfully on spoke cluster. 3. Available represents the resource exists on the spoke cluster. 4. Degraded represents the current state of resource does not match the desired state for a certain period.",
}

func (ManifestResourceStatus) SwaggerDoc() map[string]string {
	return map_ManifestResourceStatus
}

var map_ManifestWork = map[string]string{
	"":       "ManifestWork represents a manifests workload that hub wants to deploy on the spoke cluster. A manifest workload is defined as a set of kubernetes resources. ManifestWork must be created in the cluster namespace on the hub, so that agent on the corresponding spoke cluster can access this resource and deploy on the spoke cluster.",
	"spec":   "Spec represents a desired configuration of work to be deployed on the spoke cluster.",
	"status": "Status represents the current status of work",
}

func (ManifestWork) SwaggerDoc() map[string]string {
	return map_ManifestWork
}

var map_ManifestWorkList = map[string]string{
	"":         "ManifestWorkList is a collection of manifestworks.",
	"metadata": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
	"items":    "Items is a list of manifestworks.",
}

func (ManifestWorkList) SwaggerDoc() map[string]string {
	return map_ManifestWorkList
}

var map_ManifestWorkSpec = map[string]string{
	"":         "ManifestWorkSpec represents a desired configuration of manifests to be deployed on the spoke cluster.",
	"workload": "Workload represents the manifest workload to be deployed on spoke cluster",
}

func (ManifestWorkSpec) SwaggerDoc() map[string]string {
	return map_ManifestWorkSpec
}

var map_ManifestWorkStatus = map[string]string{
	"":               "ManifestWorkStatus represents the current status of spoke manifest workload",
	"conditions":     "Conditions contains the different condition statuses for this work. Valid condition types are: 1. Applied represents workload in ManifestWork is applied successfully on spoke cluster. 2. Progressing represents workload in ManifestWork is being applied on spoke cluster. 3. Available represents workload in ManifestWork exists on the spoke cluster. 4. Degraded represents the current state of workload does not match the desired state for a certain period.",
	"resourceStatus": "ResourceStatus represents the status of each resource in manifestwork deployed on spoke cluster. The agent on spoke cluster syncs the condition from spoke to the hub.",
}

func (ManifestWorkStatus) SwaggerDoc() map[string]string {
	return map_ManifestWorkStatus
}

var map_ManifestsTemplate = map[string]string{
	"":          "ManifestsTemplate represents the manifest workload to be deployed on spoke cluster",
	"manifests": "Manifests represents a list of kuberenetes resources to be deployed on the spoke cluster.",
}

func (ManifestsTemplate) SwaggerDoc() map[string]string {
	return map_ManifestsTemplate
}

var map_StatusCondition = map[string]string{
	"":                   "StatusCondition contains condition information for a spoke work.",
	"type":               "Type is the type of the spoke work condition.",
	"status":             "Status is the status of the condition. One of True, False, Unknown.",
	"lastTransitionTime": "LastTransitionTime is the last time the condition changed from one status to another.",
	"reason":             "Reason is a (brief) reason for the condition's last status change.",
	"message":            "Message is a human-readable message indicating details about the last status change.",
}

func (StatusCondition) SwaggerDoc() map[string]string {
	return map_StatusCondition
}

// AUTO-GENERATED FUNCTIONS END HERE