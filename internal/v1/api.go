// Package v1 provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package v1

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// AWSS3UploadRequestOptions defines model for AWSS3UploadRequestOptions.
type AWSS3UploadRequestOptions map[string]interface{}

// AWSS3UploadStatus defines model for AWSS3UploadStatus.
type AWSS3UploadStatus struct {
	Url string `json:"url"`
}

// AWSUploadRequestOptions defines model for AWSUploadRequestOptions.
type AWSUploadRequestOptions struct {
	ShareWithAccounts []string `json:"share_with_accounts"`
}

// AWSUploadStatus defines model for AWSUploadStatus.
type AWSUploadStatus struct {
	Ami    string `json:"ami"`
	Region string `json:"region"`
}

// ArchitectureItem defines model for ArchitectureItem.
type ArchitectureItem struct {
	Arch       string   `json:"arch"`
	ImageTypes []string `json:"image_types"`
}

// Architectures defines model for Architectures.
type Architectures []ArchitectureItem

// AzureUploadRequestOptions defines model for AzureUploadRequestOptions.
type AzureUploadRequestOptions struct {

	// Name of the resource group where the image should be uploaded.
	ResourceGroup string `json:"resource_group"`

	// ID of subscription where the image should be uploaded.
	SubscriptionId string `json:"subscription_id"`

	// ID of the tenant where the image should be uploaded. This link explains how
	// to find it in the Azure Portal:
	// https://docs.microsoft.com/en-us/azure/active-directory/fundamentals/active-directory-how-to-find-tenant
	TenantId string `json:"tenant_id"`
}

// AzureUploadStatus defines model for AzureUploadStatus.
type AzureUploadStatus struct {
	ImageName string `json:"image_name"`
}

// ComposeMetadata defines model for ComposeMetadata.
type ComposeMetadata struct {

	// ID (hash) of the built commit
	OstreeCommit *string `json:"ostree_commit,omitempty"`

	// Package list including NEVRA
	Packages *[]PackageMetadata `json:"packages,omitempty"`
}

// ComposeRequest defines model for ComposeRequest.
type ComposeRequest struct {
	Customizations *Customizations `json:"customizations,omitempty"`
	Distribution   Distributions   `json:"distribution"`
	ImageName      *string         `json:"image_name,omitempty"`

	// Array of exactly one image request. Having more image requests in one compose is currently not supported.
	ImageRequests []ImageRequest `json:"image_requests"`
}

// ComposeResponse defines model for ComposeResponse.
type ComposeResponse struct {
	Id string `json:"id"`
}

// ComposeStatus defines model for ComposeStatus.
type ComposeStatus struct {
	ImageStatus ImageStatus `json:"image_status"`
}

// ComposeStatusError defines model for ComposeStatusError.
type ComposeStatusError struct {
	Details *interface{} `json:"details,omitempty"`
	Id      int          `json:"id"`
	Reason  string       `json:"reason"`
}

