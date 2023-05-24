package prvdr

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"google.golang.org/grpc"
)

type Provider interface {
	DataSources(ctx context.Context) []func() datasource.DataSource
	Resources(ctx context.Context) []func() resource.Resource
}

type CanConfigureGrpc[T any] interface {
	ConfigureGrpc(ctx context.Context, data T) (grpc.ClientConnInterface, diag.Diagnostics)
}

type CanConfigure interface {
	Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse)
}

type CanConfigValidators interface {
	ConfigValidators(context.Context) []provider.ConfigValidator
}

type CanMetadata interface {
	Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse)
}

type CanMetaSchema interface {
	MetaSchema(context.Context, provider.MetaSchemaRequest, *provider.MetaSchemaResponse)
}

type CanSchema interface {
	Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse)
}

type CanValidateConfig interface {
	ValidateConfig(context.Context, provider.ValidateConfigRequest, *provider.ValidateConfigResponse)
}
