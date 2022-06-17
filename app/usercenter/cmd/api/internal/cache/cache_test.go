package cache

import (
	"database/sql"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/syncx"
	"go-zero-demo-single/app/usercenter/cmd/api/internal/config"
	"testing"
)

func TestCacheSet(t *testing.T) {
	var c config.Config
	conf.MustLoad("/home/zzx/go-proj/GOPATH/src/go-zero-demo-single/app/usercenter/cmd/api/etc/usercenter.yaml", &c)

	cc := cache.New(c.Cache, syncx.NewSingleFlight(), cache.NewStat("sqlc"), sql.ErrNoRows)

	cc.Set("test", "test redis")
}
