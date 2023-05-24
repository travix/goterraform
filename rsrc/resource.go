package rsrc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Resource required functions on resource.
type Resource[T any] interface {
	Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse, data T) (T, diag.Diagnostics)
	Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse, data T) (T, diag.Diagnostics)
	Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse, data T) (T, diag.Diagnostics)
	Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse, data T) diag.Diagnostics
}

type CanMetadata interface {
	Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse)
}

type CanSchema interface {
	Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse)
}

type CanConfigure interface {
	Configure(context.Context, resource.ConfigureRequest, *resource.ConfigureResponse)
}

type CanConfigValidators interface {
	ConfigValidators(context.Context) []resource.ConfigValidator
}

type CanImportState interface {
	ImportState(context.Context, resource.ImportStateRequest, *resource.ImportStateResponse)
}

type CanModifyPlan interface {
	ModifyPlan(context.Context, resource.ModifyPlanRequest, *resource.ModifyPlanResponse)
}

type CanUpgradeState interface {
	UpgradeState(context.Context) map[int64]resource.StateUpgrader
}

type CanValidateConfig interface {
	ValidateConfig(context.Context, resource.ValidateConfigRequest, *resource.ValidateConfigResponse)
}
