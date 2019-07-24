// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/bullblock-io/tezTracker/gen/restapi/operations/accounts"
	"github.com/bullblock-io/tezTracker/gen/restapi/operations/app_info"
	"github.com/bullblock-io/tezTracker/gen/restapi/operations/blocks"
	"github.com/bullblock-io/tezTracker/gen/restapi/operations/fees"
	"github.com/bullblock-io/tezTracker/gen/restapi/operations/metadata"
	"github.com/bullblock-io/tezTracker/gen/restapi/operations/operation_groups"
	"github.com/bullblock-io/tezTracker/gen/restapi/operations/operations_list"
	"github.com/bullblock-io/tezTracker/gen/restapi/operations/query"
)

// NewTezTrackerAPI creates a new TezTracker instance
func NewTezTrackerAPI(spec *loads.Document) *TezTrackerAPI {
	return &TezTrackerAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		CsvProducer: runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
			return errors.NotImplemented("csv producer has not yet been implemented")
		}),
		TxtProducer: runtime.TextProducer(),
		MetadataGetV2MetadataPlatformNetworkEntitiesHandler: metadata.GetV2MetadataPlatformNetworkEntitiesHandlerFunc(func(params metadata.GetV2MetadataPlatformNetworkEntitiesParams) middleware.Responder {
			return middleware.NotImplemented("operation MetadataGetV2MetadataPlatformNetworkEntities has not yet been implemented")
		}),
		MetadataGetV2MetadataPlatformNetworkEntityAttributeHandler: metadata.GetV2MetadataPlatformNetworkEntityAttributeHandlerFunc(func(params metadata.GetV2MetadataPlatformNetworkEntityAttributeParams) middleware.Responder {
			return middleware.NotImplemented("operation MetadataGetV2MetadataPlatformNetworkEntityAttribute has not yet been implemented")
		}),
		MetadataGetV2MetadataPlatformNetworkEntityAttributeFilterHandler: metadata.GetV2MetadataPlatformNetworkEntityAttributeFilterHandlerFunc(func(params metadata.GetV2MetadataPlatformNetworkEntityAttributeFilterParams) middleware.Responder {
			return middleware.NotImplemented("operation MetadataGetV2MetadataPlatformNetworkEntityAttributeFilter has not yet been implemented")
		}),
		MetadataGetV2MetadataPlatformNetworkEntityAttributesHandler: metadata.GetV2MetadataPlatformNetworkEntityAttributesHandlerFunc(func(params metadata.GetV2MetadataPlatformNetworkEntityAttributesParams) middleware.Responder {
			return middleware.NotImplemented("operation MetadataGetV2MetadataPlatformNetworkEntityAttributes has not yet been implemented")
		}),
		MetadataGetV2MetadataPlatformNetworksHandler: metadata.GetV2MetadataPlatformNetworksHandlerFunc(func(params metadata.GetV2MetadataPlatformNetworksParams) middleware.Responder {
			return middleware.NotImplemented("operation MetadataGetV2MetadataPlatformNetworks has not yet been implemented")
		}),
		MetadataGetV2MetadataPlatformsHandler: metadata.GetV2MetadataPlatformsHandlerFunc(func(params metadata.GetV2MetadataPlatformsParams) middleware.Responder {
			return middleware.NotImplemented("operation MetadataGetV2MetadataPlatforms has not yet been implemented")
		}),
		QueryPostV2DataPlatformNetworkEntityHandler: query.PostV2DataPlatformNetworkEntityHandlerFunc(func(params query.PostV2DataPlatformNetworkEntityParams) middleware.Responder {
			return middleware.NotImplemented("operation QueryPostV2DataPlatformNetworkEntity has not yet been implemented")
		}),
		AccountsGetAccountHandler: accounts.GetAccountHandlerFunc(func(params accounts.GetAccountParams) middleware.Responder {
			return middleware.NotImplemented("operation AccountsGetAccount has not yet been implemented")
		}),
		AccountsGetAccountsListHandler: accounts.GetAccountsListHandlerFunc(func(params accounts.GetAccountsListParams) middleware.Responder {
			return middleware.NotImplemented("operation AccountsGetAccountsList has not yet been implemented")
		}),
		FeesGetAvgFeesHandler: fees.GetAvgFeesHandlerFunc(func(params fees.GetAvgFeesParams) middleware.Responder {
			return middleware.NotImplemented("operation FeesGetAvgFees has not yet been implemented")
		}),
		AccountsGetBakersListHandler: accounts.GetBakersListHandlerFunc(func(params accounts.GetBakersListParams) middleware.Responder {
			return middleware.NotImplemented("operation AccountsGetBakersList has not yet been implemented")
		}),
		BlocksGetBlockHandler: blocks.GetBlockHandlerFunc(func(params blocks.GetBlockParams) middleware.Responder {
			return middleware.NotImplemented("operation BlocksGetBlock has not yet been implemented")
		}),
		BlocksGetBlockEndorsementsHandler: blocks.GetBlockEndorsementsHandlerFunc(func(params blocks.GetBlockEndorsementsParams) middleware.Responder {
			return middleware.NotImplemented("operation BlocksGetBlockEndorsements has not yet been implemented")
		}),
		BlocksGetBlocksHeadHandler: blocks.GetBlocksHeadHandlerFunc(func(params blocks.GetBlocksHeadParams) middleware.Responder {
			return middleware.NotImplemented("operation BlocksGetBlocksHead has not yet been implemented")
		}),
		BlocksGetBlocksListHandler: blocks.GetBlocksListHandlerFunc(func(params blocks.GetBlocksListParams) middleware.Responder {
			return middleware.NotImplemented("operation BlocksGetBlocksList has not yet been implemented")
		}),
		AppInfoGetInfoHandler: app_info.GetInfoHandlerFunc(func(params app_info.GetInfoParams) middleware.Responder {
			return middleware.NotImplemented("operation AppInfoGetInfo has not yet been implemented")
		}),
		OperationGroupsGetOperationGroupHandler: operation_groups.GetOperationGroupHandlerFunc(func(params operation_groups.GetOperationGroupParams) middleware.Responder {
			return middleware.NotImplemented("operation OperationGroupsGetOperationGroup has not yet been implemented")
		}),
		OperationGroupsGetOperationGroupsHandler: operation_groups.GetOperationGroupsHandlerFunc(func(params operation_groups.GetOperationGroupsParams) middleware.Responder {
			return middleware.NotImplemented("operation OperationGroupsGetOperationGroups has not yet been implemented")
		}),
		OperationsListGetOperationsListHandler: operations_list.GetOperationsListHandlerFunc(func(params operations_list.GetOperationsListParams) middleware.Responder {
			return middleware.NotImplemented("operation OperationsListGetOperationsList has not yet been implemented")
		}),
	}
}

