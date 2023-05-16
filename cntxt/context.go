package cntxt

import "context"

type Key string

func IfValue[T any](ctx context.Context, key string) (T, bool) {
	val := ctx.Value(Key(key))
	if value, ok := val.(T); ok {
		return value, ok
	}
	var nilT T
	return nilT, false
}

func Value[T any](ctx context.Context, key string) T {
	val, _ := IfValue[T](ctx, key)
	return val
}

func WithValue(ctx context.Context, key string, value any) context.Context {
	return context.WithValue(ctx, Key(key), value)
}
