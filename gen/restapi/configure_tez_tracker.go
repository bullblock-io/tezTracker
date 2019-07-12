// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"io"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/bullblock-io/tezTracker/gen/restapi/operations"
	"github.com/bullblock-io/tezTracker/gen/restapi/operations/accounts"
	"github.com/bullblock-io/tezTracker/gen/restapi/operations/app_info"
	"github.com/bullblock-io/tezTracker/gen/restapi/operations/blocks"
	"github.com/bullblock-io/tezTracker/gen/restapi/operations/fees"
	"github.com/bullblock-io/tezTracker/gen/restapi/operations/metadata"
	"github.com/bullblock-io/tezTracker/gen/restapi/operations/operation_groups"
	"github.com/bullblock-io/tezTracker/gen/restapi/operations/operations"
	"github.com/bullblock-io/tezTracker/gen/restapi/operations/query"
)

//go:generate swagger generate server --target ../../gen --name TezTracker --spec ../../swagger/swager.yml --exclude-main

func configureFlags(api *operations.TezTrackerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TezTrackerAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.CsvProducer = runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
		return errors.NotImplemented("csv producer has not yet been implemented")
	})
	api.TxtProducer = runtime.TextProducer()

	if api.AppInfoGetInfoHandler == nil {
		api.AppInfoGetInfoHandler = app_info.GetInfoHandlerFunc(func(params app_info.GetInfoParams) middleware.Responder {
			return middleware.NotImplemented("operation app_info.GetInfo has not yet been implemented")
		})
	}
	if api.AccountsGetV2DataPlatformNetworkAccountsHandler == nil {
		api.AccountsGetV2DataPlatformNetworkAccountsHandler = accounts.GetV2DataPlatformNetworkAccountsHandlerFunc(func(params accounts.GetV2DataPlatformNetworkAccountsParams) middleware.Responder {
			return middleware.NotImplemented("operation accounts.GetV2DataPlatformNetworkAccounts has not yet been implemented")
		})
	}
	if api.AccountsGetV2DataPlatformNetworkAccountsAccountIDHandler == nil {
		api.AccountsGetV2DataPlatformNetworkAccountsAccountIDHandler = accounts.GetV2DataPlatformNetworkAccountsAccountIDHandlerFunc(func(params accounts.GetV2DataPlatformNetworkAccountsAccountIDParams) middleware.Responder {
			return middleware.NotImplemented("operation accounts.GetV2DataPlatformNetworkAccountsAccountID has not yet been implemented")
		})
	}
	if api.BlocksGetV2DataPlatformNetworkBlocksHandler == nil {
		api.BlocksGetV2DataPlatformNetworkBlocksHandler = blocks.GetV2DataPlatformNetworkBlocksHandlerFunc(func(params blocks.GetV2DataPlatformNetworkBlocksParams) middleware.Responder {
			return middleware.NotImplemented("operation blocks.GetV2DataPlatformNetworkBlocks has not yet been implemented")
		})
	}
	if api.BlocksGetV2DataPlatformNetworkBlocksHashHandler == nil {
		api.BlocksGetV2DataPlatformNetworkBlocksHashHandler = blocks.GetV2DataPlatformNetworkBlocksHashHandlerFunc(func(params blocks.GetV2DataPlatformNetworkBlocksHashParams) middleware.Responder {
			return middleware.NotImplemented("operation blocks.GetV2DataPlatformNetworkBlocksHash has not yet been implemented")
		})
	}
	if api.BlocksGetV2DataPlatformNetworkBlocksHeadHandler == nil {
		api.BlocksGetV2DataPlatformNetworkBlocksHeadHandler = blocks.GetV2DataPlatformNetworkBlocksHeadHandlerFunc(func(params blocks.GetV2DataPlatformNetworkBlocksHeadParams) middleware.Responder {
			return middleware.NotImplemented("operation blocks.GetV2DataPlatformNetworkBlocksHead has not yet been implemented")
		})
	}
	if api.OperationGroupsGetV2DataPlatformNetworkOperationGroupsHandler == nil {
		api.OperationGroupsGetV2DataPlatformNetworkOperationGroupsHandler = operation_groups.GetV2DataPlatformNetworkOperationGroupsHandlerFunc(func(params operation_groups.GetV2DataPlatformNetworkOperationGroupsParams) middleware.Responder {
			return middleware.NotImplemented("operation operation_groups.GetV2DataPlatformNetworkOperationGroups has not yet been implemented")
		})
	}
	if api.OperationGroupsGetV2DataPlatformNetworkOperationGroupsOperationGroupIDHandler == nil {
		api.OperationGroupsGetV2DataPlatformNetworkOperationGroupsOperationGroupIDHandler = operation_groups.GetV2DataPlatformNetworkOperationGroupsOperationGroupIDHandlerFunc(func(params operation_groups.GetV2DataPlatformNetworkOperationGroupsOperationGroupIDParams) middleware.Responder {
			return middleware.NotImplemented("operation operation_groups.GetV2DataPlatformNetworkOperationGroupsOperationGroupID has not yet been implemented")
		})
	}
	if api.GetV2DataPlatformNetworkOperationsHandler == nil {
		api.GetV2DataPlatformNetworkOperationsHandler = operations.GetV2DataPlatformNetworkOperationsHandlerFunc(func(params operations.GetV2DataPlatformNetworkOperationsParams) middleware.Responder {
			return middleware.NotImplemented("operation .GetV2DataPlatformNetworkOperations has not yet been implemented")
		})
	}
	if api.FeesGetV2DataPlatformNetworkOperationsAvgFeesHandler == nil {
		api.FeesGetV2DataPlatformNetworkOperationsAvgFeesHandler = fees.GetV2DataPlatformNetworkOperationsAvgFeesHandlerFunc(func(params fees.GetV2DataPlatformNetworkOperationsAvgFeesParams) middleware.Responder {
			return middleware.NotImplemented("operation fees.GetV2DataPlatformNetworkOperationsAvgFees has not yet been implemented")
		})
	}
	if api.MetadataGetV2MetadataPlatformNetworkEntitiesHandler == nil {
		api.MetadataGetV2MetadataPlatformNetworkEntitiesHandler = metadata.GetV2MetadataPlatformNetworkEntitiesHandlerFunc(func(params metadata.GetV2MetadataPlatformNetworkEntitiesParams) middleware.Responder {
			return middleware.NotImplemented("operation metadata.GetV2MetadataPlatformNetworkEntities has not yet been implemented")
		})
	}
	if api.MetadataGetV2MetadataPlatformNetworkEntityAttributeHandler == nil {
		api.MetadataGetV2MetadataPlatformNetworkEntityAttributeHandler = metadata.GetV2MetadataPlatformNetworkEntityAttributeHandlerFunc(func(params metadata.GetV2MetadataPlatformNetworkEntityAttributeParams) middleware.Responder {
			return middleware.NotImplemented("operation metadata.GetV2MetadataPlatformNetworkEntityAttribute has not yet been implemented")
		})
	}
	if api.MetadataGetV2MetadataPlatformNetworkEntityAttributeFilterHandler == nil {
		api.MetadataGetV2MetadataPlatformNetworkEntityAttributeFilterHandler = metadata.GetV2MetadataPlatformNetworkEntityAttributeFilterHandlerFunc(func(params metadata.GetV2MetadataPlatformNetworkEntityAttributeFilterParams) middleware.Responder {
			return middleware.NotImplemented("operation metadata.GetV2MetadataPlatformNetworkEntityAttributeFilter has not yet been implemented")
		})
	}
	if api.MetadataGetV2MetadataPlatformNetworkEntityAttributesHandler == nil {
		api.MetadataGetV2MetadataPlatformNetworkEntityAttributesHandler = metadata.GetV2MetadataPlatformNetworkEntityAttributesHandlerFunc(func(params metadata.GetV2MetadataPlatformNetworkEntityAttributesParams) middleware.Responder {
			return middleware.NotImplemented("operation metadata.GetV2MetadataPlatformNetworkEntityAttributes has not yet been implemented")
		})
	}
	if api.MetadataGetV2MetadataPlatformNetworksHandler == nil {
		api.MetadataGetV2MetadataPlatformNetworksHandler = metadata.GetV2MetadataPlatformNetworksHandlerFunc(func(params metadata.GetV2MetadataPlatformNetworksParams) middleware.Responder {
			return middleware.NotImplemented("operation metadata.GetV2MetadataPlatformNetworks has not yet been implemented")
		})
	}
	if api.MetadataGetV2MetadataPlatformsHandler == nil {
		api.MetadataGetV2MetadataPlatformsHandler = metadata.GetV2MetadataPlatformsHandlerFunc(func(params metadata.GetV2MetadataPlatformsParams) middleware.Responder {
			return middleware.NotImplemented("operation metadata.GetV2MetadataPlatforms has not yet been implemented")
		})
	}
	if api.QueryPostV2DataPlatformNetworkEntityHandler == nil {
		api.QueryPostV2DataPlatformNetworkEntityHandler = query.PostV2DataPlatformNetworkEntityHandlerFunc(func(params query.PostV2DataPlatformNetworkEntityParams) middleware.Responder {
			return middleware.NotImplemented("operation query.PostV2DataPlatformNetworkEntity has not yet been implemented")
		})
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