// ComposesResponse defines model for ComposesResponse.
type ComposesResponse struct {
	Data  []ComposesResponseItem `json:"data"`
	Links struct {
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"links"`
	Meta struct {
		Count int `json:"count"`
	} `json:"meta"`
}

// ComposesResponseItem defines model for ComposesResponseItem.
type ComposesResponseItem struct {
	CreatedAt string      `json:"created_at"`
	Id        string      `json:"id"`
	ImageName *string     `json:"image_name,omitempty"`
	Request   interface{} `json:"request"`
}

// Customizations defines model for Customizations.
type Customizations struct {
	Filesystem          *[]Filesystem `json:"filesystem,omitempty"`
	Packages            *[]string     `json:"packages,omitempty"`
	PayloadRepositories *[]Repository `json:"payload_repositories,omitempty"`
	Subscription        *Subscription `json:"subscription,omitempty"`
}

// DistributionItem defines model for DistributionItem.
type DistributionItem struct {
	Description string `json:"description"`
	Name        string `json:"name"`
}

// Distributions defines model for Distributions.
type Distributions string

// List of Distributions
const (
	Distributions_centos_8 Distributions = "centos-8"
	Distributions_centos_9 Distributions = "centos-9"
	Distributions_rhel_84  Distributions = "rhel-84"
	Distributions_rhel_85  Distributions = "rhel-85"
	Distributions_rhel_86  Distributions = "rhel-86"
	Distributions_rhel_90  Distributions = "rhel-90"
)

// DistributionsResponse defines model for DistributionsResponse.
type DistributionsResponse []DistributionItem

// Filesystem defines model for Filesystem.
type Filesystem struct {
	MinSize    uint64 `json:"min_size"`
	Mountpoint string `json:"mountpoint"`
}

// GCPUploadRequestOptions defines model for GCPUploadRequestOptions.
type GCPUploadRequestOptions struct {

	// List of valid Google accounts to share the imported Compute Node image with.
	// Each string must contain a specifier of the account type. Valid formats are:
	//   - 'user:{emailid}': An email address that represents a specific
	//     Google account. For example, 'alice@example.com'.
	//   - 'serviceAccount:{emailid}': An email address that represents a
	//     service account. For example, 'my-other-app@appspot.gserviceaccount.com'.
	//   - 'group:{emailid}': An email address that represents a Google group.
	//     For example, 'admins@example.com'.
	//   - 'domain:{domain}': The G Suite domain (primary) that represents all
	//     the users of that domain. For example, 'google.com' or 'example.com'.
	//     If not specified, the imported Compute Node image is not shared with any
	//     account.
	ShareWithAccounts []string `json:"share_with_accounts"`
}

// GCPUploadStatus defines model for GCPUploadStatus.
type GCPUploadStatus struct {
	ImageName string `json:"image_name"`
	ProjectId string `json:"project_id"`
}

// HTTPError defines model for HTTPError.
type HTTPError struct {
	Detail string `json:"detail"`
	Title  string `json:"title"`
}

// HTTPErrorList defines model for HTTPErrorList.
type HTTPErrorList struct {
	Errors []HTTPError `json:"errors"`
}

// ImageRequest defines model for ImageRequest.
type ImageRequest struct {

	// CPU architecture of the image, only x86_64 is currently supported.
	Architecture  string        `json:"architecture"`
	ImageType     ImageTypes    `json:"image_type"`
	Ostree        *OSTree       `json:"ostree,omitempty"`
	UploadRequest UploadRequest `json:"upload_request"`
}

// ImageStatus defines model for ImageStatus.
type ImageStatus struct {
	Error        *ComposeStatusError `json:"error,omitempty"`
	Status       string              `json:"status"`
	UploadStatus *UploadStatus       `json:"upload_status,omitempty"`
}

// ImageTypes defines model for ImageTypes.
type ImageTypes string

// List of ImageTypes
const (
	ImageTypes_ami                 ImageTypes = "ami"
	ImageTypes_aws                 ImageTypes = "aws"
	ImageTypes_azure               ImageTypes = "azure"
	ImageTypes_edge_commit         ImageTypes = "edge-commit"
	ImageTypes_edge_container      ImageTypes = "edge-container"
	ImageTypes_edge_installer      ImageTypes = "edge-installer"
	ImageTypes_gcp                 ImageTypes = "gcp"
	ImageTypes_guest_image         ImageTypes = "guest-image"
	ImageTypes_image_installer     ImageTypes = "image-installer"
	ImageTypes_rhel_edge_commit    ImageTypes = "rhel-edge-commit"
	ImageTypes_rhel_edge_installer ImageTypes = "rhel-edge-installer"
	ImageTypes_vhd                 ImageTypes = "vhd"
	ImageTypes_vsphere             ImageTypes = "vsphere"
)

// OSTree defines model for OSTree.
type OSTree struct {

	// Can be either a commit (example: 02604b2da6e954bd34b8b82a835e5a77d2b60ffa), or a branch-like reference (example: rhel/8/x86_64/edge)
	Parent *string `json:"parent,omitempty"`
	Ref    *string `json:"ref,omitempty"`
	Url    *string `json:"url,omitempty"`
}

// Package defines model for Package.
type Package struct {
	Name    string `json:"name"`
	Summary string `json:"summary"`
}

// PackageMetadata defines model for PackageMetadata.
type PackageMetadata struct {
	Arch      string  `json:"arch"`
	Epoch     *string `json:"epoch,omitempty"`
	Name      string  `json:"name"`
	Release   string  `json:"release"`
	Sigmd5    string  `json:"sigmd5"`
	Signature *string `json:"signature,omitempty"`
	Type      string  `json:"type"`
	Version   string  `json:"version"`
}

// PackagesResponse defines model for PackagesResponse.
type PackagesResponse struct {
	Data  []Package `json:"data"`
	Links struct {
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"links"`
	Meta struct {
		Count int `json:"count"`
	} `json:"meta"`
}

// Readiness defines model for Readiness.
type Readiness struct {
	Readiness string `json:"readiness"`
}

// Repository defines model for Repository.
type Repository struct {
	Baseurl    *string `json:"baseurl,omitempty"`
	CheckGpg   *bool   `json:"check_gpg,omitempty"`
	Gpgkey     *string `json:"gpgkey,omitempty"`
	IgnoreSsl  *bool   `json:"ignore_ssl,omitempty"`
	Metalink   *string `json:"metalink,omitempty"`
	Mirrorlist *string `json:"mirrorlist,omitempty"`
	Rhsm       bool    `json:"rhsm"`
}

// Subscription defines model for Subscription.
type Subscription struct {
	ActivationKey string `json:"activation-key"`
	BaseUrl       string `json:"base-url"`
	Insights      bool   `json:"insights"`
	Organization  int    `json:"organization"`
	ServerUrl     string `json:"server-url"`
}

// UploadRequest defines model for UploadRequest.
type UploadRequest struct {
	Options interface{} `json:"options"`
	Type    UploadTypes `json:"type"`
}

// UploadStatus defines model for UploadStatus.
type UploadStatus struct {
	Options interface{} `json:"options"`
	Status  string      `json:"status"`
	Type    UploadTypes `json:"type"`
}

// UploadTypes defines model for UploadTypes.
type UploadTypes string

// List of UploadTypes
const (
	UploadTypes_aws    UploadTypes = "aws"
	UploadTypes_aws_s3 UploadTypes = "aws.s3"
	UploadTypes_azure  UploadTypes = "azure"
	UploadTypes_gcp    UploadTypes = "gcp"
)

// Version defines model for Version.
type Version struct {
	Version string `json:"version"`
}

// ComposeImageJSONBody defines parameters for ComposeImage.
type ComposeImageJSONBody ComposeRequest

// GetComposesParams defines parameters for GetComposes.
type GetComposesParams struct {

	// max amount of composes, default 100
	Limit *int `json:"limit,omitempty"`

	// composes page offset, default 0
	Offset *int `json:"offset,omitempty"`
}

// GetPackagesParams defines parameters for GetPackages.
type GetPackagesParams struct {

	// distribution to look up packages for
	Distribution Distributions `json:"distribution"`

	// architecture to look up packages for
	Architecture string `json:"architecture"`

	// packages to look for
	Search string `json:"search"`

	// max amount of packages, default 100
	Limit *int `json:"limit,omitempty"`

	// packages page offset, default 0
	Offset *int `json:"offset,omitempty"`
}

// ComposeImageRequestBody defines body for ComposeImage for application/json ContentType.
type ComposeImageJSONRequestBody ComposeImageJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// get the architectures and their image types available for a given distribution
	// (GET /architectures/{distribution})
	GetArchitectures(ctx echo.Context, distribution string) error
	// compose image
	// (POST /compose)
	ComposeImage(ctx echo.Context) error
	// get a collection of previous compose requests for the logged in user
	// (GET /composes)
	GetComposes(ctx echo.Context, params GetComposesParams) error
	// get status of an image compose
	// (GET /composes/{composeId})
	GetComposeStatus(ctx echo.Context, composeId string) error
	// get metadata of an image compose
	// (GET /composes/{composeId}/metadata)
	GetComposeMetadata(ctx echo.Context, composeId string) error
	// get the available distributions
	// (GET /distributions)
	GetDistributions(ctx echo.Context) error
	// get the openapi json specification
	// (GET /openapi.json)
	GetOpenapiJson(ctx echo.Context) error

	// (GET /packages)
	GetPackages(ctx echo.Context, params GetPackagesParams) error
	// return the readiness
	// (GET /ready)
	GetReadiness(ctx echo.Context) error
	// get the service version
	// (GET /version)
	GetVersion(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetArchitectures converts echo context to params.
func (w *ServerInterfaceWrapper) GetArchitectures(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "distribution" -------------
	var distribution string

	err = runtime.BindStyledParameter("simple", false, "distribution", ctx.Param("distribution"), &distribution)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter distribution: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetArchitectures(ctx, distribution)
	return err
}

// ComposeImage converts echo context to params.
func (w *ServerInterfaceWrapper) ComposeImage(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ComposeImage(ctx)
	return err
}

// GetComposes converts echo context to params.
func (w *ServerInterfaceWrapper) GetComposes(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetComposesParams
	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetComposes(ctx, params)
	return err
}

// GetComposeStatus converts echo context to params.
func (w *ServerInterfaceWrapper) GetComposeStatus(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "composeId" -------------
	var composeId string

	err = runtime.BindStyledParameter("simple", false, "composeId", ctx.Param("composeId"), &composeId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter composeId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetComposeStatus(ctx, composeId)
	return err
}

// GetComposeMetadata converts echo context to params.
func (w *ServerInterfaceWrapper) GetComposeMetadata(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "composeId" -------------
	var composeId string

	err = runtime.BindStyledParameter("simple", false, "composeId", ctx.Param("composeId"), &composeId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter composeId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetComposeMetadata(ctx, composeId)
	return err
}

// GetDistributions converts echo context to params.
func (w *ServerInterfaceWrapper) GetDistributions(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetDistributions(ctx)
	return err
}

// GetOpenapiJson converts echo context to params.
func (w *ServerInterfaceWrapper) GetOpenapiJson(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetOpenapiJson(ctx)
	return err
}

// GetPackages converts echo context to params.
func (w *ServerInterfaceWrapper) GetPackages(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPackagesParams
	// ------------- Required query parameter "distribution" -------------

	err = runtime.BindQueryParameter("form", true, true, "distribution", ctx.QueryParams(), &params.Distribution)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter distribution: %s", err))
	}

	// ------------- Required query parameter "architecture" -------------

	err = runtime.BindQueryParameter("form", true, true, "architecture", ctx.QueryParams(), &params.Architecture)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter architecture: %s", err))
	}

	// ------------- Required query parameter "search" -------------

	err = runtime.BindQueryParameter("form", true, true, "search", ctx.QueryParams(), &params.Search)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter search: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPackages(ctx, params)
	return err
}

// GetReadiness converts echo context to params.
func (w *ServerInterfaceWrapper) GetReadiness(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetReadiness(ctx)
	return err
}

// GetVersion converts echo context to params.
func (w *ServerInterfaceWrapper) GetVersion(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetVersion(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET("/architectures/:distribution", wrapper.GetArchitectures)
	router.POST("/compose", wrapper.ComposeImage)
	router.GET("/composes", wrapper.GetComposes)
	router.GET("/composes/:composeId", wrapper.GetComposeStatus)
	router.GET("/composes/:composeId/metadata", wrapper.GetComposeMetadata)
	router.GET("/distributions", wrapper.GetDistributions)
	router.GET("/openapi.json", wrapper.GetOpenapiJson)
	router.GET("/packages", wrapper.GetPackages)
	router.GET("/ready", wrapper.GetReadiness)
	router.GET("/version", wrapper.GetVersion)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