/*TezTrackerAPI the tez tracker API */
type TezTrackerAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer
	// CsvProducer registers a producer for a "text/csv" mime type
	CsvProducer runtime.Producer
	// TxtProducer registers a producer for a "text/plain" mime type
	TxtProducer runtime.Producer

	// MetadataGetV2MetadataPlatformNetworkEntitiesHandler sets the operation handler for the get v2 metadata platform network entities operation
	MetadataGetV2MetadataPlatformNetworkEntitiesHandler metadata.GetV2MetadataPlatformNetworkEntitiesHandler
	// MetadataGetV2MetadataPlatformNetworkEntityAttributeHandler sets the operation handler for the get v2 metadata platform network entity attribute operation
	MetadataGetV2MetadataPlatformNetworkEntityAttributeHandler metadata.GetV2MetadataPlatformNetworkEntityAttributeHandler
	// MetadataGetV2MetadataPlatformNetworkEntityAttributeFilterHandler sets the operation handler for the get v2 metadata platform network entity attribute filter operation
	MetadataGetV2MetadataPlatformNetworkEntityAttributeFilterHandler metadata.GetV2MetadataPlatformNetworkEntityAttributeFilterHandler
	// MetadataGetV2MetadataPlatformNetworkEntityAttributesHandler sets the operation handler for the get v2 metadata platform network entity attributes operation
	MetadataGetV2MetadataPlatformNetworkEntityAttributesHandler metadata.GetV2MetadataPlatformNetworkEntityAttributesHandler
	// MetadataGetV2MetadataPlatformNetworksHandler sets the operation handler for the get v2 metadata platform networks operation
	MetadataGetV2MetadataPlatformNetworksHandler metadata.GetV2MetadataPlatformNetworksHandler
	// MetadataGetV2MetadataPlatformsHandler sets the operation handler for the get v2 metadata platforms operation
	MetadataGetV2MetadataPlatformsHandler metadata.GetV2MetadataPlatformsHandler
	// QueryPostV2DataPlatformNetworkEntityHandler sets the operation handler for the post v2 data platform network entity operation
	QueryPostV2DataPlatformNetworkEntityHandler query.PostV2DataPlatformNetworkEntityHandler
	// AccountsGetAccountHandler sets the operation handler for the get account operation
	AccountsGetAccountHandler accounts.GetAccountHandler
	// AccountsGetAccountsListHandler sets the operation handler for the get accounts list operation
	AccountsGetAccountsListHandler accounts.GetAccountsListHandler
	// FeesGetAvgFeesHandler sets the operation handler for the get avg fees operation
	FeesGetAvgFeesHandler fees.GetAvgFeesHandler
	// AccountsGetBakersListHandler sets the operation handler for the get bakers list operation
	AccountsGetBakersListHandler accounts.GetBakersListHandler
	// BlocksGetBlockHandler sets the operation handler for the get block operation
	BlocksGetBlockHandler blocks.GetBlockHandler
	// BlocksGetBlockEndorsementsHandler sets the operation handler for the get block endorsements operation
	BlocksGetBlockEndorsementsHandler blocks.GetBlockEndorsementsHandler
	// BlocksGetBlocksHeadHandler sets the operation handler for the get blocks head operation
	BlocksGetBlocksHeadHandler blocks.GetBlocksHeadHandler
	// BlocksGetBlocksListHandler sets the operation handler for the get blocks list operation
	BlocksGetBlocksListHandler blocks.GetBlocksListHandler
	// AppInfoGetInfoHandler sets the operation handler for the get info operation
	AppInfoGetInfoHandler app_info.GetInfoHandler
	// OperationGroupsGetOperationGroupHandler sets the operation handler for the get operation group operation
	OperationGroupsGetOperationGroupHandler operation_groups.GetOperationGroupHandler
	// OperationGroupsGetOperationGroupsHandler sets the operation handler for the get operation groups operation
	OperationGroupsGetOperationGroupsHandler operation_groups.GetOperationGroupsHandler
	// OperationsListGetOperationsListHandler sets the operation handler for the get operations list operation
	OperationsListGetOperationsListHandler operations_list.GetOperationsListHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *TezTrackerAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *TezTrackerAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *TezTrackerAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *TezTrackerAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *TezTrackerAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *TezTrackerAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *TezTrackerAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the TezTrackerAPI
