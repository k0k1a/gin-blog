module github.com/k0k1a/go-gin-example

go 1.17

replace (
	github.com/k0k1a/go-gin-example/conf => ../go-gin-example/pkg/conf
	github.com/k0k1a/go-gin-example/middleware => ../go-gin-example/middleware
	github.com/k0k1a/go-gin-example/models => ../go-gin-example/models
	github.com/k0k1a/go-gin-example/pkg/setting => ../go-gin-example/pkg/setting
	github.com/k0k1a/go-gin-example/routers => ../go-gin-example/routers
)

require github.com/go-ini/ini v1.66.6 // indirect
