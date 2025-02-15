export ROOT_MOD=github.com/readlnh/biz-demo/gomall
.PHONY: gen-demo-proto
gen-demo-proto:
	@cd demo/demo_proto && cwgo server -I ../../idl --type RPC --module github.com/readlnh/biz-demo/gomall/demo/demo-proto --service demo_proto --idl ../../idl/echo.proto

.PHONY: gen-demo-thrift
gen-demo-thrift:
		@cd demo/demo_thrift && cwgo server -I ../../idl --type RPC --module github.com/readlnh/biz-demo/gomall/demo/demo-thrift --service demo_thrift --idl ../../idl/echo.thrift

.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server -I ../../idl --type HTTP --service frontend --module github.com/readlnh/biz-demo/gomall/app/frontend --idl ../../idl/frontend/home.proto

.PHONY: gen-frontend-idl
gen-frontend-idl:
	@cd app/frontend && cwgo server -I ../../idl --type HTTP --service frontend --module ${ROOT_MOD}/app/frontend --idl ../../idl/frontend/*.proto

.PHONY: gen-user
gen-user: 
	@cd rpc_gen && cwgo client --type RPC --service user --module ${ROOT_MOD}/rpc_gen  -I ../idl  --idl ../idl/user.proto
	@cd app/user && cwgo server --type RPC --service user --module ${ROOT_MOD}/app/user --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/user.proto

.PHONY: gen-product
gen-product: 
	@cd rpc_gen && cwgo client --type RPC --service product --module ${ROOT_MOD}/rpc_gen  -I ../idl  --idl ../idl/product.proto
	@cd app/product && cwgo server --type RPC --service product --module ${ROOT_MOD}/app/product --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/product.proto