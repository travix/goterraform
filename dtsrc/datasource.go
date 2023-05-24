package dtsrc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// Datasource required functions on datasource.
type Datasource[T any] interface {
	Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse, data T) (T, diag.Diagnostics)
}

type CanMetadata interface {
	Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse)
}

type CanSchema interface {
	Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse)
}

type CanConfigure interface {
	Configure(context.Context, datasource.ConfigureRequest, *datasource.ConfigureResponse)
}

type CanConfigValidators interface {
	ConfigValidators(context.Context) []datasource.ConfigValidator
}

type CanValidateConfig interface {
	ValidateConfig(context.Context, datasource.ValidateConfigRequest, *datasource.ValidateConfigResponse)
}
