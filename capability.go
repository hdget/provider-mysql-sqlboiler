package sqlboiler_mysql

import (
	"github.com/hdget/common/types"
	"github.com/hdget/provider-mysql-sqlboiler/pkg"
	"go.uber.org/fx"
)

var Capability = &types.Capability{
	Category: types.ProviderCategoryDb,
	Module: fx.Module(
		string(types.ProviderNameDbSqlBoilerMysql),
		fx.Provide(pkg.New),
	),
}
