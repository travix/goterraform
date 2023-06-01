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

func GetModel[T Model](ctx context.Context, raw tftypes.Value, get func(ctx context.Context, target interface{}) diag.Diagnostics) (*T, diag.Diagnostics) {
	var val *map[string]tftypes.Value
	var data T
	err := raw.As(&val)
	if err != nil {
		tflog.Warn(ctx, fmt.Sprintf("failed to convert raw value to map[string]tftypes.Value: %q", err))
		goto GETDATA
	}
	for k, v := range data.AttributeDefaultValues(ctx) {
		if x, ok := (*val)[k]; !ok || x.IsNull() {
			(*val)[k] = v
		}
	}
GETDATA:
	diagnostics := get(ctx, &data)
	if diagnostics.HasError() {
		return nil, diagnostics
	}
	return &data, nil
}
