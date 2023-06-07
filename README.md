# adminv-api
adminv-api




```api
goctl model mysql ddl -c -src ./rpc/doc/sql.sql -dir ./rpc/model/sys_model --style go_zero

goctl model mysql ddl -src ./rpc/doc/sql.sql -dir ./rpc/model/sys_model --style go_zero

###
goctl rpc protoc ./rpc/sys/sys.proto --go_out=./rpc/sys/ --go-grpc_out=./rpc/sys/ --zrpc_out=.rpc/sys  --style go_zero

###
goctl api go --api ./api/doc/admin.api --dir ./api/   --style go_zero
```