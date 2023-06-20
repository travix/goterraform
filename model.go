package gotf

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type Model interface {
	AttributeDefaultValues(ctx context.Context) map[string]tftypes.Value
}

func GetModel[T Model](ctx context.Context, raw tftypes.Value, target T, get func(ctx context.Context, target interface{}) diag.Diagnostics) diag.Diagnostics {
	var val *map[string]tftypes.Value
	err := raw.As(&val)
	if err != nil {
		tflog.Warn(ctx, fmt.Sprintf("failed to convert raw value to map[string]tftypes.Value: %q", err))
		goto GETDATA
	}
	for k, v := range target.AttributeDefaultValues(ctx) {
		if x, ok := (*val)[k]; !ok || x.IsNull() {
			(*val)[k] = v
		}
	}
GETDATA:
	diagnostics := get(ctx, target)
	if diagnostics.HasError() {
		return diagnostics
	}
	return nil
}
