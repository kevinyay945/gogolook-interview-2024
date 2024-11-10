.PHONY: init
init:
	go install go.uber.org/mock/mockgen@v0.5.0
	go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.2.0

.PHONY: openapi_http
openapi_http:
	oapi-codegen -generate types -o "http/openapi_types.gen.go" -package "http" "assets/swagger/swagger.yml"
	oapi-codegen -generate server -o "http/openapi_api.gen.go" -package "http" "assets/swagger/swagger.yml"

