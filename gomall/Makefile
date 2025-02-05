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
	@cd app/frontend && cwgo server -I ../../idl --type HTTP --service frontend --module github.com/readlnh/biz-demo/gomall/app/frontend --idl ../../idl/frontend/*.proto
