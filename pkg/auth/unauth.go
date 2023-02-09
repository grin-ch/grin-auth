package auth

import "context"

type Unauther struct {
}

// AuthFuncOverride 覆盖authFunc
// 跳过部分接口的权限验证
func (*Unauther) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, nil
}
