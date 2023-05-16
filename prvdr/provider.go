package prvdr

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

type Provider interface {
	DataSources(ctx context.Context) []func() datasource.DataSource
	Resources(ctx context.Context) []func() resource.Resource
}

type CanMetadata interface {
	Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse)
}

type CanSchema interface {
	Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse)
}

type CanConfigure interface {
	Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse)
}

type CanConfigValidators interface {
	ConfigValidators(context.Context) []provider.ConfigValidator
}

type CanMetaSchema interface {
	MetaSchema(context.Context, provider.MetaSchemaRequest, *provider.MetaSchemaResponse)
}

type CanValidateConfig interface {
	ValidateConfig(context.Context, provider.ValidateConfigRequest, *provider.ValidateConfigResponse)
}
