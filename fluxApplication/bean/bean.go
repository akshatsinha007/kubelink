package bean

import (
	client "github.com/devtron-labs/kubelink/grpc"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	FluxKustomizationGroup   = "kustomize.toolkit.fluxcd.io"
	FluxAppKustomizationKind = "Kustomization"
	FluxKustomizationVersion = "v1"
	AllNamespaces            = ""
	FluxHelmReleaseGroup     = "helm.toolkit.fluxcd.io"
	FluxAppHelmreleaseKind   = "HelmRelease"
	FluxHelmReleaseVersion   = "v2"

	//DevtronCDNamespae            = "devtroncd"
	//ArgoLabelForManagedResources = "app.kubernetes.io/instance"
)
const (
	Destination string = "destination"
	Server      string = "server"
	STATUS      string = "status"
	INVENTORY   string = "inventory"
	ENTRIES     string = "entries"
	ID          string = "id"
	VERSION     string = "v"
)
const (
	FieldSeparator  = "_"
	ColonTranscoded = "__"
)

var GvkForKustomizationFluxApp = schema.GroupVersionKind{
	Group:   FluxKustomizationGroup,
	Kind:    FluxAppKustomizationKind,
	Version: FluxKustomizationVersion,
}

var GvkForHelmreleaseFluxApp = schema.GroupVersionKind{
	Group:   FluxHelmReleaseGroup,
	Kind:    FluxAppHelmreleaseKind,
	Version: FluxHelmReleaseVersion,
}

type FluxApplicationListDto struct {
	ClusterId   int    `json:"clusterId"`
	ClusterName string `json:"clusterName"`
	FluxAppDto  []*FluxApplicationDto
}
type FluxApplicationDto struct {
	Name               string             `json:"appName"`
	HealthStatus       string             `json:"appStatus"`
	SyncStatus         string             `json:"syncStatus"`
	EnvironmentDetails *EnvironmentDetail `json:"environmentDetail"`
	IsKustomizeApp     bool               `json:"isKustomizeApp"`
}
type EnvironmentDetail struct {
	ClusterId   int    `json:"clusterId"`
	ClusterName string `json:"clusterName"`
	Namespace   string `json:"namespace"`
}

type ObjectMetadataCompact struct {
	Id      string
	Version string
}

type FluxpplicationDetailDto struct {
	FluxAppDto *FluxApplicationListDto `json:"fluxAppDto"`
	Manifest   map[string]interface{}  `json:"manifest"`
}

type FluxKustomization struct {
	AppKsDetailDto   *FluxKsAppDetail                    `json:"appDetailDto"`
	Resources        *client.ExternalResourceTreeRequest `json:"fluxResource"`
	Kustomizations   []*FluxKustomization                `json:"kustomizations"`
	FluxHelmReleases []*client.AppDetailRequest          `json:"fluxHelmReleases"`
	ParentKsApp      string                              `json:"parentKsApp"`
	Manifest         unstructured.Unstructured           `json:"manifest"`
}

type FluxResource struct {
	Gvk       schema.GroupVersionKind
	Name      string
	Namespace string
}

type FluxHelmReleases struct {
	ReleaseName string
	Name        string
	Namespace   string
}
type FluxHelmResource struct {
	ReleaseName string
	Name        string
	Namespace   string
}

type FluxKsAppDetail struct {
	Name      string
	Namespace string
	GroupKind schema.GroupKind
}