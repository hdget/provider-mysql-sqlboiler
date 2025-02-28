package mysql_sqlboiler

import (
	"github.com/hdget/common/types"
	"github.com/hdget/provider-mysql-sqlboiler/pkg"
	"go.uber.org/fx"
)

const (
	providerName = "mysql-sqlboiler"
)

var Capability = &types.Capability{
	Category: types.ProviderCategoryDb,
	Name:     providerName,
	Module: fx.Module(
		providerName,
		fx.Provide(pkg.New),
	),
}
