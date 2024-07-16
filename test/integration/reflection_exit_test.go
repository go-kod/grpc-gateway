package integration

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-kod/kod"
	"github.com/nautilus/graphql"
	"github.com/sysulq/graphql-gateway/pkg/server"
	"github.com/sysulq/graphql-gateway/test"
	"go.uber.org/mock/gomock"
)

func TestReflectionExit(t *testing.T) {
	infos := test.SetupDeps(t)

	mockConfig := server.NewMockConfig(gomock.NewController(t))
	mockConfig.EXPECT().Config().Return(&server.ConfigInfo{
		Grpc: server.Grpc{
			Services: []*server.Service{
				{
					Address:    infos.ConstructsServerAddr.Addr().String(),
					Reflection: true,
				},
				{
					Address:    infos.OptionsServerAddr.Addr().String(),
					Reflection: true,
				},
			},
		},
	}).AnyTimes()

	kod.RunTest(t, func(ctx context.Context, s server.Gateway) {
		gatewayUrl := test.SetupGateway(t, s)
		querier := graphql.NewSingleRequestQueryer(gatewayUrl)

		t.Run("stop do not panic", func(t *testing.T) {
			infos.OptionServer.Stop()

			recv := map[string]interface{}{}
			if err := querier.Query(context.Background(), &graphql.QueryInput{
				Query: constructsAnyQuery,
			}, &recv); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(recv, constructsAnyResponse) {
				t.Errorf("mutation failed: expected: %s got: %s", constructsAnyResponse, recv)
			}
		})
	}, kod.WithFakes(kod.Fake[server.Config](mockConfig)))
}