func (o *TezTrackerAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.CsvProducer == nil {
		unregistered = append(unregistered, "CsvProducer")
	}

	if o.TxtProducer == nil {
		unregistered = append(unregistered, "TxtProducer")
	}

	if o.MetadataGetV2MetadataPlatformNetworkEntitiesHandler == nil {
		unregistered = append(unregistered, "metadata.GetV2MetadataPlatformNetworkEntitiesHandler")
	}

	if o.MetadataGetV2MetadataPlatformNetworkEntityAttributeHandler == nil {
		unregistered = append(unregistered, "metadata.GetV2MetadataPlatformNetworkEntityAttributeHandler")
	}

	if o.MetadataGetV2MetadataPlatformNetworkEntityAttributeFilterHandler == nil {
		unregistered = append(unregistered, "metadata.GetV2MetadataPlatformNetworkEntityAttributeFilterHandler")
	}

	if o.MetadataGetV2MetadataPlatformNetworkEntityAttributesHandler == nil {
		unregistered = append(unregistered, "metadata.GetV2MetadataPlatformNetworkEntityAttributesHandler")
	}

	if o.MetadataGetV2MetadataPlatformNetworksHandler == nil {
		unregistered = append(unregistered, "metadata.GetV2MetadataPlatformNetworksHandler")
	}

	if o.MetadataGetV2MetadataPlatformsHandler == nil {
		unregistered = append(unregistered, "metadata.GetV2MetadataPlatformsHandler")
	}

	if o.QueryPostV2DataPlatformNetworkEntityHandler == nil {
		unregistered = append(unregistered, "query.PostV2DataPlatformNetworkEntityHandler")
	}

	if o.AccountsGetAccountHandler == nil {
		unregistered = append(unregistered, "accounts.GetAccountHandler")
	}

	if o.AccountsGetAccountsListHandler == nil {
		unregistered = append(unregistered, "accounts.GetAccountsListHandler")
	}

	if o.FeesGetAvgFeesHandler == nil {
		unregistered = append(unregistered, "fees.GetAvgFeesHandler")
	}

	if o.AccountsGetBakersListHandler == nil {
		unregistered = append(unregistered, "accounts.GetBakersListHandler")
	}

	if o.BlocksGetBlockHandler == nil {
		unregistered = append(unregistered, "blocks.GetBlockHandler")
	}

	if o.BlocksGetBlockEndorsementsHandler == nil {
		unregistered = append(unregistered, "blocks.GetBlockEndorsementsHandler")
	}

	if o.BlocksGetBlocksHeadHandler == nil {
		unregistered = append(unregistered, "blocks.GetBlocksHeadHandler")
	}

	if o.BlocksGetBlocksListHandler == nil {
		unregistered = append(unregistered, "blocks.GetBlocksListHandler")
	}

	if o.AppInfoGetInfoHandler == nil {
		unregistered = append(unregistered, "app_info.GetInfoHandler")
	}

	if o.OperationGroupsGetOperationGroupHandler == nil {
		unregistered = append(unregistered, "operation_groups.GetOperationGroupHandler")
	}

	if o.OperationGroupsGetOperationGroupsHandler == nil {
		unregistered = append(unregistered, "operation_groups.GetOperationGroupsHandler")
	}

	if o.OperationsListGetOperationsListHandler == nil {
		unregistered = append(unregistered, "operations_list.GetOperationsListHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *TezTrackerAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *TezTrackerAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *TezTrackerAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *TezTrackerAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *TezTrackerAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		case "text/csv":
			result["text/csv"] = o.CsvProducer

		case "text/plain":
			result["text/plain"] = o.TxtProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *TezTrackerAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the tez tracker API
func (o *TezTrackerAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *TezTrackerAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/metadata/{platform}/{network}/entities"] = metadata.NewGetV2MetadataPlatformNetworkEntities(o.context, o.MetadataGetV2MetadataPlatformNetworkEntitiesHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/metadata/{platform}/{network}/{entity}/{attribute}"] = metadata.NewGetV2MetadataPlatformNetworkEntityAttribute(o.context, o.MetadataGetV2MetadataPlatformNetworkEntityAttributeHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/metadata/{platform}/{network}/{entity}/{attribute}/{filter}"] = metadata.NewGetV2MetadataPlatformNetworkEntityAttributeFilter(o.context, o.MetadataGetV2MetadataPlatformNetworkEntityAttributeFilterHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/metadata/{platform}/{network}/{entity}/attributes"] = metadata.NewGetV2MetadataPlatformNetworkEntityAttributes(o.context, o.MetadataGetV2MetadataPlatformNetworkEntityAttributesHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/metadata/{platform}/networks"] = metadata.NewGetV2MetadataPlatformNetworks(o.context, o.MetadataGetV2MetadataPlatformNetworksHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/metadata/platforms"] = metadata.NewGetV2MetadataPlatforms(o.context, o.MetadataGetV2MetadataPlatformsHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v2/data/{platform}/{network}/{entity}"] = query.NewPostV2DataPlatformNetworkEntity(o.context, o.QueryPostV2DataPlatformNetworkEntityHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/data/{platform}/{network}/accounts/{accountId}"] = accounts.NewGetAccount(o.context, o.AccountsGetAccountHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/data/{platform}/{network}/accounts"] = accounts.NewGetAccountsList(o.context, o.AccountsGetAccountsListHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/data/{platform}/{network}/operations/avgFees"] = fees.NewGetAvgFees(o.context, o.FeesGetAvgFeesHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/data/{platform}/{network}/bakers"] = accounts.NewGetBakersList(o.context, o.AccountsGetBakersListHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/data/{platform}/{network}/blocks/{hash}"] = blocks.NewGetBlock(o.context, o.BlocksGetBlockHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/data/{platform}/{network}/blocks/{hash}/endorsements"] = blocks.NewGetBlockEndorsements(o.context, o.BlocksGetBlockEndorsementsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/data/{platform}/{network}/blocks/head"] = blocks.NewGetBlocksHead(o.context, o.BlocksGetBlocksHeadHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/data/{platform}/{network}/blocks"] = blocks.NewGetBlocksList(o.context, o.BlocksGetBlocksListHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/data/{platform}/{network}/info"] = app_info.NewGetInfo(o.context, o.AppInfoGetInfoHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/data/{platform}/{network}/operation_groups/{operationGroupId}"] = operation_groups.NewGetOperationGroup(o.context, o.OperationGroupsGetOperationGroupHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/data/{platform}/{network}/operation_groups"] = operation_groups.NewGetOperationGroups(o.context, o.OperationGroupsGetOperationGroupsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/data/{platform}/{network}/operations"] = operations_list.NewGetOperationsList(o.context, o.OperationsListGetOperationsListHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *TezTrackerAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *TezTrackerAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *TezTrackerAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *TezTrackerAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
