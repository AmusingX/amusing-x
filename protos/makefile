.DEFAULT_GOAL := all

all: generate-pangu generate-charon genreate-pluto genreate-ganymede genreate-comm

generate-pangu:
	protoc -I . \
          --go_out ./pangu/service \
          --go_opt paths=source_relative \
          --go-grpc_out ./pangu/service --go-grpc_opt paths=source_relative \
          --grpc-gateway_out ./pangu/service --grpc-gateway_opt paths=source_relative \
          ./pangu/proto/*.proto

generate-charon:
	protoc -I . \
              --go_out ./charon/service \
              --go_opt paths=source_relative \
              --go-grpc_out ./charon/service --go-grpc_opt paths=source_relative \
              --grpc-gateway_out ./charon/service --grpc-gateway_opt paths=source_relative \
              ./charon/proto/*.proto

genreate-pluto:
	protoc -I . \
              --go_out ./pluto/service \
              --go_opt paths=source_relative \
              --go-grpc_out ./pluto/service --go-grpc_opt paths=source_relative \
              --grpc-gateway_out ./pluto/service --grpc-gateway_opt paths=source_relative \
              ./pluto/proto/*.proto

genreate-ganymede:
	protoc -I . \
              --go_out ./ganymede/service \
              --go_opt paths=source_relative \
              --go-grpc_out ./ganymede/service --go-grpc_opt paths=source_relative \
              --grpc-gateway_out ./ganymede/service --grpc-gateway_opt paths=source_relative \
              ./ganymede/proto/*.proto

genreate-comm:
	protoc -I . \
              --go_out ./ \
              --go_opt paths=source_relative \
              --go-grpc_out ./ --go-grpc_opt paths=source_relative \
              --grpc-gateway_out ./ --grpc-gateway_opt paths=source_relative \
              ./comm/response/*.proto

genreate-ganymede:
	protoc -I . \
              --go_out ./europa/service \
              --go_opt paths=source_relative \
              --go-grpc_out ./europa/service --go-grpc_opt paths=source_relative \
              --grpc-gateway_out ./europa/service --grpc-gateway_opt paths=source_relative \
              ./europa/proto/*.proto
