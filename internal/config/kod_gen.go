// Code generated by "kod generate". DO NOT EDIT.
//go:build !ignoreKodGen

package config

import (
	"context"
	"github.com/go-kod/kod"
	"github.com/go-kod/kod/interceptor"
	"reflect"
)

// Full method names for components.
const ()

func init() {
	kod.Register(&kod.Registration{
		Name:      "github.com/go-kod/grpc-gateway/internal/config/Config",
		Interface: reflect.TypeOf((*Config)(nil)).Elem(),
		Impl:      reflect.TypeOf(config{}),
		Refs:      ``,
		LocalStubFn: func(ctx context.Context, info *kod.LocalStubFnInfo) any {
			return config_local_stub{
				impl:        info.Impl.(Config),
				interceptor: info.Interceptor,
			}
		},
	})
}

// CodeGen version check.
var _ kod.CodeGenLatestVersion = kod.CodeGenVersion[[0][1]struct{}](`
ERROR: You generated this file with 'kod generate' (codegen
version v0.1.0). The generated code is incompatible with the version of the
github.com/go-kod/kod module that you're using. The kod module
version can be found in your go.mod file or by running the following command.

    go list -m github.com/go-kod/kod

We recommend updating the kod module and the 'kod generate' command by
running the following.

    go get github.com/go-kod/kod@latest
    go install github.com/go-kod/kod/cmd/kod@latest

Then, re-run 'kod generate' and re-build your code. If the problem persists,
please file an issue at https://github.com/go-kod/kod/issues.
`)

// kod.InstanceOf checks.
var _ kod.InstanceOf[Config] = (*config)(nil)

// Local stub implementations.
// config_local_stub is a local stub implementation of [Config].
type config_local_stub struct {
	impl        Config
	interceptor interceptor.Interceptor
}

// Check that [config_local_stub] implements the [Config] interface.
var _ Config = (*config_local_stub)(nil)

// Config wraps the method [config.Config].
func (s config_local_stub) Config() (r0 *ConfigInfo) {
	// Because the first argument is not context.Context, so interceptors are not supported.
	r0 = s.impl.Config()
	return
}
