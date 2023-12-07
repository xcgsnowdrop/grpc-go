# 完整的jwt认证示例
- 包含token刷新逻辑

# 重新编译proto文件
- 进入目录: grpc-go/examples/auth/
- 运行：protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/auth_service.proto